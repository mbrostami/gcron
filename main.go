package main

import (
	"bufio"
	"flag"
	"hash/fnv"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/mbrostami/gcron/configs"
	"github.com/mbrostami/gcron/cron"
	"github.com/mbrostami/gcron/helpers"
	"github.com/mbrostami/gcron/output"
	"github.com/mbrostami/gcron/validators"
	"github.com/rs/xid"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
)

func main() {
	executable := flag.String("c", "", "Command to execute")
	flagLockEnabled := flag.Bool("lock.enable", false, "Enable mutex lock")
	flagLockRemote := flag.Bool("lock.remote", false, "Use rpc mutex lock")
	flagLockName := flag.String("lock.name", "", "Mutex name")
	flagOverride := flag.String("override", "", "Override command status by regex match in output")
	flagDelay := flag.Int("delay", 0, "Delay running command in seconds")
	flag.Bool("out.tags", false, "Output tags")
	flag.Bool("out.hide.systime", false, "Hide system time tag")
	flag.Bool("out.hide.usertime", false, "Hide user time tag")
	flag.Bool("out.hide.duration", false, "Hide duration tag")
	flag.Bool("out.hide.uid", false, "Hide uid tag")
	// flag.Bool("out.clean", false, "Only command output")
	flag.Bool("log.enable", false, "Enable log")
	flag.String("server.tcp.port", "", "TCP Server port")
	flag.String("server.tcp.host", "", "TCP Server host")
	flag.String("server.udp.host", "", "UDP Server host")
	flag.String("server.udp.port", "", "UDP Server port")
	flag.String("server.rpc.host", "", "RPC Server host")
	flag.String("server.rpc.port", "", "RPC Server port")
	flag.String("server.unix.path", "/tmp/gcron-server.sock", "UNIX socket path")
	flag.String("log.level", "warning", "Log level")

	cfg := configs.GetConfig(flag.CommandLine)
	crontask := cron.Task{
		Command:   *executable,
		FLock:     *flagLockEnabled,
		FLockName: *flagLockName,
		FOverride: *flagOverride,
		FDelay:    *flagDelay,
	}
	processCommand(cfg, crontask, *flagLockRemote)
}

func processCommand(cfg configs.Config, crontask cron.Task, remoteLock bool) {

	if crontask.Validate() {

		log.SetLevel(cfg.GetLogLevel())
		// Setup log
		log.SetFormatter(&nested.Formatter{
			NoColors: false,
		})

		crontask.GUID = xid.New().String() // sortable guid
		hostname, _ := os.Hostname()
		crontask.Hostname = hostname
		crontask.UID = hash(crontask.Command)

		var grpcHandler output.GrpcHandler
		if cfg.Server.RPC.Enabled {
			grpcHandler, _ = output.NewGrpcHandler(cfg.Server.RPC.Host, cfg.Server.RPC.Port)
			defer grpcHandler.Close()
			if crontask.FLock && remoteLock { // remote lock can only be used with rpc
				locked, _ := grpcHandler.Lock(strconv.FormatUint(uint64(crontask.UID), 10))
				if !locked {
					log.Fatal("Task is already running...")
				}
			}
		}

		var mtx *helpers.Mutex
		if crontask.FLock && !remoteLock { // remotelock with locallock can not be enabled at the same time
			mutexName := strconv.FormatUint(uint64(crontask.UID), 10)
			if crontask.FLockName != "" {
				mutexName = crontask.FLockName
			}
			mtx, err := helpers.NewMutex(mutexName)
			if err != nil {
				log.Fatalf("Couldn't create lock: %v", err)
			}
			locked, _ := mtx.Lock()
			if !locked {
				log.Fatal("Task is already running...")
			}
		}

		// FIXME: Prevent IO Block
		log.SetOutput(os.Stdout)
		if cfg.Log.Enable {
			f, err := os.OpenFile(cfg.Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Panicf("error opening file: %v", err)
			}
			defer f.Close()
			writers := io.MultiWriter(
				os.Stdout,
				f,
			)
			log.SetOutput(writers)
		}

		// Delay running command
		if crontask.FDelay > 0 {
			time.Sleep(time.Duration(crontask.FDelay) * time.Second)
		}

		cmd := exec.Command("bash", "-c", crontask.Command)

		stdOutReader, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatalf("%v error: %s", os.Stderr, err)
		}
		stdErrReader, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalf("%v error: %s", os.Stderr, err)
		}

		stdChan := make(chan []byte)
		done := make(chan bool)
		scanner := bufio.NewScanner(stdOutReader)
		errScanner := bufio.NewScanner(stdErrReader)
		// Stream output
		go func() {
			for scanner.Scan() {
				log.Infof("%s", string(scanner.Bytes()))
				stdChan <- scanner.Bytes()
			}
			for errScanner.Scan() {
				log.Warnf("%s", string(errScanner.Bytes()))
				stdChan <- errScanner.Bytes()
			}
			close(stdChan)
			done <- true
		}()

		crontask.StartTime = time.Now()
		crontask.Success = false
		cmd.Start()
		crontask.Pid = cmd.Process.Pid
		var statusByRegex = false
		for output := range stdChan {
			// output = append(output, []byte("\n")...)
			crontask.Output = append(crontask.Output, output...)
			if crontask.FOverride != "" {
				statusByRegex = statusByRegex || validators.NewRegex(crontask.FOverride).Validate(string(output))
			}
			if cfg.Server.RPC.Enabled {
				grpcHandler.Log(crontask.GUID, output)
			}
		}
		<-done
		if crontask.FLock && remoteLock {
			if cfg.Server.RPC.Enabled {
				grpcHandler.Release(strconv.FormatUint(uint64(crontask.UID), 10))
			}
		} else if crontask.FLock && !remoteLock {
			if mtx != nil {
				mtx.Release()
			}
		}
		cmd.Wait()
		crontask.Success = cmd.ProcessState.Success()
		if crontask.FOverride != "" {
			crontask.Success = statusByRegex
		}

		proc, _ := process.NewProcess(int32(cmd.Process.Pid))
		parent, _ := process.NewProcess(int32(os.Getppid()))
		crontask.Parent, _ = parent.Name()
		crontask.Username, _ = proc.Username()
		crontask.SystemTime = cmd.ProcessState.SystemTime()
		crontask.UserTime = cmd.ProcessState.UserTime()
		crontask.ExitCode = cmd.ProcessState.ExitCode()
		crontask.EndTime = time.Now()

		if cfg.Server.RPC.Enabled {
			grpcHandler.Finish(crontask)
		}
		// Log tags
		if cfg.Out.Tags == true {
			// var customOutput string
			fields := log.Fields{}
			if !cfg.Out.Hide.UID {
				fields["uid"] = crontask.UID
			}
			if !cfg.Out.Hide.SysTime {
				fields["systime"] = crontask.SystemTime.Seconds()
			}
			if !cfg.Out.Hide.UserTime {
				fields["usertime"] = crontask.UserTime.Seconds()
			}
			if !cfg.Out.Hide.Duration {
				fields["duration"] = crontask.EndTime.Sub(crontask.StartTime).Seconds()
			}
			fields["status"] = crontask.Success
			log.WithFields(fields).Info("[tags]")
		}
		// Send crontask over tcp udp and unix socket
		// FIXME: Stream output instead of sending all at once
		if cfg.Server.TCP.Enabled {
			output.SendOverTCP(
				cfg.Server.TCP.Host,
				cfg.Server.TCP.Port,
				crontask,
			)
		}
		if cfg.Server.UDP.Enabled {
			output.SendOverUPD(
				cfg.Server.UDP.Host,
				cfg.Server.UDP.Port,
				crontask,
			)
		}
		if cfg.Server.Unix.Enabled {
			output.SendOverUNIX(
				cfg.Server.Unix.Path,
				crontask,
			)
		}
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func report(cfg configs.Config) {
	return
}
