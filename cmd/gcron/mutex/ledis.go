package mutex

import (
	"github.com/mbrostami/gcron/internal/db"
)

// LedisdbMutex to use as mutex lock
type LedisdbMutex struct {
	Server *db.LedisDB
}

// Lock create lock
func (l *LedisdbMutex) Lock(key string, timeout int) (bool, error) {
	return l.Server.Lock(key, int32(timeout))
}

// Release release lock
func (l *LedisdbMutex) Release(key string) (bool, error) {
	return l.Server.Release(key)
}
