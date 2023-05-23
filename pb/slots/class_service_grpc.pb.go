// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: slots/class_service.proto

package slots

import (
	context "context"
	pb "github.com/snple/kokomi/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClassService_Create_FullMethodName          = "/slots.ClassService/Create"
	ClassService_Update_FullMethodName          = "/slots.ClassService/Update"
	ClassService_View_FullMethodName            = "/slots.ClassService/View"
	ClassService_ViewByName_FullMethodName      = "/slots.ClassService/ViewByName"
	ClassService_Delete_FullMethodName          = "/slots.ClassService/Delete"
	ClassService_List_FullMethodName            = "/slots.ClassService/List"
	ClassService_ViewWithDeleted_FullMethodName = "/slots.ClassService/ViewWithDeleted"
	ClassService_Pull_FullMethodName            = "/slots.ClassService/Pull"
	ClassService_Sync_FullMethodName            = "/slots.ClassService/Sync"
)

// ClassServiceClient is the client API for ClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassServiceClient interface {
	Create(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.Class, error)
	Update(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.Class, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Class, error)
	ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Class, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *ListClassRequest, opts ...grpc.CallOption) (*ListClassResponse, error)
	ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Class, error)
	Pull(ctx context.Context, in *PullClassRequest, opts ...grpc.CallOption) (*PullClassResponse, error)
	Sync(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.MyBool, error)
}

type classServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassServiceClient(cc grpc.ClientConnInterface) ClassServiceClient {
	return &classServiceClient{cc}
}

func (c *classServiceClient) Create(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.Class, error) {
	out := new(pb.Class)
	err := c.cc.Invoke(ctx, ClassService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) Update(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.Class, error) {
	out := new(pb.Class)
	err := c.cc.Invoke(ctx, ClassService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Class, error) {
	out := new(pb.Class)
	err := c.cc.Invoke(ctx, ClassService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Class, error) {
	out := new(pb.Class)
	err := c.cc.Invoke(ctx, ClassService_ViewByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, ClassService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) List(ctx context.Context, in *ListClassRequest, opts ...grpc.CallOption) (*ListClassResponse, error) {
	out := new(ListClassResponse)
	err := c.cc.Invoke(ctx, ClassService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Class, error) {
	out := new(pb.Class)
	err := c.cc.Invoke(ctx, ClassService_ViewWithDeleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) Pull(ctx context.Context, in *PullClassRequest, opts ...grpc.CallOption) (*PullClassResponse, error) {
	out := new(PullClassResponse)
	err := c.cc.Invoke(ctx, ClassService_Pull_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) Sync(ctx context.Context, in *pb.Class, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, ClassService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServiceServer is the server API for ClassService service.
// All implementations must embed UnimplementedClassServiceServer
// for forward compatibility
type ClassServiceServer interface {
	Create(context.Context, *pb.Class) (*pb.Class, error)
	Update(context.Context, *pb.Class) (*pb.Class, error)
	View(context.Context, *pb.Id) (*pb.Class, error)
	ViewByName(context.Context, *pb.Name) (*pb.Class, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *ListClassRequest) (*ListClassResponse, error)
	ViewWithDeleted(context.Context, *pb.Id) (*pb.Class, error)
	Pull(context.Context, *PullClassRequest) (*PullClassResponse, error)
	Sync(context.Context, *pb.Class) (*pb.MyBool, error)
	mustEmbedUnimplementedClassServiceServer()
}

// UnimplementedClassServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClassServiceServer struct {
}

func (UnimplementedClassServiceServer) Create(context.Context, *pb.Class) (*pb.Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClassServiceServer) Update(context.Context, *pb.Class) (*pb.Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClassServiceServer) View(context.Context, *pb.Id) (*pb.Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedClassServiceServer) ViewByName(context.Context, *pb.Name) (*pb.Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewByName not implemented")
}
func (UnimplementedClassServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClassServiceServer) List(context.Context, *ListClassRequest) (*ListClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClassServiceServer) ViewWithDeleted(context.Context, *pb.Id) (*pb.Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedClassServiceServer) Pull(context.Context, *PullClassRequest) (*PullClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedClassServiceServer) Sync(context.Context, *pb.Class) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedClassServiceServer) mustEmbedUnimplementedClassServiceServer() {}

// UnsafeClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServiceServer will
// result in compilation errors.
type UnsafeClassServiceServer interface {
	mustEmbedUnimplementedClassServiceServer()
}

func RegisterClassServiceServer(s grpc.ServiceRegistrar, srv ClassServiceServer) {
	s.RegisterService(&ClassService_ServiceDesc, srv)
}

func _ClassService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).Create(ctx, req.(*pb.Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).Update(ctx, req.(*pb.Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_ViewByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).ViewByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_ViewByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).ViewByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).List(ctx, req.(*ListClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).ViewWithDeleted(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).Pull(ctx, req.(*PullClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).Sync(ctx, req.(*pb.Class))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassService_ServiceDesc is the grpc.ServiceDesc for ClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slots.ClassService",
	HandlerType: (*ClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClassService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClassService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _ClassService_View_Handler,
		},
		{
			MethodName: "ViewByName",
			Handler:    _ClassService_ViewByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClassService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClassService_List_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _ClassService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _ClassService_Pull_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _ClassService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slots/class_service.proto",
}

const (
	AttrService_Create_FullMethodName                  = "/slots.AttrService/Create"
	AttrService_Update_FullMethodName                  = "/slots.AttrService/Update"
	AttrService_View_FullMethodName                    = "/slots.AttrService/View"
	AttrService_ViewByName_FullMethodName              = "/slots.AttrService/ViewByName"
	AttrService_Delete_FullMethodName                  = "/slots.AttrService/Delete"
	AttrService_List_FullMethodName                    = "/slots.AttrService/List"
	AttrService_GetValue_FullMethodName                = "/slots.AttrService/GetValue"
	AttrService_SetValue_FullMethodName                = "/slots.AttrService/SetValue"
	AttrService_SetValueUnchecked_FullMethodName       = "/slots.AttrService/SetValueUnchecked"
	AttrService_GetValueByName_FullMethodName          = "/slots.AttrService/GetValueByName"
	AttrService_SetValueByName_FullMethodName          = "/slots.AttrService/SetValueByName"
	AttrService_SetValueByNameUnchecked_FullMethodName = "/slots.AttrService/SetValueByNameUnchecked"
	AttrService_ViewWithDeleted_FullMethodName         = "/slots.AttrService/ViewWithDeleted"
	AttrService_Pull_FullMethodName                    = "/slots.AttrService/Pull"
	AttrService_Sync_FullMethodName                    = "/slots.AttrService/Sync"
)

// AttrServiceClient is the client API for AttrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttrServiceClient interface {
	Create(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.Attr, error)
	Update(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.Attr, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Attr, error)
	ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Attr, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *ListAttrRequest, opts ...grpc.CallOption) (*ListAttrResponse, error)
	GetValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.AttrValue, error)
	SetValue(ctx context.Context, in *pb.AttrValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	SetValueUnchecked(ctx context.Context, in *pb.AttrValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetValueByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.AttrNameValue, error)
	SetValueByName(ctx context.Context, in *pb.AttrNameValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	SetValueByNameUnchecked(ctx context.Context, in *pb.AttrNameValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Attr, error)
	Pull(ctx context.Context, in *PullAttrRequest, opts ...grpc.CallOption) (*PullAttrResponse, error)
	Sync(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.MyBool, error)
}

type attrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttrServiceClient(cc grpc.ClientConnInterface) AttrServiceClient {
	return &attrServiceClient{cc}
}

func (c *attrServiceClient) Create(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.Attr, error) {
	out := new(pb.Attr)
	err := c.cc.Invoke(ctx, AttrService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) Update(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.Attr, error) {
	out := new(pb.Attr)
	err := c.cc.Invoke(ctx, AttrService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Attr, error) {
	out := new(pb.Attr)
	err := c.cc.Invoke(ctx, AttrService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Attr, error) {
	out := new(pb.Attr)
	err := c.cc.Invoke(ctx, AttrService_ViewByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) List(ctx context.Context, in *ListAttrRequest, opts ...grpc.CallOption) (*ListAttrResponse, error) {
	out := new(ListAttrResponse)
	err := c.cc.Invoke(ctx, AttrService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) GetValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.AttrValue, error) {
	out := new(pb.AttrValue)
	err := c.cc.Invoke(ctx, AttrService_GetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) SetValue(ctx context.Context, in *pb.AttrValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_SetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) SetValueUnchecked(ctx context.Context, in *pb.AttrValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_SetValueUnchecked_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) GetValueByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.AttrNameValue, error) {
	out := new(pb.AttrNameValue)
	err := c.cc.Invoke(ctx, AttrService_GetValueByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) SetValueByName(ctx context.Context, in *pb.AttrNameValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_SetValueByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) SetValueByNameUnchecked(ctx context.Context, in *pb.AttrNameValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_SetValueByNameUnchecked_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Attr, error) {
	out := new(pb.Attr)
	err := c.cc.Invoke(ctx, AttrService_ViewWithDeleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) Pull(ctx context.Context, in *PullAttrRequest, opts ...grpc.CallOption) (*PullAttrResponse, error) {
	out := new(PullAttrResponse)
	err := c.cc.Invoke(ctx, AttrService_Pull_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrServiceClient) Sync(ctx context.Context, in *pb.Attr, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, AttrService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttrServiceServer is the server API for AttrService service.
// All implementations must embed UnimplementedAttrServiceServer
// for forward compatibility
type AttrServiceServer interface {
	Create(context.Context, *pb.Attr) (*pb.Attr, error)
	Update(context.Context, *pb.Attr) (*pb.Attr, error)
	View(context.Context, *pb.Id) (*pb.Attr, error)
	ViewByName(context.Context, *pb.Name) (*pb.Attr, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *ListAttrRequest) (*ListAttrResponse, error)
	GetValue(context.Context, *pb.Id) (*pb.AttrValue, error)
	SetValue(context.Context, *pb.AttrValue) (*pb.MyBool, error)
	SetValueUnchecked(context.Context, *pb.AttrValue) (*pb.MyBool, error)
	GetValueByName(context.Context, *pb.Name) (*pb.AttrNameValue, error)
	SetValueByName(context.Context, *pb.AttrNameValue) (*pb.MyBool, error)
	SetValueByNameUnchecked(context.Context, *pb.AttrNameValue) (*pb.MyBool, error)
	ViewWithDeleted(context.Context, *pb.Id) (*pb.Attr, error)
	Pull(context.Context, *PullAttrRequest) (*PullAttrResponse, error)
	Sync(context.Context, *pb.Attr) (*pb.MyBool, error)
	mustEmbedUnimplementedAttrServiceServer()
}

// UnimplementedAttrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttrServiceServer struct {
}

func (UnimplementedAttrServiceServer) Create(context.Context, *pb.Attr) (*pb.Attr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttrServiceServer) Update(context.Context, *pb.Attr) (*pb.Attr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAttrServiceServer) View(context.Context, *pb.Id) (*pb.Attr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedAttrServiceServer) ViewByName(context.Context, *pb.Name) (*pb.Attr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewByName not implemented")
}
func (UnimplementedAttrServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAttrServiceServer) List(context.Context, *ListAttrRequest) (*ListAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAttrServiceServer) GetValue(context.Context, *pb.Id) (*pb.AttrValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedAttrServiceServer) SetValue(context.Context, *pb.AttrValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedAttrServiceServer) SetValueUnchecked(context.Context, *pb.AttrValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueUnchecked not implemented")
}
func (UnimplementedAttrServiceServer) GetValueByName(context.Context, *pb.Name) (*pb.AttrNameValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueByName not implemented")
}
func (UnimplementedAttrServiceServer) SetValueByName(context.Context, *pb.AttrNameValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueByName not implemented")
}
func (UnimplementedAttrServiceServer) SetValueByNameUnchecked(context.Context, *pb.AttrNameValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueByNameUnchecked not implemented")
}
func (UnimplementedAttrServiceServer) ViewWithDeleted(context.Context, *pb.Id) (*pb.Attr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedAttrServiceServer) Pull(context.Context, *PullAttrRequest) (*PullAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedAttrServiceServer) Sync(context.Context, *pb.Attr) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedAttrServiceServer) mustEmbedUnimplementedAttrServiceServer() {}

// UnsafeAttrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttrServiceServer will
// result in compilation errors.
type UnsafeAttrServiceServer interface {
	mustEmbedUnimplementedAttrServiceServer()
}

func RegisterAttrServiceServer(s grpc.ServiceRegistrar, srv AttrServiceServer) {
	s.RegisterService(&AttrService_ServiceDesc, srv)
}

func _AttrService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Attr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).Create(ctx, req.(*pb.Attr))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Attr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).Update(ctx, req.(*pb.Attr))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_ViewByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).ViewByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_ViewByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).ViewByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).List(ctx, req.(*ListAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).GetValue(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.AttrValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_SetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).SetValue(ctx, req.(*pb.AttrValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_SetValueUnchecked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.AttrValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).SetValueUnchecked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_SetValueUnchecked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).SetValueUnchecked(ctx, req.(*pb.AttrValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_GetValueByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).GetValueByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_GetValueByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).GetValueByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_SetValueByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.AttrNameValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).SetValueByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_SetValueByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).SetValueByName(ctx, req.(*pb.AttrNameValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_SetValueByNameUnchecked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.AttrNameValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).SetValueByNameUnchecked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_SetValueByNameUnchecked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).SetValueByNameUnchecked(ctx, req.(*pb.AttrNameValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).ViewWithDeleted(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).Pull(ctx, req.(*PullAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttrService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Attr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttrServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttrService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttrServiceServer).Sync(ctx, req.(*pb.Attr))
	}
	return interceptor(ctx, in, info, handler)
}

// AttrService_ServiceDesc is the grpc.ServiceDesc for AttrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slots.AttrService",
	HandlerType: (*AttrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttrService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AttrService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _AttrService_View_Handler,
		},
		{
			MethodName: "ViewByName",
			Handler:    _AttrService_ViewByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AttrService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AttrService_List_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _AttrService_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _AttrService_SetValue_Handler,
		},
		{
			MethodName: "SetValueUnchecked",
			Handler:    _AttrService_SetValueUnchecked_Handler,
		},
		{
			MethodName: "GetValueByName",
			Handler:    _AttrService_GetValueByName_Handler,
		},
		{
			MethodName: "SetValueByName",
			Handler:    _AttrService_SetValueByName_Handler,
		},
		{
			MethodName: "SetValueByNameUnchecked",
			Handler:    _AttrService_SetValueByNameUnchecked_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _AttrService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _AttrService_Pull_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _AttrService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slots/class_service.proto",
}
