package cron

import "time"

// Task keeps cronjob information
type Task struct {
	FLock      bool
	FLockName  string
	FOverride  string
	FDelay     int
	Pid        int
	GUID       string // global unique id
	UID        uint32 // hash based on command string
	Parent     string
	Hostname   string
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
