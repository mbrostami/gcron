package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	executable := flag.String("exec", "", "Command to execute")
	flag.Parse()
	crontask := Cron{
       command: *executable,
    }
	if crontask.Validate() {
		cmd := exec.Command("bash", "-c", crontask.command)
		fmt.Printf("Command \"%s\" started!\n", crontask.command)
		fmt.Printf("------ OUTPUT BEGIN\n")
		out, err := cmd.Output()
		crontask.output = out
		if err != nil {
			if status, ok := err.(*exec.ExitError); ok {
				crontask.exitCode = status.ExitCode()
				crontask.output = status.Stderr
			}
		}
		fmt.Printf("%s", crontask.output)
		fmt.Printf("------ OUTPUT END %d\n", crontask.exitCode)
		fmt.Printf("Command finished \n")
	}
}
