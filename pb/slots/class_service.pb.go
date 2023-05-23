// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: slots/class_service.proto

package slots

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

// class
type ListClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	Tags string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListClassRequest) Reset() {
	*x = ListClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassRequest) ProtoMessage() {}

func (x *ListClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassRequest.ProtoReflect.Descriptor instead.
func (*ListClassRequest) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListClassRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListClassRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ListClassRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Class []*pb.Class `protobuf:"bytes,3,rep,name=class,proto3" json:"class,omitempty"`
}

func (x *ListClassResponse) Reset() {
	*x = ListClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassResponse) ProtoMessage() {}

func (x *ListClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassResponse.ProtoReflect.Descriptor instead.
func (*ListClassResponse) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListClassResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListClassResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListClassResponse) GetClass() []*pb.Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type LinkClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LinkClassRequest) Reset() {
	*x = LinkClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkClassRequest) ProtoMessage() {}

func (x *LinkClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkClassRequest.ProtoReflect.Descriptor instead.
func (*LinkClassRequest) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{2}
}

func (x *LinkClassRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LinkClassRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PullClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PullClassRequest) Reset() {
	*x = PullClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullClassRequest) ProtoMessage() {}

func (x *PullClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullClassRequest.ProtoReflect.Descriptor instead.
func (*PullClassRequest) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{3}
}

func (x *PullClassRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullClassRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullClassRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PullClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64       `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Class []*pb.Class `protobuf:"bytes,3,rep,name=class,proto3" json:"class,omitempty"`
}

func (x *PullClassResponse) Reset() {
	*x = PullClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullClassResponse) ProtoMessage() {}

func (x *PullClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullClassResponse.ProtoReflect.Descriptor instead.
func (*PullClassResponse) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{4}
}

func (x *PullClassResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullClassResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullClassResponse) GetClass() []*pb.Class {
	if x != nil {
		return x.Class
	}
	return nil
}

// attr
type ListAttrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	ClassId string `protobuf:"bytes,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Tags    string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Type    string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListAttrRequest) Reset() {
	*x = ListAttrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttrRequest) ProtoMessage() {}

func (x *ListAttrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttrRequest.ProtoReflect.Descriptor instead.
func (*ListAttrRequest) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListAttrRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListAttrRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *ListAttrRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ListAttrRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Attr  []*pb.Attr `protobuf:"bytes,3,rep,name=attr,proto3" json:"attr,omitempty"`
}

func (x *ListAttrResponse) Reset() {
	*x = ListAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttrResponse) ProtoMessage() {}

func (x *ListAttrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttrResponse.ProtoReflect.Descriptor instead.
func (*ListAttrResponse) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListAttrResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListAttrResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListAttrResponse) GetAttr() []*pb.Attr {
	if x != nil {
		return x.Attr
	}
	return nil
}

type PullAttrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	ClassId string `protobuf:"bytes,4,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Type    string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PullAttrRequest) Reset() {
	*x = PullAttrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullAttrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullAttrRequest) ProtoMessage() {}

func (x *PullAttrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullAttrRequest.ProtoReflect.Descriptor instead.
func (*PullAttrRequest) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{7}
}

func (x *PullAttrRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullAttrRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullAttrRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *PullAttrRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PullAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64      `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Attr  []*pb.Attr `protobuf:"bytes,3,rep,name=attr,proto3" json:"attr,omitempty"`
}

func (x *PullAttrResponse) Reset() {
	*x = PullAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_class_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullAttrResponse) ProtoMessage() {}

func (x *PullAttrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_class_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullAttrResponse.ProtoReflect.Descriptor instead.
func (*PullAttrResponse) Descriptor() ([]byte, []int) {
	return file_slots_class_service_proto_rawDescGZIP(), []int{8}
}

func (x *PullAttrResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullAttrResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullAttrResponse) GetAttr() []*pb.Attr {
	if x != nil {
		return x.Attr
	}
	return nil
}

var File_slots_class_service_proto protoreflect.FileDescriptor

var file_slots_class_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6c, 0x6f,
	0x74, 0x73, 0x1a, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52,
	0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x22, 0x72, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x22, 0x6c,
	0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x10,
	0x50, 0x75, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x32, 0xf7, 0x02, 0x0a, 0x0c, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x00, 0x12, 0x20, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x00, 0x12,
	0x1b, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0a,
	0x56, 0x69, 0x65, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22,
	0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x6c, 0x6f, 0x74,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26,
	0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x17,
	0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x32, 0x8a, 0x05, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x22, 0x00,
	0x12, 0x22, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x11, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x6e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x1e, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x6c, 0x6f, 0x74, 0x73, 0x3b, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_slots_class_service_proto_rawDescOnce sync.Once
	file_slots_class_service_proto_rawDescData = file_slots_class_service_proto_rawDesc
)

func file_slots_class_service_proto_rawDescGZIP() []byte {
	file_slots_class_service_proto_rawDescOnce.Do(func() {
		file_slots_class_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_slots_class_service_proto_rawDescData)
	})
	return file_slots_class_service_proto_rawDescData
}

var file_slots_class_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_slots_class_service_proto_goTypes = []interface{}{
	(*ListClassRequest)(nil),  // 0: slots.ListClassRequest
	(*ListClassResponse)(nil), // 1: slots.ListClassResponse
	(*LinkClassRequest)(nil),  // 2: slots.LinkClassRequest
	(*PullClassRequest)(nil),  // 3: slots.PullClassRequest
	(*PullClassResponse)(nil), // 4: slots.PullClassResponse
	(*ListAttrRequest)(nil),   // 5: slots.ListAttrRequest
	(*ListAttrResponse)(nil),  // 6: slots.ListAttrResponse
	(*PullAttrRequest)(nil),   // 7: slots.PullAttrRequest
	(*PullAttrResponse)(nil),  // 8: slots.PullAttrResponse
	(*pb.Page)(nil),           // 9: pb.Page
	(*pb.Class)(nil),          // 10: pb.Class
	(*pb.Attr)(nil),           // 11: pb.Attr
	(*pb.Id)(nil),             // 12: pb.Id
	(*pb.Name)(nil),           // 13: pb.Name
	(*pb.AttrValue)(nil),      // 14: pb.AttrValue
	(*pb.AttrNameValue)(nil),  // 15: pb.AttrNameValue
	(*pb.MyBool)(nil),         // 16: pb.MyBool
}
var file_slots_class_service_proto_depIdxs = []int32{
	9,  // 0: slots.ListClassRequest.page:type_name -> pb.Page
	9,  // 1: slots.ListClassResponse.page:type_name -> pb.Page
	10, // 2: slots.ListClassResponse.class:type_name -> pb.Class
	10, // 3: slots.PullClassResponse.class:type_name -> pb.Class
	9,  // 4: slots.ListAttrRequest.page:type_name -> pb.Page
	9,  // 5: slots.ListAttrResponse.page:type_name -> pb.Page
	11, // 6: slots.ListAttrResponse.attr:type_name -> pb.Attr
	11, // 7: slots.PullAttrResponse.attr:type_name -> pb.Attr
	10, // 8: slots.ClassService.Create:input_type -> pb.Class
	10, // 9: slots.ClassService.Update:input_type -> pb.Class
	12, // 10: slots.ClassService.View:input_type -> pb.Id
	13, // 11: slots.ClassService.ViewByName:input_type -> pb.Name
	12, // 12: slots.ClassService.Delete:input_type -> pb.Id
	0,  // 13: slots.ClassService.List:input_type -> slots.ListClassRequest
	12, // 14: slots.ClassService.ViewWithDeleted:input_type -> pb.Id
	3,  // 15: slots.ClassService.Pull:input_type -> slots.PullClassRequest
	10, // 16: slots.ClassService.Sync:input_type -> pb.Class
	11, // 17: slots.AttrService.Create:input_type -> pb.Attr
	11, // 18: slots.AttrService.Update:input_type -> pb.Attr
	12, // 19: slots.AttrService.View:input_type -> pb.Id
	13, // 20: slots.AttrService.ViewByName:input_type -> pb.Name
	12, // 21: slots.AttrService.Delete:input_type -> pb.Id
	5,  // 22: slots.AttrService.List:input_type -> slots.ListAttrRequest
	12, // 23: slots.AttrService.GetValue:input_type -> pb.Id
	14, // 24: slots.AttrService.SetValue:input_type -> pb.AttrValue
	14, // 25: slots.AttrService.SetValueUnchecked:input_type -> pb.AttrValue
	13, // 26: slots.AttrService.GetValueByName:input_type -> pb.Name
	15, // 27: slots.AttrService.SetValueByName:input_type -> pb.AttrNameValue
	15, // 28: slots.AttrService.SetValueByNameUnchecked:input_type -> pb.AttrNameValue
	12, // 29: slots.AttrService.ViewWithDeleted:input_type -> pb.Id
	7,  // 30: slots.AttrService.Pull:input_type -> slots.PullAttrRequest
	11, // 31: slots.AttrService.Sync:input_type -> pb.Attr
	10, // 32: slots.ClassService.Create:output_type -> pb.Class
	10, // 33: slots.ClassService.Update:output_type -> pb.Class
	10, // 34: slots.ClassService.View:output_type -> pb.Class
	10, // 35: slots.ClassService.ViewByName:output_type -> pb.Class
	16, // 36: slots.ClassService.Delete:output_type -> pb.MyBool
	1,  // 37: slots.ClassService.List:output_type -> slots.ListClassResponse
	10, // 38: slots.ClassService.ViewWithDeleted:output_type -> pb.Class
	4,  // 39: slots.ClassService.Pull:output_type -> slots.PullClassResponse
	16, // 40: slots.ClassService.Sync:output_type -> pb.MyBool
	11, // 41: slots.AttrService.Create:output_type -> pb.Attr
	11, // 42: slots.AttrService.Update:output_type -> pb.Attr
	11, // 43: slots.AttrService.View:output_type -> pb.Attr
	11, // 44: slots.AttrService.ViewByName:output_type -> pb.Attr
	16, // 45: slots.AttrService.Delete:output_type -> pb.MyBool
	6,  // 46: slots.AttrService.List:output_type -> slots.ListAttrResponse
	14, // 47: slots.AttrService.GetValue:output_type -> pb.AttrValue
	16, // 48: slots.AttrService.SetValue:output_type -> pb.MyBool
	16, // 49: slots.AttrService.SetValueUnchecked:output_type -> pb.MyBool
	15, // 50: slots.AttrService.GetValueByName:output_type -> pb.AttrNameValue
	16, // 51: slots.AttrService.SetValueByName:output_type -> pb.MyBool
	16, // 52: slots.AttrService.SetValueByNameUnchecked:output_type -> pb.MyBool
	11, // 53: slots.AttrService.ViewWithDeleted:output_type -> pb.Attr
	8,  // 54: slots.AttrService.Pull:output_type -> slots.PullAttrResponse
	16, // 55: slots.AttrService.Sync:output_type -> pb.MyBool
	32, // [32:56] is the sub-list for method output_type
	8,  // [8:32] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_slots_class_service_proto_init() }
func file_slots_class_service_proto_init() {
	if File_slots_class_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slots_class_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClassRequest); i {
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
		file_slots_class_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClassResponse); i {
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
		file_slots_class_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkClassRequest); i {
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
		file_slots_class_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullClassRequest); i {
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
		file_slots_class_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullClassResponse); i {
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
		file_slots_class_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttrRequest); i {
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
		file_slots_class_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttrResponse); i {
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
		file_slots_class_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullAttrRequest); i {
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
		file_slots_class_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullAttrResponse); i {
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
			RawDescriptor: file_slots_class_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_slots_class_service_proto_goTypes,
		DependencyIndexes: file_slots_class_service_proto_depIdxs,
		MessageInfos:      file_slots_class_service_proto_msgTypes,
	}.Build()
	File_slots_class_service_proto = out.File
	file_slots_class_service_proto_rawDesc = nil
	file_slots_class_service_proto_goTypes = nil
	file_slots_class_service_proto_depIdxs = nil
}
