package main

import (
	"bufio"
	"flag"
	"gcron/configs"
	"gcron/cron"
	"gcron/output"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/rs/xid"
	"github.com/shirou/gopsutil/process"
)

func main() {
	executable := flag.String("exec", "echo", "Command to execute")
	configpath := flag.String("config", ".", "Config file path")

	// Override config file values
	flag.Bool("out.notime", false, "Clean output")
	flag.Bool("out.clean", false, "Clean output")
	flag.String("server.tcp.port", "", "TCP Server port")
	flag.String("server.tcp.host", "", "TCP Server host")
	flag.String("server.udp.port", "", "UDP Server port")
	flag.String("server.udp.host", "", "UDP Server host")
	flag.String("server.unix.path", "/tmp/gcron-server.sock", "UNIX socket path")

	cfg := configs.GetConfig(*configpath, flag.CommandLine)
	crontask := cron.Task{
		Command: *executable,
	}
	if crontask.Validate() {
		if cfg.Out.Notime == false {
			log.SetFlags(log.Ldate | log.Ltime)
		} else {
			log.SetFlags(0)
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
			log.Printf("%v error: %s", os.Stderr, err)
			os.Exit(1)
		}
		stdErrReader, err := cmd.StderrPipe()
		if err != nil {
			log.Printf("%v error: %s", os.Stderr, err)
			os.Exit(1)
		}

		stdChan := make(chan []byte)
		scanner := bufio.NewScanner(stdOutReader)
		errScanner := bufio.NewScanner(stdErrReader)
		crontask.GUID = xid.New().String() // sortable guid
		go func() {
			for scanner.Scan() {
				stdChan <- scanner.Bytes()
			}
			for errScanner.Scan() {
				stdChan <- errScanner.Bytes()
			}
			close(stdChan)
		}()

		crontask.StartTime = time.Now()
		crontask.Success = false
		cmd.Start()
		crontask.Pid = cmd.Process.Pid
		for output := range stdChan {
			log.Printf("%s", string(output))
			crontask.Output = append(crontask.Output, output...)
		}
		cmd.Wait()

		p, _ := process.NewProcess(int32(cmd.Process.Pid))
		crontask.Username, _ = p.Username()
		crontask.Success = cmd.ProcessState.Success()
		crontask.SystemTime = cmd.ProcessState.SystemTime()
		crontask.UserTime = cmd.ProcessState.UserTime()
		crontask.ExitCode = cmd.ProcessState.ExitCode()
		crontask.EndTime = time.Now()

		// Log tags
		if cfg.Out.Clean == false {
			log.Printf(
				"[guid:%s] [duration:%vs] [systime:%vs] [usertime:%vs] [status:%v]",
				crontask.GUID,
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
