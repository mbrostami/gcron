package helpers

import (
	"github.com/mbrostami/go-filemutex"
)

// Mutex use for mutex lock
type Mutex struct {
	handler *filemutex.FileMutex
}

// NewMutex create new filemutex handler
func NewMutex(key string) (*Mutex, error) {
	mutex, err := filemutex.New("/tmp/gcron-" + key + ".lock")
	return &Mutex{handler: mutex}, err
}

// Lock locks by flock
func (m Mutex) Lock() (bool, error) {
	err := m.handler.TryLock()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Release lock
func (m Mutex) Release() (bool, error) {
	defer m.handler.Close()
	err := m.handler.Unlock()
	if err != nil {
		return false, err
	}
	return true, nil
}
