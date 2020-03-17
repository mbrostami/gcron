package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	executable := flag.String("exec", "", "Command to execute")
	configpath := flag.String("config", ".", "Config file path")
	flag.Parse()
	crontask := Cron{
		Command: *executable,
	}
	if crontask.Validate() {
		cfg := GetConfig(*configpath)
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

		log.Printf("Command \"%s\" started!\n", crontask.Command)
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
			stdChan <- "[uid][tags]" + "\n"
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

		// Send crontask over tcp -- later over udp and unix socket
		// FIXME: Stream output instead of sending all at once
		conn := tcpConnection(cfg)
		defer conn.Close()

		p, _ := process.NewProcess(int32(cmd.Process.Pid))
		crontask.Username, _ = p.Username()
		crontask.Success = cmd.ProcessState.Success()
		crontask.SystemTime = cmd.ProcessState.SystemTime()
		crontask.UserTime = cmd.ProcessState.UserTime()
		crontask.ExitCode = cmd.ProcessState.ExitCode()
		crontask.EndTime = time.Now()

		binaryBuff := new(bytes.Buffer)
		gobobj := gob.NewEncoder(binaryBuff)
		gobobj.Encode(crontask)
		conn.Write(binaryBuff.Bytes())
	}
}

func tcpConnection(cfg Config) *net.TCPConn {
	if cfg.Server.TCP.Host != "" {
		servAddr := cfg.Server.TCP.Host + ":" + cfg.Server.TCP.Port
		//log.Printf("%s", servAddr)
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}
		return conn
	}
	return nil
}

// func udpConnection(cfg Config) *net.UDPConn {
// 	// if cfg.Server.UDP.Host != nil {
// 	servAddr := cfg.Server.UDP.Host + ":" + string(cfg.Server.UDP.Port)
// 	udpAddr, err := net.ResolveUDPAddr("udp", servAddr)
// 	if err != nil {
// 		println("ResolveUDPAddr failed:", err.Error())
// 		os.Exit(1)
// 	}
// 	conn, err := net.DialUDP("udp", nil, udpAddr)
// 	if err != nil {
// 		println("Dial UDP failed:", err.Error())
// 		os.Exit(1)
// 	}
// 	return conn
// 	// }
// 	// return nil
// }

// func unixConnection(cfg Config) *net.UnixConn {
// 	// if cfg.Server.Unix.Host != nil {
// 	servAddr := cfg.Server.Unix.Host + ":" + string(cfg.Server.Unix.Port)
// 	unixAddr, err := net.ResolveUnixAddr("unix", servAddr)
// 	if err != nil {
// 		println("ResolveUDPAddr failed:", err.Error())
// 		os.Exit(1)
// 	}
// 	conn, err := net.DialUnix("udp", nil, unixAddr)
// 	if err != nil {
// 		println("Dial UDP failed:", err.Error())
// 		os.Exit(1)
// 	}
// 	return conn
// 	// }
// 	// return nil
// }
