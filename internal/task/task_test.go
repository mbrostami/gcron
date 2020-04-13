package task_test

import (
	"fmt"
	"testing"

	"github.com/mbrostami/gcron/internal/task"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		command string
		valid   bool
	}{
		{"", false},
		{"echo HelloWolrd", true},
	}
	for _, item := range tests {
		testname := fmt.Sprintf("command: %s, valid: %v", item.command, item.valid)
		t.Run(testname, func(t *testing.T) {
			task := task.Task{}
			task.Command = item.command
			res, _ := task.Validate()
			if res != item.valid {
				t.Errorf("got %v, want %v", res, item.valid)
			}
		})
	}
}
