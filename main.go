package main

import (
	"flag"
	"fmt"
	"os/exec"
)
type cron struct {
    command string
	exitCode int
	output []byte
	valid bool
}
func main() {
	executable := flag.String("exec", "", "Command to execute")
	flag.Parse()
	crontask := createCron(*executable)
	if crontask.valid {
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

func createCron(executable string) *cron {
   crontask := cron{
       command: executable,
   }
   crontask.valid = false
   if executable != "" {
     crontask.valid = true
   }
   return &crontask
}
