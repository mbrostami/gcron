package mutex

import (
	"github.com/mbrostami/gcron/cmd/gcron/out"
)

// GrpcMutex to use as mutex lock
type GrpcMutex struct {
	Server *out.GrpcHandler
}

// Lock create lock
func (g *GrpcMutex) Lock(key string, timeout int) (bool, error) {
	return g.Server.Lock(key, int32(timeout))
}

// Release release lock
func (g *GrpcMutex) Release(key string) (bool, error) {
	return g.Server.Release(key)
}
