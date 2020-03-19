package cron

import "time"

// Task keeps cronjob information
type Task struct {
	Pid        int
	GUID       string
	Username   string
	Command    string
	StartTime  time.Time
	EndTime    time.Time
	ExitCode   int
	Output     []byte
	SystemTime time.Duration
	UserTime   time.Duration
	Success    bool
}

// Validate the command
func (task Task) Validate() bool {
	if task.Command != "" {
		return true
	}
	return false
}
