package mutex

import (
	"github.com/mbrostami/go-filemutex"
)

// FileMutex use for mutex lock
type FileMutex struct {
	handler *filemutex.FileMutex
}

// Lock locks by flock
func (m *FileMutex) Lock(key string, timeout int) (bool, error) {
	if m.handler == nil {
		mutex, err := filemutex.New("/tmp/gcron-" + key + ".lock")
		if err != nil {
			return false, err
		}
		m.handler = mutex
	}
	err := m.handler.TryLock()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Release lock
func (m *FileMutex) Release(key string) (bool, error) {
	defer m.handler.Close()
	err := m.handler.Unlock()
	if err != nil {
		return false, err
	}
	return true, nil
}
