package main

import "time"

// Cron keeps cronjob information
type Cron struct {
	Pid        int
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
