// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: nodes/var_service.proto

package nodes

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
	VarService_Create_FullMethodName                  = "/nodes.VarService/Create"
	VarService_Update_FullMethodName                  = "/nodes.VarService/Update"
	VarService_View_FullMethodName                    = "/nodes.VarService/View"
	VarService_ViewByName_FullMethodName              = "/nodes.VarService/ViewByName"
	VarService_Delete_FullMethodName                  = "/nodes.VarService/Delete"
	VarService_List_FullMethodName                    = "/nodes.VarService/List"
	VarService_GetValue_FullMethodName                = "/nodes.VarService/GetValue"
	VarService_SetValue_FullMethodName                = "/nodes.VarService/SetValue"
	VarService_SetValueUnchecked_FullMethodName       = "/nodes.VarService/SetValueUnchecked"
	VarService_GetValueByName_FullMethodName          = "/nodes.VarService/GetValueByName"
	VarService_SetValueByName_FullMethodName          = "/nodes.VarService/SetValueByName"
	VarService_SetValueByNameUnchecked_FullMethodName = "/nodes.VarService/SetValueByNameUnchecked"
	VarService_ViewWithDeleted_FullMethodName         = "/nodes.VarService/ViewWithDeleted"
	VarService_Pull_FullMethodName                    = "/nodes.VarService/Pull"
	VarService_Sync_FullMethodName                    = "/nodes.VarService/Sync"
)

// VarServiceClient is the client API for VarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VarServiceClient interface {
	Create(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.Var, error)
	Update(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.Var, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Var, error)
	ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Var, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *ListVarRequest, opts ...grpc.CallOption) (*ListVarResponse, error)
	GetValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.VarValue, error)
	SetValue(ctx context.Context, in *pb.VarValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	SetValueUnchecked(ctx context.Context, in *pb.VarValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetValueByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.VarNameValue, error)
	SetValueByName(ctx context.Context, in *pb.VarNameValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	SetValueByNameUnchecked(ctx context.Context, in *pb.VarNameValue, opts ...grpc.CallOption) (*pb.MyBool, error)
	ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Var, error)
	Pull(ctx context.Context, in *PullVarRequest, opts ...grpc.CallOption) (*PullVarResponse, error)
	Sync(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.MyBool, error)
}

type varServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVarServiceClient(cc grpc.ClientConnInterface) VarServiceClient {
	return &varServiceClient{cc}
}

func (c *varServiceClient) Create(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.Var, error) {
	out := new(pb.Var)
	err := c.cc.Invoke(ctx, VarService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) Update(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.Var, error) {
	out := new(pb.Var)
	err := c.cc.Invoke(ctx, VarService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Var, error) {
	out := new(pb.Var)
	err := c.cc.Invoke(ctx, VarService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Var, error) {
	out := new(pb.Var)
	err := c.cc.Invoke(ctx, VarService_ViewByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) List(ctx context.Context, in *ListVarRequest, opts ...grpc.CallOption) (*ListVarResponse, error) {
	out := new(ListVarResponse)
	err := c.cc.Invoke(ctx, VarService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) GetValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.VarValue, error) {
	out := new(pb.VarValue)
	err := c.cc.Invoke(ctx, VarService_GetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) SetValue(ctx context.Context, in *pb.VarValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_SetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) SetValueUnchecked(ctx context.Context, in *pb.VarValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_SetValueUnchecked_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) GetValueByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.VarNameValue, error) {
	out := new(pb.VarNameValue)
	err := c.cc.Invoke(ctx, VarService_GetValueByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) SetValueByName(ctx context.Context, in *pb.VarNameValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_SetValueByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) SetValueByNameUnchecked(ctx context.Context, in *pb.VarNameValue, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_SetValueByNameUnchecked_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Var, error) {
	out := new(pb.Var)
	err := c.cc.Invoke(ctx, VarService_ViewWithDeleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) Pull(ctx context.Context, in *PullVarRequest, opts ...grpc.CallOption) (*PullVarResponse, error) {
	out := new(PullVarResponse)
	err := c.cc.Invoke(ctx, VarService_Pull_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *varServiceClient) Sync(ctx context.Context, in *pb.Var, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, VarService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VarServiceServer is the server API for VarService service.
// All implementations must embed UnimplementedVarServiceServer
// for forward compatibility
type VarServiceServer interface {
	Create(context.Context, *pb.Var) (*pb.Var, error)
	Update(context.Context, *pb.Var) (*pb.Var, error)
	View(context.Context, *pb.Id) (*pb.Var, error)
	ViewByName(context.Context, *pb.Name) (*pb.Var, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *ListVarRequest) (*ListVarResponse, error)
	GetValue(context.Context, *pb.Id) (*pb.VarValue, error)
	SetValue(context.Context, *pb.VarValue) (*pb.MyBool, error)
	SetValueUnchecked(context.Context, *pb.VarValue) (*pb.MyBool, error)
	GetValueByName(context.Context, *pb.Name) (*pb.VarNameValue, error)
	SetValueByName(context.Context, *pb.VarNameValue) (*pb.MyBool, error)
	SetValueByNameUnchecked(context.Context, *pb.VarNameValue) (*pb.MyBool, error)
	ViewWithDeleted(context.Context, *pb.Id) (*pb.Var, error)
	Pull(context.Context, *PullVarRequest) (*PullVarResponse, error)
	Sync(context.Context, *pb.Var) (*pb.MyBool, error)
	mustEmbedUnimplementedVarServiceServer()
}

// UnimplementedVarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVarServiceServer struct {
}

func (UnimplementedVarServiceServer) Create(context.Context, *pb.Var) (*pb.Var, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVarServiceServer) Update(context.Context, *pb.Var) (*pb.Var, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVarServiceServer) View(context.Context, *pb.Id) (*pb.Var, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedVarServiceServer) ViewByName(context.Context, *pb.Name) (*pb.Var, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewByName not implemented")
}
func (UnimplementedVarServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVarServiceServer) List(context.Context, *ListVarRequest) (*ListVarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedVarServiceServer) GetValue(context.Context, *pb.Id) (*pb.VarValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedVarServiceServer) SetValue(context.Context, *pb.VarValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedVarServiceServer) SetValueUnchecked(context.Context, *pb.VarValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueUnchecked not implemented")
}
func (UnimplementedVarServiceServer) GetValueByName(context.Context, *pb.Name) (*pb.VarNameValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueByName not implemented")
}
func (UnimplementedVarServiceServer) SetValueByName(context.Context, *pb.VarNameValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueByName not implemented")
}
func (UnimplementedVarServiceServer) SetValueByNameUnchecked(context.Context, *pb.VarNameValue) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValueByNameUnchecked not implemented")
}
func (UnimplementedVarServiceServer) ViewWithDeleted(context.Context, *pb.Id) (*pb.Var, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedVarServiceServer) Pull(context.Context, *PullVarRequest) (*PullVarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedVarServiceServer) Sync(context.Context, *pb.Var) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedVarServiceServer) mustEmbedUnimplementedVarServiceServer() {}

// UnsafeVarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VarServiceServer will
// result in compilation errors.
type UnsafeVarServiceServer interface {
	mustEmbedUnimplementedVarServiceServer()
}

func RegisterVarServiceServer(s grpc.ServiceRegistrar, srv VarServiceServer) {
	s.RegisterService(&VarService_ServiceDesc, srv)
}

func _VarService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Var)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).Create(ctx, req.(*pb.Var))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Var)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).Update(ctx, req.(*pb.Var))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_ViewByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).ViewByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_ViewByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).ViewByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).List(ctx, req.(*ListVarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).GetValue(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.VarValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_SetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).SetValue(ctx, req.(*pb.VarValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_SetValueUnchecked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.VarValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).SetValueUnchecked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_SetValueUnchecked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).SetValueUnchecked(ctx, req.(*pb.VarValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_GetValueByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).GetValueByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_GetValueByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).GetValueByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_SetValueByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.VarNameValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).SetValueByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_SetValueByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).SetValueByName(ctx, req.(*pb.VarNameValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_SetValueByNameUnchecked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.VarNameValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).SetValueByNameUnchecked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_SetValueByNameUnchecked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).SetValueByNameUnchecked(ctx, req.(*pb.VarNameValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).ViewWithDeleted(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullVarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).Pull(ctx, req.(*PullVarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VarService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Var)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VarServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VarService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VarServiceServer).Sync(ctx, req.(*pb.Var))
	}
	return interceptor(ctx, in, info, handler)
}

// VarService_ServiceDesc is the grpc.ServiceDesc for VarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodes.VarService",
	HandlerType: (*VarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VarService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VarService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _VarService_View_Handler,
		},
		{
			MethodName: "ViewByName",
			Handler:    _VarService_ViewByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VarService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VarService_List_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _VarService_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _VarService_SetValue_Handler,
		},
		{
			MethodName: "SetValueUnchecked",
			Handler:    _VarService_SetValueUnchecked_Handler,
		},
		{
			MethodName: "GetValueByName",
			Handler:    _VarService_GetValueByName_Handler,
		},
		{
			MethodName: "SetValueByName",
			Handler:    _VarService_SetValueByName_Handler,
		},
		{
			MethodName: "SetValueByNameUnchecked",
			Handler:    _VarService_SetValueByNameUnchecked_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _VarService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _VarService_Pull_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _VarService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodes/var_service.proto",
}
