package main

import "time"

type Cron struct {
	command   string
	startTime time.Time
	endTime   time.Time
	exitCode  int
	output    []byte
}

func (cron Cron) Validate() bool {
	if cron.command != "" {
		return true
	}
	return false
}
