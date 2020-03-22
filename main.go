package main

import (
	"bufio"
	"flag"
	"fmt"
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
	"github.com/mbrostami/gcron/validators"

	"github.com/rs/xid"
	"github.com/shirou/gopsutil/process"
)

func main() {
	execFlagSet := flag.NewFlagSet("exec", flag.ExitOnError)
	reportsFlagSet := flag.NewFlagSet("report", flag.ExitOnError)

	executable := execFlagSet.String("c", "", "Command to execute")
	flagLockEnabled := execFlagSet.Bool("lock.enable", false, "Enable mutex lock")
	flagLockName := execFlagSet.String("lock.name", "", "Mutex name")
	flagOverride := execFlagSet.String("override", "", "Override command status by regex match in output")
	execFlagSet.Bool("out.tags", false, "Output tags")
	execFlagSet.Bool("out.hide.systime", false, "Hide system time tag")
	execFlagSet.Bool("out.hide.usertime", false, "Hide user time tag")
	execFlagSet.Bool("out.hide.duration", false, "Hide duration tag")
	execFlagSet.Bool("out.hide.uid", false, "Hide uid tag")
	execFlagSet.Bool("out.clean", false, "Only command output")
	execFlagSet.String("server.tcp.port", "", "TCP Server port")
	execFlagSet.String("server.tcp.host", "", "TCP Server host")
	execFlagSet.String("server.udp.port", "", "UDP Server port")
	execFlagSet.String("server.udp.host", "", "UDP Server host")
	execFlagSet.String("server.unix.path", "/tmp/gcron-server.sock", "UNIX socket path")

	reportsFlagSet.String("format", "text", "Test")
	if len(os.Args) == 2 {
		fmt.Println("usage: gcron <command> [<args>]")
		fmt.Println("  exec     Execute command")
		fmt.Println("  report   Generate reports")
		fmt.Println("help: gcron <command> --help")

		return
	}

	switch os.Args[1] {
	case "exec":
		cfg := configs.GetConfig(execFlagSet)
		crontask := cron.Task{
			Command:   *executable,
			FLock:     *flagLockEnabled,
			FLockName: *flagLockName,
			FOverride: *flagOverride,
		}
		processCommand(cfg, crontask)

	case "report":
		cfg := configs.GetConfig(reportsFlagSet)
		report(cfg)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}

func processCommand(cfg configs.Config, crontask cron.Task) {

	if crontask.Validate() {

		hostname, _ := os.Hostname()
		crontask.Hostname = hostname
		crontask.UID = hash(crontask.Command)

		var mtx *helpers.Mutex
		if crontask.FLock {
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

		crontask.GUID = xid.New().String() // sortable guid
		stdChan := make(chan []byte)
		done := make(chan bool)
		scanner := bufio.NewScanner(stdOutReader)
		errScanner := bufio.NewScanner(stdErrReader)
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
		var statusByRegex = false
		for output := range stdChan {
			log.Printf("%s", string(output))
			crontask.Output = append(crontask.Output, output...)
			if crontask.FOverride != "" {
				statusByRegex = statusByRegex || validators.NewRegex(crontask.FOverride).Validate(string(output))
			}
		}
		<-done
		if mtx != nil {
			mtx.Release()
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

		// Log tags
		if cfg.Out.Tags == true {
			var customOutput string
			if !cfg.Out.Hide.UID {
				customOutput += fmt.Sprintf("[uid:%vs] ", crontask.UID)
			}
			if !cfg.Out.Hide.SysTime {
				customOutput += fmt.Sprintf("[systime:%vs] ", crontask.SystemTime.Seconds())
			}
			if !cfg.Out.Hide.UserTime {
				customOutput += fmt.Sprintf("[usertime:%vs] ", crontask.UserTime.Seconds())
			}
			if !cfg.Out.Hide.Duration {
				customOutput += fmt.Sprintf("[duration:%vs] ", crontask.EndTime.Sub(crontask.StartTime).Seconds())
			}
			log.Printf(
				"%s[status:%v]",
				customOutput,
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

func report(cfg configs.Config) {
	return
}
