package output

import (
	"context"
	"time"

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
	stream     pb.Gcron_StartLogClient
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

// StartLogStream start log stream
func (g GrpcHandler) StartLogStream() (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream, err := g.client.StartLog(ctx)
	if err != nil {
		return false, err
	}
	g.stream = stream
	return true, nil
}

// Log send log over stream
func (g GrpcHandler) Log(guid string, message []byte) error {
	return g.stream.Send(&pb.LogEntry{GUID: guid, Output: message})
}

// CloseStream close stream
func (g GrpcHandler) CloseStream() (bool, error) {
	boolValue, err := g.stream.CloseAndRecv()
	return boolValue.GetValue(), err
}

// Done finialize the task
func (g GrpcHandler) Done(crontask cron.Task) (bool, error) {
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
	finished, err := g.client.Done(context.Background(), grpcTask)
	if err != nil {
		return false, err
	}
	return finished.GetValue(), nil
}

// Close close the connection
func (g GrpcHandler) Close() {
	g.connection.Close()
}
