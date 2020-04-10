package task

import (
	"hash/fnv"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

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

// SetBasics set defaults
func (task Task) SetBasics() {
	task.GUID = xid.New().String()
	hostname, _ := os.Hostname()
	task.Hostname = hostname
	task.UID = hash(task.Command)
}

// Validate the command
func (task Task) Validate() (bool, error) {
	if task.Command != "" {
		return true, nil
	}
	err := errors.Errorf("Command %s is not valie", task.Command)
	return false, err
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
