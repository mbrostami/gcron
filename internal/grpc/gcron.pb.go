// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gcron.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LogEntry struct {
	GUID                 string   `protobuf:"bytes,1,opt,name=GUID,proto3" json:"GUID,omitempty"`
	Output               []byte   `protobuf:"bytes,2,opt,name=Output,proto3" json:"Output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd548e04c98879c7, []int{0}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetGUID() string {
	if m != nil {
		return m.GUID
	}
	return ""
}

func (m *LogEntry) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type LockMessage struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Timeout              int32    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockMessage) Reset()         { *m = LockMessage{} }
func (m *LockMessage) String() string { return proto.CompactTextString(m) }
func (*LockMessage) ProtoMessage()    {}
func (*LockMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd548e04c98879c7, []int{1}
}

func (m *LockMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockMessage.Unmarshal(m, b)
}
func (m *LockMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockMessage.Marshal(b, m, deterministic)
}
func (m *LockMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockMessage.Merge(m, src)
}
func (m *LockMessage) XXX_Size() int {
	return xxx_messageInfo_LockMessage.Size(m)
}
func (m *LockMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LockMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LockMessage proto.InternalMessageInfo

func (m *LockMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LockMessage) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type Task struct {
	FLock                bool                 `protobuf:"varint,1,opt,name=FLock,proto3" json:"FLock,omitempty"`
	FLockName            string               `protobuf:"bytes,2,opt,name=FLockName,proto3" json:"FLockName,omitempty"`
	FOverride            string               `protobuf:"bytes,3,opt,name=FOverride,proto3" json:"FOverride,omitempty"`
	FDelay               int32                `protobuf:"varint,4,opt,name=FDelay,proto3" json:"FDelay,omitempty"`
	Pid                  int32                `protobuf:"varint,5,opt,name=Pid,proto3" json:"Pid,omitempty"`
	GUID                 string               `protobuf:"bytes,6,opt,name=GUID,proto3" json:"GUID,omitempty"`
	UID                  uint32               `protobuf:"varint,7,opt,name=UID,proto3" json:"UID,omitempty"`
	Parent               string               `protobuf:"bytes,8,opt,name=Parent,proto3" json:"Parent,omitempty"`
	Hostname             string               `protobuf:"bytes,9,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	Username             string               `protobuf:"bytes,10,opt,name=Username,proto3" json:"Username,omitempty"`
	Command              string               `protobuf:"bytes,11,opt,name=Command,proto3" json:"Command,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,12,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,13,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	ExitCode             int32                `protobuf:"varint,14,opt,name=ExitCode,proto3" json:"ExitCode,omitempty"`
	Output               []byte               `protobuf:"bytes,15,opt,name=Output,proto3" json:"Output,omitempty"`
	SystemTime           *duration.Duration   `protobuf:"bytes,16,opt,name=SystemTime,proto3" json:"SystemTime,omitempty"`
	UserTime             *duration.Duration   `protobuf:"bytes,17,opt,name=UserTime,proto3" json:"UserTime,omitempty"`
	Success              bool                 `protobuf:"varint,18,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd548e04c98879c7, []int{2}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetFLock() bool {
	if m != nil {
		return m.FLock
	}
	return false
}

func (m *Task) GetFLockName() string {
	if m != nil {
		return m.FLockName
	}
	return ""
}

func (m *Task) GetFOverride() string {
	if m != nil {
		return m.FOverride
	}
	return ""
}

func (m *Task) GetFDelay() int32 {
	if m != nil {
		return m.FDelay
	}
	return 0
}

func (m *Task) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Task) GetGUID() string {
	if m != nil {
		return m.GUID
	}
	return ""
}

func (m *Task) GetUID() uint32 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *Task) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Task) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Task) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Task) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Task) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Task) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Task) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *Task) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Task) GetSystemTime() *duration.Duration {
	if m != nil {
		return m.SystemTime
	}
	return nil
}

func (m *Task) GetUserTime() *duration.Duration {
	if m != nil {
		return m.UserTime
	}
	return nil
}

func (m *Task) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "grpc.LogEntry")
	proto.RegisterType((*LockMessage)(nil), "grpc.LockMessage")
	proto.RegisterType((*Task)(nil), "grpc.Task")
}

func init() {
	proto.RegisterFile("gcron.proto", fileDescriptor_fd548e04c98879c7)
}

var fileDescriptor_fd548e04c98879c7 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xc4, 0xf9, 0x9a, 0xb4, 0x25, 0x5d, 0x21, 0xb4, 0x44, 0x15, 0x44, 0x39, 0xe5, 0xe4,
	0xa2, 0xf2, 0x59, 0x8e, 0x34, 0x29, 0x20, 0x05, 0x5a, 0x39, 0x29, 0xf7, 0x6d, 0x3c, 0x58, 0x56,
	0x62, 0xaf, 0xb5, 0xbb, 0x06, 0xf2, 0xf7, 0xf8, 0x4f, 0xdc, 0xd1, 0x8e, 0xbd, 0xa9, 0xd5, 0x48,
	0xcd, 0x6d, 0xdf, 0xbc, 0xf7, 0x76, 0x66, 0x67, 0x76, 0xa0, 0x17, 0x2f, 0x95, 0xcc, 0x82, 0x5c,
	0x49, 0x23, 0x99, 0x1f, 0xab, 0x7c, 0x39, 0x78, 0x11, 0x4b, 0x19, 0xaf, 0xf1, 0x94, 0x62, 0xb7,
	0xc5, 0xcf, 0xd3, 0xdf, 0x4a, 0xe4, 0x39, 0x2a, 0x5d, 0xaa, 0x06, 0x2f, 0xef, 0xf3, 0x26, 0x49,
	0x51, 0x1b, 0x91, 0xe6, 0x95, 0x60, 0xe7, 0x82, 0xa8, 0x50, 0xc2, 0x24, 0x2e, 0xcd, 0xe8, 0x1d,
	0x74, 0x66, 0x32, 0x9e, 0x66, 0x46, 0x6d, 0x18, 0x03, 0xff, 0xf3, 0xcd, 0xd7, 0x09, 0xf7, 0x86,
	0xde, 0xb8, 0x1b, 0xd2, 0x99, 0x3d, 0x83, 0xd6, 0x55, 0x61, 0xf2, 0xc2, 0xf0, 0xc7, 0x43, 0x6f,
	0x7c, 0x10, 0x56, 0x68, 0x74, 0x0e, 0xbd, 0x99, 0x5c, 0xae, 0xbe, 0xa1, 0xd6, 0x22, 0x46, 0xd6,
	0x87, 0xc6, 0x0a, 0x37, 0x95, 0xd3, 0x1e, 0x19, 0x87, 0xb6, 0xad, 0x45, 0x56, 0xce, 0x66, 0xe8,
	0xe0, 0xe8, 0xaf, 0x0f, 0xfe, 0x42, 0xe8, 0x15, 0x7b, 0x0a, 0xcd, 0x4b, 0x7b, 0x09, 0xd9, 0x3a,
	0x61, 0x09, 0xd8, 0x09, 0x74, 0xe9, 0xf0, 0x5d, 0xa4, 0x48, 0xd6, 0x6e, 0x78, 0x17, 0x20, 0xf6,
	0xea, 0x17, 0x2a, 0x95, 0x44, 0xc8, 0x1b, 0x15, 0xeb, 0x02, 0xb6, 0xda, 0xcb, 0x09, 0xae, 0xc5,
	0x86, 0xfb, 0x94, 0xb3, 0x42, 0xb6, 0xbc, 0xeb, 0x24, 0xe2, 0x4d, 0x0a, 0xda, 0xe3, 0xf6, 0xad,
	0xad, 0xda, 0x5b, 0xfb, 0xd0, 0xb0, 0xa1, 0xf6, 0xd0, 0x1b, 0x1f, 0x86, 0x8d, 0xea, 0xf5, 0xd7,
	0x42, 0x61, 0x66, 0x78, 0x87, 0x74, 0x15, 0x62, 0x03, 0xe8, 0x7c, 0x91, 0xda, 0x64, 0xb6, 0xc4,
	0x2e, 0x31, 0x5b, 0x6c, 0xb9, 0x1b, 0x8d, 0x8a, 0x38, 0x28, 0x39, 0x87, 0x6d, 0x53, 0x2e, 0x64,
	0x9a, 0x8a, 0x2c, 0xe2, 0x3d, 0xa2, 0x1c, 0x64, 0x1f, 0xa0, 0x3b, 0x37, 0x42, 0x99, 0x45, 0x92,
	0x22, 0x3f, 0x18, 0x7a, 0xe3, 0xde, 0xd9, 0x20, 0x28, 0x67, 0x17, 0xb8, 0xd9, 0x05, 0x0b, 0x37,
	0xdc, 0xf0, 0x4e, 0xcc, 0xde, 0x40, 0x7b, 0x9a, 0x45, 0xe4, 0x3b, 0xdc, 0xeb, 0x73, 0x52, 0x5b,
	0xe5, 0xf4, 0x4f, 0x62, 0x2e, 0x64, 0x84, 0xfc, 0x88, 0xda, 0xb2, 0xc5, 0xb5, 0x99, 0x3f, 0xa9,
	0xcf, 0x9c, 0x9d, 0x03, 0xcc, 0x37, 0xda, 0x60, 0x4a, 0xc9, 0xfa, 0x94, 0xec, 0xf9, 0x4e, 0xb2,
	0x49, 0xf5, 0xc1, 0xc2, 0x9a, 0x98, 0xbd, 0x2d, 0x9b, 0x42, 0xc6, 0xe3, 0x7d, 0xc6, 0xad, 0xd4,
	0xf6, 0x6b, 0x5e, 0x2c, 0x97, 0xa8, 0x35, 0x67, 0xf4, 0x47, 0x1c, 0x3c, 0xfb, 0xe7, 0x41, 0x93,
	0xd6, 0x85, 0xbd, 0x07, 0x9f, 0xfe, 0xcd, 0x71, 0x60, 0x37, 0x26, 0xa8, 0xfd, 0xca, 0xc1, 0x6e,
	0x27, 0x3e, 0x49, 0xb9, 0xfe, 0x21, 0xd6, 0x05, 0x8e, 0x1e, 0xb1, 0x29, 0xb4, 0x43, 0x5c, 0xa3,
	0xd0, 0xc8, 0x4e, 0x76, 0x84, 0x73, 0xa3, 0x92, 0x2c, 0x26, 0xe9, 0x9e, 0x6b, 0x3e, 0x42, 0x87,
	0x86, 0x31, 0x93, 0x31, 0x3b, 0x72, 0x35, 0x94, 0x1b, 0xf5, 0xb0, 0x73, 0xec, 0xb1, 0x57, 0xe0,
	0x4f, 0x64, 0x86, 0x0c, 0x4a, 0x9f, 0xdd, 0x8a, 0x87, 0x3d, 0xb7, 0x2d, 0x8a, 0xbe, 0xfe, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0xa3, 0x17, 0x9c, 0x2c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GcronClient is the client API for Gcron service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GcronClient interface {
	// Lock remote lock based on (UID) or Custom Lock Name
	Lock(ctx context.Context, in *LockMessage, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	// Release release lock based on (HOSTNAME+USERNAME+UID) or Custom Lock Name
	Release(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	StartLog(ctx context.Context, opts ...grpc.CallOption) (Gcron_StartLogClient, error)
	Done(ctx context.Context, in *Task, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type gcronClient struct {
	cc grpc.ClientConnInterface
}

func NewGcronClient(cc grpc.ClientConnInterface) GcronClient {
	return &gcronClient{cc}
}

func (c *gcronClient) Lock(ctx context.Context, in *LockMessage, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/grpc.gcron/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcronClient) Release(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/grpc.gcron/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcronClient) StartLog(ctx context.Context, opts ...grpc.CallOption) (Gcron_StartLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gcron_serviceDesc.Streams[0], "/grpc.gcron/StartLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &gcronStartLogClient{stream}
	return x, nil
}

type Gcron_StartLogClient interface {
	Send(*LogEntry) error
	CloseAndRecv() (*wrappers.BoolValue, error)
	grpc.ClientStream
}

type gcronStartLogClient struct {
	grpc.ClientStream
}

func (x *gcronStartLogClient) Send(m *LogEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gcronStartLogClient) CloseAndRecv() (*wrappers.BoolValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrappers.BoolValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gcronClient) Done(ctx context.Context, in *Task, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/grpc.gcron/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcronServer is the server API for Gcron service.
type GcronServer interface {
	// Lock remote lock based on (UID) or Custom Lock Name
	Lock(context.Context, *LockMessage) (*wrappers.BoolValue, error)
	// Release release lock based on (HOSTNAME+USERNAME+UID) or Custom Lock Name
	Release(context.Context, *wrappers.StringValue) (*wrappers.BoolValue, error)
	StartLog(Gcron_StartLogServer) error
	Done(context.Context, *Task) (*wrappers.BoolValue, error)
}

// UnimplementedGcronServer can be embedded to have forward compatible implementations.
type UnimplementedGcronServer struct {
}

func (*UnimplementedGcronServer) Lock(ctx context.Context, req *LockMessage) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (*UnimplementedGcronServer) Release(ctx context.Context, req *wrappers.StringValue) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (*UnimplementedGcronServer) StartLog(srv Gcron_StartLogServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLog not implemented")
}
func (*UnimplementedGcronServer) Done(ctx context.Context, req *Task) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}

func RegisterGcronServer(s *grpc.Server, srv GcronServer) {
	s.RegisterService(&_Gcron_serviceDesc, srv)
}

func _Gcron_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcronServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gcron/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcronServer).Lock(ctx, req.(*LockMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gcron_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcronServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gcron/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcronServer).Release(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gcron_StartLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GcronServer).StartLog(&gcronStartLogServer{stream})
}

type Gcron_StartLogServer interface {
	SendAndClose(*wrappers.BoolValue) error
	Recv() (*LogEntry, error)
	grpc.ServerStream
}

type gcronStartLogServer struct {
	grpc.ServerStream
}

func (x *gcronStartLogServer) SendAndClose(m *wrappers.BoolValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gcronStartLogServer) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gcron_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcronServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gcron/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcronServer).Done(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gcron_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gcron",
	HandlerType: (*GcronServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _Gcron_Lock_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _Gcron_Release_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Gcron_Done_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLog",
			Handler:       _Gcron_StartLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "gcron.proto",
}