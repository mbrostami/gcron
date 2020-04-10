package mutex

// Interface lock interface
type Interface interface {
	Lock(key string, timeout int) (bool, error)
	Release(key string) (bool, error)
}
