package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"log/syslog"
	"os"
	"os/exec"
	"time"
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
		wrt := io.MultiWriter(os.Stdout, f, createSysLog(getConfig()))
		log.SetOutput(wrt)

		cmd := exec.Command("bash", "-c", crontask.command)

		log.Printf("Command \"%s\" started!\n", crontask.command)
		log.Println("------ OUTPUT BEGIN")

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
		go func() {
			for scanner.Scan() {
				stdChan <- scanner.Text() + "\n"
			}
			for errScanner.Scan() {
				stdChan <- "Error: " + errScanner.Text() + "\n"
			}
			close(stdChan)
		}()

		crontask.startTime = time.Now()
		cmd.Start()
		for output := range stdChan {
			log.Printf("%s", output)
			crontask.output = append(crontask.output, output...)
		}
		cmd.Wait()
		crontask.endTime = time.Now()
		log.Printf("------ OUTPUT END %d", crontask.exitCode)
	}
}

func createSysLog(cfg Config) *syslog.Writer {
	writer, err := syslog.Dial("udp", "localhost:5140", cfg.GetLogLevel(), "gcron")
	if err != nil {
		log.Fatal("failed to dial syslog")
	}
	return writer
}
