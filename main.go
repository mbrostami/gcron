package main

import (
	"bufio"
	"flag"
	"hash/fnv"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/mbrostami/gcron/configs"
	"github.com/mbrostami/gcron/cron"
	"github.com/mbrostami/gcron/helpers"
	"github.com/mbrostami/gcron/output"

	"github.com/rs/xid"
	"github.com/shirou/gopsutil/process"
)

func main() {
	executable := flag.String("exec", "echo", "Command to execute")
	// flagRunAfter := flag.Int("run.after", 0, "Run command after given seconds")
	flagLockEnabled := flag.Bool("lock.enable", false, "Enable mutex lock")
	flagLockName := flag.String("lock.name", "", "Mutex name")

	// Override config file values
	flag.Bool("out.tags", false, "Output tags")
	flag.Bool("out.clean", false, "Only command output")
	flag.String("server.tcp.port", "", "TCP Server port")
	flag.String("server.tcp.host", "", "TCP Server host")
	flag.String("server.udp.port", "", "UDP Server port")
	flag.String("server.udp.host", "", "UDP Server host")
	flag.String("server.unix.path", "/tmp/gcron-server.sock", "UNIX socket path")

	cfg := configs.GetConfig(flag.CommandLine)
	crontask := cron.Task{
		Command: *executable,
	}
	if crontask.Validate() {

		hostname, _ := os.Hostname()
		crontask.Hostname = hostname
		crontask.UID = hash(crontask.Command)

		var mtx *helpers.Mutex
		if *flagLockEnabled {
			mutexName := strconv.FormatUint(uint64(crontask.UID), 10)
			if *flagLockName != "" {
				mutexName = *flagLockName
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

		// Setup log
		log.SetFlags(0)
		if cfg.Out.Clean == false {
			log.SetFlags(log.Ldate | log.Ltime)
		}
		f, err := os.OpenFile(cfg.Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(
			os.Stdout,
			f,
			// createSysLog(getConfig()),
		)
		log.SetOutput(wrt)

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
		crontask.GUID = xid.New().String() // sortable guid
		// Stream output
		go func() {
			for scanner.Scan() {
				stdChan <- scanner.Bytes()
			}
			for errScanner.Scan() {
				stdChan <- errScanner.Bytes()
			}
			close(stdChan)
			done <- true
		}()

		crontask.StartTime = time.Now()
		crontask.Success = false
		cmd.Start()
		crontask.Pid = cmd.Process.Pid
		for output := range stdChan {
			log.Printf("%s", string(output))
			crontask.Output = append(crontask.Output, output...)
		}
		<-done
		if mtx != nil {
			mtx.Release()
		}
		cmd.Wait()

		proc, _ := process.NewProcess(int32(cmd.Process.Pid))
		parent, _ := process.NewProcess(int32(os.Getppid()))
		crontask.Parent, _ = parent.Name()
		crontask.Username, _ = proc.Username()
		crontask.Success = cmd.ProcessState.Success()
		crontask.SystemTime = cmd.ProcessState.SystemTime()
		crontask.UserTime = cmd.ProcessState.UserTime()
		crontask.ExitCode = cmd.ProcessState.ExitCode()
		crontask.EndTime = time.Now()

		// Log tags
		if cfg.Out.Tags == true {
			log.Printf(
				"[uid:%d] [duration:%vs] [systime:%vs] [usertime:%vs] [status:%v]",
				crontask.UID,
				crontask.EndTime.Sub(crontask.StartTime).Seconds(),
				crontask.SystemTime.Seconds(),
				crontask.UserTime.Seconds(),
				crontask.Success,
			)
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
