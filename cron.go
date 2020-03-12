package main

type Cron struct {
    command string
	exitCode int
	output []byte
}

func (cron Cron) Validate() bool {
	if cron.command != "" {
		return true
	}
	return false
} 