package db

import pb "github.com/mbrostami/gcron/internal/grpc"

// TaskCollection list of tasks
type TaskCollection struct {
	Tasks map[int]*pb.Task
}

// DB database interface
type DB interface {
	Store(uid uint32, task *pb.Task) (string, error)
	Get(uid uint32, start int, stop int) *TaskCollection
	Close()

	SetTask(task *pb.Task) (bool, error)
	GetTasks(from int32, limit int32) *TaskCollection

	Lock(key string, timeout int32) (bool, error)
	Release(key string) (bool, error)
}
