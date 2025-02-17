// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: edges/sync_service.proto

package edges

import (
	context "context"
	pb "github.com/snple/kokomi/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SyncService_SetDeviceUpdated_FullMethodName    = "/edges.SyncService/SetDeviceUpdated"
	SyncService_GetDeviceUpdated_FullMethodName    = "/edges.SyncService/GetDeviceUpdated"
	SyncService_WaitDeviceUpdated_FullMethodName   = "/edges.SyncService/WaitDeviceUpdated"
	SyncService_SetSlotUpdated_FullMethodName      = "/edges.SyncService/SetSlotUpdated"
	SyncService_GetSlotUpdated_FullMethodName      = "/edges.SyncService/GetSlotUpdated"
	SyncService_SetSourceUpdated_FullMethodName    = "/edges.SyncService/SetSourceUpdated"
	SyncService_GetSourceUpdated_FullMethodName    = "/edges.SyncService/GetSourceUpdated"
	SyncService_SetTagUpdated_FullMethodName       = "/edges.SyncService/SetTagUpdated"
	SyncService_GetTagUpdated_FullMethodName       = "/edges.SyncService/GetTagUpdated"
	SyncService_SetConstUpdated_FullMethodName     = "/edges.SyncService/SetConstUpdated"
	SyncService_GetConstUpdated_FullMethodName     = "/edges.SyncService/GetConstUpdated"
	SyncService_SetTagValueUpdated_FullMethodName  = "/edges.SyncService/SetTagValueUpdated"
	SyncService_GetTagValueUpdated_FullMethodName  = "/edges.SyncService/GetTagValueUpdated"
	SyncService_WaitTagValueUpdated_FullMethodName = "/edges.SyncService/WaitTagValueUpdated"
	SyncService_SetTagWriteUpdated_FullMethodName  = "/edges.SyncService/SetTagWriteUpdated"
	SyncService_GetTagWriteUpdated_FullMethodName  = "/edges.SyncService/GetTagWriteUpdated"
	SyncService_WaitTagWriteUpdated_FullMethodName = "/edges.SyncService/WaitTagWriteUpdated"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	SetDeviceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	WaitDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error)
	SetSlotUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetSlotUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetSourceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetSourceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetTagUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetTagUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetConstUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetConstUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetTagValueUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	WaitTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error)
	SetTagWriteUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetTagWriteUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	WaitTagWriteUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SetDeviceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetDeviceUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetDeviceUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) WaitDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SyncService_ServiceDesc.Streams[0], SyncService_WaitDeviceUpdated_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[pb.MyEmpty, pb.MyBool]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitDeviceUpdatedClient = grpc.ServerStreamingClient[pb.MyBool]

func (c *syncServiceClient) SetSlotUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetSlotUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetSlotUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetSlotUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetSourceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetSourceUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetSourceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetSourceUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetTagUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetTagUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTagUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetTagUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetConstUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetConstUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetConstUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetConstUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetTagValueUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetTagValueUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetTagValueUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) WaitTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SyncService_ServiceDesc.Streams[1], SyncService_WaitTagValueUpdated_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[pb.MyEmpty, pb.MyBool]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitTagValueUpdatedClient = grpc.ServerStreamingClient[pb.MyBool]

func (c *syncServiceClient) SetTagWriteUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetTagWriteUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTagWriteUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetTagWriteUpdated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) WaitTagWriteUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.MyBool], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SyncService_ServiceDesc.Streams[2], SyncService_WaitTagWriteUpdated_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[pb.MyEmpty, pb.MyBool]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitTagWriteUpdatedClient = grpc.ServerStreamingClient[pb.MyBool]

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility.
type SyncServiceServer interface {
	SetDeviceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetDeviceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	WaitDeviceUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error
	SetSlotUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetSlotUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetSourceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetSourceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetTagUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetTagUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetConstUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetConstUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetTagValueUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetTagValueUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	WaitTagValueUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error
	SetTagWriteUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetTagWriteUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	WaitTagWriteUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSyncServiceServer struct{}

func (UnimplementedSyncServiceServer) SetDeviceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetDeviceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) WaitDeviceUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error {
	return status.Errorf(codes.Unimplemented, "method WaitDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetSlotUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSlotUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetSlotUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlotUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetSourceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSourceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetSourceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetTagUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetTagUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetConstUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConstUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetConstUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConstUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetTagValueUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetTagValueUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) WaitTagValueUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error {
	return status.Errorf(codes.Unimplemented, "method WaitTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetTagWriteUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagWriteUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetTagWriteUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagWriteUpdated not implemented")
}
func (UnimplementedSyncServiceServer) WaitTagWriteUpdated(*pb.MyEmpty, grpc.ServerStreamingServer[pb.MyBool]) error {
	return status.Errorf(codes.Unimplemented, "method WaitTagWriteUpdated not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}
func (UnimplementedSyncServiceServer) testEmbeddedByValue()                     {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	// If the following call pancis, it indicates UnimplementedSyncServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_SetDeviceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetDeviceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetDeviceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetDeviceUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetDeviceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetDeviceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetDeviceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetDeviceUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_WaitDeviceUpdated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).WaitDeviceUpdated(m, &grpc.GenericServerStream[pb.MyEmpty, pb.MyBool]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitDeviceUpdatedServer = grpc.ServerStreamingServer[pb.MyBool]

func _SyncService_SetSlotUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetSlotUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetSlotUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetSlotUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetSlotUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetSlotUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetSlotUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetSlotUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetSourceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetSourceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetSourceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetSourceUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetSourceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetSourceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetSourceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetSourceUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetTagUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetTagUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetTagUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetTagUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTagUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTagUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetTagUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTagUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetConstUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetConstUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetConstUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetConstUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetConstUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetConstUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetConstUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetConstUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetTagValueUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetTagValueUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetTagValueUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetTagValueUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTagValueUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTagValueUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetTagValueUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTagValueUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_WaitTagValueUpdated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).WaitTagValueUpdated(m, &grpc.GenericServerStream[pb.MyEmpty, pb.MyBool]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitTagValueUpdatedServer = grpc.ServerStreamingServer[pb.MyBool]

func _SyncService_SetTagWriteUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetTagWriteUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetTagWriteUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetTagWriteUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTagWriteUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTagWriteUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetTagWriteUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTagWriteUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_WaitTagWriteUpdated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).WaitTagWriteUpdated(m, &grpc.GenericServerStream[pb.MyEmpty, pb.MyBool]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SyncService_WaitTagWriteUpdatedServer = grpc.ServerStreamingServer[pb.MyBool]

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edges.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDeviceUpdated",
			Handler:    _SyncService_SetDeviceUpdated_Handler,
		},
		{
			MethodName: "GetDeviceUpdated",
			Handler:    _SyncService_GetDeviceUpdated_Handler,
		},
		{
			MethodName: "SetSlotUpdated",
			Handler:    _SyncService_SetSlotUpdated_Handler,
		},
		{
			MethodName: "GetSlotUpdated",
			Handler:    _SyncService_GetSlotUpdated_Handler,
		},
		{
			MethodName: "SetSourceUpdated",
			Handler:    _SyncService_SetSourceUpdated_Handler,
		},
		{
			MethodName: "GetSourceUpdated",
			Handler:    _SyncService_GetSourceUpdated_Handler,
		},
		{
			MethodName: "SetTagUpdated",
			Handler:    _SyncService_SetTagUpdated_Handler,
		},
		{
			MethodName: "GetTagUpdated",
			Handler:    _SyncService_GetTagUpdated_Handler,
		},
		{
			MethodName: "SetConstUpdated",
			Handler:    _SyncService_SetConstUpdated_Handler,
		},
		{
			MethodName: "GetConstUpdated",
			Handler:    _SyncService_GetConstUpdated_Handler,
		},
		{
			MethodName: "SetTagValueUpdated",
			Handler:    _SyncService_SetTagValueUpdated_Handler,
		},
		{
			MethodName: "GetTagValueUpdated",
			Handler:    _SyncService_GetTagValueUpdated_Handler,
		},
		{
			MethodName: "SetTagWriteUpdated",
			Handler:    _SyncService_SetTagWriteUpdated_Handler,
		},
		{
			MethodName: "GetTagWriteUpdated",
			Handler:    _SyncService_GetTagWriteUpdated_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitDeviceUpdated",
			Handler:       _SyncService_WaitDeviceUpdated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WaitTagValueUpdated",
			Handler:       _SyncService_WaitTagValueUpdated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WaitTagWriteUpdated",
			Handler:       _SyncService_WaitTagWriteUpdated_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "edges/sync_service.proto",
}
