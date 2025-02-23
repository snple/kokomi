// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: cores/slot_service.proto

package cores

import (
	context "context"
	pb "github.com/snple/beacon/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SlotService_Create_FullMethodName          = "/cores.SlotService/Create"
	SlotService_Update_FullMethodName          = "/cores.SlotService/Update"
	SlotService_View_FullMethodName            = "/cores.SlotService/View"
	SlotService_Name_FullMethodName            = "/cores.SlotService/Name"
	SlotService_NameFull_FullMethodName        = "/cores.SlotService/NameFull"
	SlotService_Delete_FullMethodName          = "/cores.SlotService/Delete"
	SlotService_List_FullMethodName            = "/cores.SlotService/List"
	SlotService_Link_FullMethodName            = "/cores.SlotService/Link"
	SlotService_Clone_FullMethodName           = "/cores.SlotService/Clone"
	SlotService_ViewWithDeleted_FullMethodName = "/cores.SlotService/ViewWithDeleted"
	SlotService_Pull_FullMethodName            = "/cores.SlotService/Pull"
	SlotService_Sync_FullMethodName            = "/cores.SlotService/Sync"
)

// SlotServiceClient is the client API for SlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlotServiceClient interface {
	Create(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error)
	Update(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error)
	Name(ctx context.Context, in *SlotNameRequest, opts ...grpc.CallOption) (*pb.Slot, error)
	NameFull(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Slot, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *SlotListRequest, opts ...grpc.CallOption) (*SlotListResponse, error)
	Link(ctx context.Context, in *SlotLinkRequest, opts ...grpc.CallOption) (*pb.MyBool, error)
	Clone(ctx context.Context, in *SlotCloneRequest, opts ...grpc.CallOption) (*pb.MyBool, error)
	ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error)
	Pull(ctx context.Context, in *SlotPullRequest, opts ...grpc.CallOption) (*SlotPullResponse, error)
	Sync(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.MyBool, error)
}

type slotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSlotServiceClient(cc grpc.ClientConnInterface) SlotServiceClient {
	return &slotServiceClient{cc}
}

func (c *slotServiceClient) Create(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Update(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_View_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Name(ctx context.Context, in *SlotNameRequest, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_Name_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) NameFull(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_NameFull_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) List(ctx context.Context, in *SlotListRequest, opts ...grpc.CallOption) (*SlotListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SlotListResponse)
	err := c.cc.Invoke(ctx, SlotService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Link(ctx context.Context, in *SlotLinkRequest, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Link_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Clone(ctx context.Context, in *SlotCloneRequest, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Clone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_ViewWithDeleted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Pull(ctx context.Context, in *SlotPullRequest, opts ...grpc.CallOption) (*SlotPullResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SlotPullResponse)
	err := c.cc.Invoke(ctx, SlotService_Pull_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Sync(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.MyBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlotServiceServer is the server API for SlotService service.
// All implementations must embed UnimplementedSlotServiceServer
// for forward compatibility.
type SlotServiceServer interface {
	Create(context.Context, *pb.Slot) (*pb.Slot, error)
	Update(context.Context, *pb.Slot) (*pb.Slot, error)
	View(context.Context, *pb.Id) (*pb.Slot, error)
	Name(context.Context, *SlotNameRequest) (*pb.Slot, error)
	NameFull(context.Context, *pb.Name) (*pb.Slot, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *SlotListRequest) (*SlotListResponse, error)
	Link(context.Context, *SlotLinkRequest) (*pb.MyBool, error)
	Clone(context.Context, *SlotCloneRequest) (*pb.MyBool, error)
	ViewWithDeleted(context.Context, *pb.Id) (*pb.Slot, error)
	Pull(context.Context, *SlotPullRequest) (*SlotPullResponse, error)
	Sync(context.Context, *pb.Slot) (*pb.MyBool, error)
	mustEmbedUnimplementedSlotServiceServer()
}

// UnimplementedSlotServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSlotServiceServer struct{}

func (UnimplementedSlotServiceServer) Create(context.Context, *pb.Slot) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSlotServiceServer) Update(context.Context, *pb.Slot) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSlotServiceServer) View(context.Context, *pb.Id) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedSlotServiceServer) Name(context.Context, *SlotNameRequest) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedSlotServiceServer) NameFull(context.Context, *pb.Name) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NameFull not implemented")
}
func (UnimplementedSlotServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSlotServiceServer) List(context.Context, *SlotListRequest) (*SlotListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSlotServiceServer) Link(context.Context, *SlotLinkRequest) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (UnimplementedSlotServiceServer) Clone(context.Context, *SlotCloneRequest) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedSlotServiceServer) ViewWithDeleted(context.Context, *pb.Id) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedSlotServiceServer) Pull(context.Context, *SlotPullRequest) (*SlotPullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedSlotServiceServer) Sync(context.Context, *pb.Slot) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedSlotServiceServer) mustEmbedUnimplementedSlotServiceServer() {}
func (UnimplementedSlotServiceServer) testEmbeddedByValue()                     {}

// UnsafeSlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlotServiceServer will
// result in compilation errors.
type UnsafeSlotServiceServer interface {
	mustEmbedUnimplementedSlotServiceServer()
}

func RegisterSlotServiceServer(s grpc.ServiceRegistrar, srv SlotServiceServer) {
	// If the following call pancis, it indicates UnimplementedSlotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SlotService_ServiceDesc, srv)
}

func _SlotService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Create(ctx, req.(*pb.Slot))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Update(ctx, req.(*pb.Slot))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Name_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Name(ctx, req.(*SlotNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_NameFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).NameFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_NameFull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).NameFull(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).List(ctx, req.(*SlotListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Link_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Link(ctx, req.(*SlotLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotCloneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Clone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Clone(ctx, req.(*SlotCloneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).ViewWithDeleted(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotPullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Pull(ctx, req.(*SlotPullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Sync(ctx, req.(*pb.Slot))
	}
	return interceptor(ctx, in, info, handler)
}

// SlotService_ServiceDesc is the grpc.ServiceDesc for SlotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cores.SlotService",
	HandlerType: (*SlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SlotService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SlotService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _SlotService_View_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _SlotService_Name_Handler,
		},
		{
			MethodName: "NameFull",
			Handler:    _SlotService_NameFull_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SlotService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SlotService_List_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _SlotService_Link_Handler,
		},
		{
			MethodName: "Clone",
			Handler:    _SlotService_Clone_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _SlotService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _SlotService_Pull_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _SlotService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cores/slot_service.proto",
}
