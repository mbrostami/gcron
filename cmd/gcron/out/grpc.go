package out

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mbrostami/gcron/internal/config"
	pb "github.com/mbrostami/gcron/internal/grpc"
	"github.com/mbrostami/gcron/internal/task"
	"google.golang.org/grpc"
)

// GrpcHandler handles grpc
type GrpcHandler struct {
	connection *grpc.ClientConn
	client     pb.GcronClient
}

// NewHandler make new grpc handler
func NewHandler(cfg config.GeneralConfig) (*GrpcHandler, error) {
	var g GrpcHandler
	host := cfg.GetKey("server.rpc.host").(string)
	port := cfg.GetKey("server.rpc.port").(string)

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return &g, err
	}
	g.connection = conn
	client := pb.NewGcronClient(g.connection)
	g.client = client
	return &g, nil
}

// Lock try to lock
func (g *GrpcHandler) Lock(lockName string, timeout int32) (bool, error) {
	locked, err := g.client.Lock(context.Background(), &pb.LockMessage{Key: lockName, Timeout: timeout})
	if err != nil {
		return false, err
	}
	return locked.GetValue(), nil
}

// Release try to lock
func (g *GrpcHandler) Release(lockName string) (bool, error) {
	released, err := g.client.Release(context.Background(), &wrappers.StringValue{Value: lockName})
	if err != nil {
		return false, err
	}
	return released.GetValue(), nil
}

// StartLogStream start log stream
func (g *GrpcHandler) StartLogStream() (pb.Gcron_StartLogClient, error) {
	stream, err := g.client.StartLog(context.Background())
	if err != nil {
		return stream, err
	}
	return stream, nil
}

// GetLogEntry get logEntry
func (g *GrpcHandler) GetLogEntry(guid string, message []byte) *pb.LogEntry {
	return &pb.LogEntry{GUID: guid, Output: message}
}

// Done finialize the task
func (g *GrpcHandler) Done(crontask task.Task) (bool, error) {
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
		SystemTime: ptypes.DurationProto(crontask.SystemTime),
		UserTime:   ptypes.DurationProto(crontask.UserTime),
		Success:    crontask.Success,
		// Output:     crontask.Output, // output already is streamed
	}
	finished, err := g.client.Done(context.Background(), grpcTask)
	if err != nil {
		return false, err
	}
	return finished.GetValue(), nil
}

// Close close the connection
func (g *GrpcHandler) Close() {
	g.connection.Close()
}
