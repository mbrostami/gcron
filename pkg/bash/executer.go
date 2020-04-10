package bash

import (
	"bufio"
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/process"
)

// Executer command executer structure
type Executer struct {
	OutputChannel chan []byte // channel to send output
	Delay         int
	Command       string
	CmdInfo       *CmdInfo
	cmd           *exec.Cmd
}

// CmdInfo command structure
type CmdInfo struct {
	StartTime  time.Time
	EndTime    time.Time
	Status     bool
	Pid        int
	Parent     string
	Username   string
	SystemTime time.Duration
	UserTime   time.Duration
	ExitCode   int
}

// MakeCommand returns a command executer
func MakeCommand(command string, delay int, outChan chan []byte) (Executer, error) {
	e := Executer{
		Command:       command,
		Delay:         delay,
		OutputChannel: outChan,
	}
	e.CmdInfo = &CmdInfo{}
	e.CmdInfo.Status = false
	e.cmd = exec.Command("bash", "-c", e.Command)
	stdOutReader, err := e.cmd.StdoutPipe()
	if err != nil {
		return e, err
	}
	stdErrReader, err := e.cmd.StderrPipe()
	if err != nil {
		return e, err
	}
	scanner := bufio.NewScanner(stdOutReader)
	errScanner := bufio.NewScanner(stdErrReader)
	go e.scan(scanner, errScanner)
	return e, nil
}

// Execute run command
func (e *Executer) Execute() error {
	// Delay running command
	if e.Delay > 0 {
		time.Sleep(time.Duration(e.Delay) * time.Second)
	}
	e.CmdInfo.StartTime = time.Now()
	err := e.cmd.Start()
	e.CmdInfo.Pid = e.cmd.Process.Pid
	return err
}

// scan stdout and stderr
func (e *Executer) scan(stdoutScanner *bufio.Scanner, stderrScanner *bufio.Scanner) {
	defer close(e.OutputChannel)
	for stdoutScanner.Scan() {
		e.OutputChannel <- stdoutScanner.Bytes()
	}
	for stderrScanner.Scan() {
		e.OutputChannel <- stderrScanner.Bytes()
	}
	e.cmd.Wait()
	e.CmdInfo.Status = e.cmd.ProcessState.Success()
	proc, _ := process.NewProcess(int32(e.cmd.Process.Pid))
	parent, _ := process.NewProcess(int32(os.Getppid()))
	e.CmdInfo.Parent, _ = parent.Name()
	e.CmdInfo.Username, _ = proc.Username()
	e.CmdInfo.SystemTime = e.cmd.ProcessState.SystemTime()
	e.CmdInfo.UserTime = e.cmd.ProcessState.UserTime()
	e.CmdInfo.ExitCode = e.cmd.ProcessState.ExitCode()
	e.CmdInfo.EndTime = time.Now()
}
