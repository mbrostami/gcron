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
	executable := flag.String("exec", "", "Command to execute")
	configpath := flag.String("config", ".", "Config file path")
	flagOutNoTime := flag.Bool("o-notime", false, "Clean output")
	flagOutClean := flag.Bool("o-clean", false, "Clean output")
	flag.Parse()
	crontask := cron.Task{
		Command: *executable,
	}
	if crontask.Validate() {
		if *flagOutNoTime == false {
			log.SetFlags(log.Ldate | log.Ltime)
		} else {
			log.SetFlags(0)
		}
		cfg := configs.GetConfig(*configpath)
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

		stdChan := make(chan string)
		scanner := bufio.NewScanner(stdOutReader)
		errScanner := bufio.NewScanner(stdErrReader)
		crontask.GUID = xid.New().String() // sortable guid
		go func() {
			for scanner.Scan() {
				stdChan <- scanner.Text() + "\n"
			}
			for errScanner.Scan() {
				stdChan <- errScanner.Text() + "\n"
			}
			close(stdChan)
		}()

		crontask.StartTime = time.Now()
		crontask.Success = false
		cmd.Start()
		crontask.Pid = cmd.Process.Pid
		for output := range stdChan {
			log.Printf("%s", output)
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
		if *flagOutClean == false {
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
		output.SendOverTCP(
			cfg.Server.TCP.Host,
			cfg.Server.TCP.Port,
			crontask,
		)
		output.SendOverUPD(
			cfg.Server.UDP.Host,
			cfg.Server.UDP.Port,
			crontask,
		)
		output.SendOverUNIX(
			cfg.Server.Unix.Path,
			crontask,
		)
	}
}
