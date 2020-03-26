package output

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mbrostami/gcron/cron"
	pb "github.com/mbrostami/gcron/grpc"
	"google.golang.org/grpc"
)

// GrpcHandler handles grpc
type GrpcHandler struct {
	connection *grpc.ClientConn
	client     pb.GcronClient
}

// NewGrpcHandler dial connection with rpc server
func NewGrpcHandler(host string, port string) (GrpcHandler, error) {
	var g GrpcHandler
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return g, err
	}
	g.connection = conn
	client := pb.NewGcronClient(g.connection)
	g.client = client
	return g, nil
}

// Initialize send initialized task
func (g GrpcHandler) Initialize(guid string) (bool, error) {
	initialized, err := g.client.InitializeTask(context.Background(), &wrappers.StringValue{Value: guid})
	if err != nil {
		return false, err
	}
	return initialized.GetValue(), nil
}

// Lock try to lock
func (g GrpcHandler) Lock(lockName string) (bool, error) {
	locked, err := g.client.Lock(context.Background(), &wrappers.StringValue{Value: lockName})
	if err != nil {
		return false, err
	}
	return locked.GetValue(), nil
}

// Release try to lock
func (g GrpcHandler) Release(lockName string) (bool, error) {
	released, err := g.client.Release(context.Background(), &wrappers.StringValue{Value: lockName})
	if err != nil {
		return false, err
	}
	return released.GetValue(), nil
}

// Log send string
func (g GrpcHandler) Log(output string) (bool, error) {
	logged, err := g.client.Log(context.Background(), &wrappers.StringValue{Value: output})
	if err != nil {
		return false, err
	}
	return logged.GetValue(), nil
}

// Finish finialize the task
func (g GrpcHandler) Finish(crontask cron.Task) (bool, error) {
	// FIXME find a mapping solution
	startTime, _ := ptypes.TimestampProto(crontask.StartTime)
	endTime, _ := ptypes.TimestampProto(crontask.EndTime)
	grpcTask := &pb.Task{
		FLock:      crontask.FLock,
		FLockName:  crontask.FLockName,
		FOverride:  crontask.FOverride,
		FDelay:     int32(crontask.FDelay),
		Pid:        int32(crontask.Pid),
		GUID:       crontask.GUID,
		UID:        crontask.UID,
		Parent:     crontask.Parent,
		Hostname:   crontask.Hostname,
		Username:   crontask.Username,
		Command:    crontask.Command,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitCode:   int32(crontask.ExitCode),
		Output:     crontask.Output,
		SystemTime: ptypes.DurationProto(crontask.SystemTime),
		UserTime:   ptypes.DurationProto(crontask.UserTime),
		Success:    crontask.Success,
	}
	finished, err := g.client.FinializeTask(context.Background(), grpcTask)
	if err != nil {
		return false, err
	}
	return finished.GetValue(), nil
}

// Close close the connection
func (g GrpcHandler) Close() {
	g.connection.Close()
}
