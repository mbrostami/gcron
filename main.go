package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"log/syslog"
	"os"
	"os/exec"
)

func main() {
	executable := flag.String("exec", "", "Command to execute")
	flag.Parse()
	crontask := Cron{
		command: *executable,
	}
	if crontask.Validate() {
		//setupLogger(getConfig().Log.Path)
		f, err := os.OpenFile(getConfig().Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)
		cmd := exec.Command("bash", "-c", crontask.command)
		logger := createSysLog()
		log.Printf("Command \"%s\" started!\n", crontask.command)
		log.Println("------ OUTPUT BEGIN")
		logger.Warning("-----Warning...")

		cmdReader, err := cmd.StdoutPipe()
		cmdErrReader, _ := cmd.StderrPipe()
		if err != nil {
			log.Printf("%s error: %s", os.Stderr, err)
			os.Exit(1)
		}

		var out string
		done := make(chan struct{})
		scanner := bufio.NewScanner(cmdReader)
		errScanner := bufio.NewScanner(cmdErrReader)
		go func() {
			for scanner.Scan() {
				log.Printf("%s", scanner.Text())
				out = "Error:" + scanner.Text()
				crontask.output = append(crontask.output, out...)
			}
			for errScanner.Scan() {
				log.Printf("Err %s", errScanner.Text())
				out = "Error:" + errScanner.Text()
				crontask.output = append(crontask.output, out...)
				//err = errScanner.Text()
			}
			done <- struct{}{}
		}()
		cmd.Start()
		<-done
		cmd.Wait()
		//	if err != nil {
		//		if status, ok := err.(*exec.ExitError); ok {
		//			crontask.exitCode = status.ExitCode()
		//			crontask.output = status.Stderr
		//			log.Printf("Error: %s", status.Stderr)
		//		}
		//	}
		log.Printf("------ OUTPUT END %d", crontask.exitCode)
		log.Println("Command finished")
	}
}
func createSysLog() *syslog.Writer {
	w, err := syslog.Dial("udp", "localhost:514", syslog.LOG_WARNING, "testtag")
	if err != nil {
		log.Fatal("failed to dial syslog")
	}
	return w
}
