// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: cores/proxy_service.proto

package cores

import (
	pb "github.com/snple/kokomi/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	DeviceId string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Tags     string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ProxyListRequest) Reset() {
	*x = ProxyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyListRequest) ProtoMessage() {}

func (x *ProxyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyListRequest.ProtoReflect.Descriptor instead.
func (*ProxyListRequest) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ProxyListRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ProxyListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ProxyListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ProxyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Proxy []*pb.Proxy `protobuf:"bytes,3,rep,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *ProxyListResponse) Reset() {
	*x = ProxyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyListResponse) ProtoMessage() {}

func (x *ProxyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyListResponse.ProtoReflect.Descriptor instead.
func (*ProxyListResponse) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ProxyListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ProxyListResponse) GetProxy() []*pb.Proxy {
	if x != nil {
		return x.Proxy
	}
	return nil
}

type ProxyNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProxyNameRequest) Reset() {
	*x = ProxyNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyNameRequest) ProtoMessage() {}

func (x *ProxyNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyNameRequest.ProtoReflect.Descriptor instead.
func (*ProxyNameRequest) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyNameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ProxyNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProxyLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProxyLinkRequest) Reset() {
	*x = ProxyLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyLinkRequest) ProtoMessage() {}

func (x *ProxyLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyLinkRequest.ProtoReflect.Descriptor instead.
func (*ProxyLinkRequest) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProxyLinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProxyLinkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ProxyCloneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *ProxyCloneRequest) Reset() {
	*x = ProxyCloneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyCloneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyCloneRequest) ProtoMessage() {}

func (x *ProxyCloneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyCloneRequest.ProtoReflect.Descriptor instead.
func (*ProxyCloneRequest) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{4}
}

func (x *ProxyCloneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProxyCloneRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ProxyPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After    int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit    uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ProxyPullRequest) Reset() {
	*x = ProxyPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPullRequest) ProtoMessage() {}

func (x *ProxyPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPullRequest.ProtoReflect.Descriptor instead.
func (*ProxyPullRequest) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{5}
}

func (x *ProxyPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ProxyPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ProxyPullRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ProxyPullRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ProxyPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64       `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Proxy []*pb.Proxy `protobuf:"bytes,3,rep,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *ProxyPullResponse) Reset() {
	*x = ProxyPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_proxy_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPullResponse) ProtoMessage() {}

func (x *ProxyPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_proxy_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPullResponse.ProtoReflect.Descriptor instead.
func (*ProxyPullResponse) Descriptor() ([]byte, []int) {
	return file_cores_proxy_service_proto_rawDescGZIP(), []int{6}
}

func (x *ProxyPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ProxyPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ProxyPullResponse) GetProxy() []*pb.Proxy {
	if x != nil {
		return x.Proxy
	}
	return nil
}

var File_cores_proxy_service_proto protoreflect.FileDescriptor

var file_cores_proxy_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75,
	0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22,
	0x43, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x40, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x6f, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x32, 0x83, 0x04, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x75, 0x6c,
	0x6c, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6c, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42,
	0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64,
	0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x79,
	0x6e, 0x63, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f,
	0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x3b,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cores_proxy_service_proto_rawDescOnce sync.Once
	file_cores_proxy_service_proto_rawDescData = file_cores_proxy_service_proto_rawDesc
)

func file_cores_proxy_service_proto_rawDescGZIP() []byte {
	file_cores_proxy_service_proto_rawDescOnce.Do(func() {
		file_cores_proxy_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cores_proxy_service_proto_rawDescData)
	})
	return file_cores_proxy_service_proto_rawDescData
}

var file_cores_proxy_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cores_proxy_service_proto_goTypes = []interface{}{
	(*ProxyListRequest)(nil),  // 0: cores.ProxyListRequest
	(*ProxyListResponse)(nil), // 1: cores.ProxyListResponse
	(*ProxyNameRequest)(nil),  // 2: cores.ProxyNameRequest
	(*ProxyLinkRequest)(nil),  // 3: cores.ProxyLinkRequest
	(*ProxyCloneRequest)(nil), // 4: cores.ProxyCloneRequest
	(*ProxyPullRequest)(nil),  // 5: cores.ProxyPullRequest
	(*ProxyPullResponse)(nil), // 6: cores.ProxyPullResponse
	(*pb.Page)(nil),           // 7: pb.Page
	(*pb.Proxy)(nil),          // 8: pb.Proxy
	(*pb.Id)(nil),             // 9: pb.Id
	(*pb.Name)(nil),           // 10: pb.Name
	(*pb.MyBool)(nil),         // 11: pb.MyBool
}
var file_cores_proxy_service_proto_depIdxs = []int32{
	7,  // 0: cores.ProxyListRequest.page:type_name -> pb.Page
	7,  // 1: cores.ProxyListResponse.page:type_name -> pb.Page
	8,  // 2: cores.ProxyListResponse.proxy:type_name -> pb.Proxy
	8,  // 3: cores.ProxyPullResponse.proxy:type_name -> pb.Proxy
	8,  // 4: cores.ProxyService.Create:input_type -> pb.Proxy
	8,  // 5: cores.ProxyService.Update:input_type -> pb.Proxy
	9,  // 6: cores.ProxyService.View:input_type -> pb.Id
	2,  // 7: cores.ProxyService.Name:input_type -> cores.ProxyNameRequest
	10, // 8: cores.ProxyService.NameFull:input_type -> pb.Name
	9,  // 9: cores.ProxyService.Delete:input_type -> pb.Id
	0,  // 10: cores.ProxyService.List:input_type -> cores.ProxyListRequest
	3,  // 11: cores.ProxyService.Link:input_type -> cores.ProxyLinkRequest
	4,  // 12: cores.ProxyService.Clone:input_type -> cores.ProxyCloneRequest
	9,  // 13: cores.ProxyService.ViewWithDeleted:input_type -> pb.Id
	5,  // 14: cores.ProxyService.Pull:input_type -> cores.ProxyPullRequest
	8,  // 15: cores.ProxyService.Sync:input_type -> pb.Proxy
	8,  // 16: cores.ProxyService.Create:output_type -> pb.Proxy
	8,  // 17: cores.ProxyService.Update:output_type -> pb.Proxy
	8,  // 18: cores.ProxyService.View:output_type -> pb.Proxy
	8,  // 19: cores.ProxyService.Name:output_type -> pb.Proxy
	8,  // 20: cores.ProxyService.NameFull:output_type -> pb.Proxy
	11, // 21: cores.ProxyService.Delete:output_type -> pb.MyBool
	1,  // 22: cores.ProxyService.List:output_type -> cores.ProxyListResponse
	11, // 23: cores.ProxyService.Link:output_type -> pb.MyBool
	11, // 24: cores.ProxyService.Clone:output_type -> pb.MyBool
	8,  // 25: cores.ProxyService.ViewWithDeleted:output_type -> pb.Proxy
	6,  // 26: cores.ProxyService.Pull:output_type -> cores.ProxyPullResponse
	11, // 27: cores.ProxyService.Sync:output_type -> pb.MyBool
	16, // [16:28] is the sub-list for method output_type
	4,  // [4:16] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cores_proxy_service_proto_init() }
func file_cores_proxy_service_proto_init() {
	if File_cores_proxy_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cores_proxy_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyNameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyLinkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyCloneRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPullRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cores_proxy_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPullResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cores_proxy_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cores_proxy_service_proto_goTypes,
		DependencyIndexes: file_cores_proxy_service_proto_depIdxs,
		MessageInfos:      file_cores_proxy_service_proto_msgTypes,
	}.Build()
	File_cores_proxy_service_proto = out.File
	file_cores_proxy_service_proto_rawDesc = nil
	file_cores_proxy_service_proto_goTypes = nil
	file_cores_proxy_service_proto_depIdxs = nil
}
