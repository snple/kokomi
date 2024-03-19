// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: nodes/option_service.proto

package nodes

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

type OptionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	Tags string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *OptionListRequest) Reset() {
	*x = OptionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_option_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionListRequest) ProtoMessage() {}

func (x *OptionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_option_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionListRequest.ProtoReflect.Descriptor instead.
func (*OptionListRequest) Descriptor() ([]byte, []int) {
	return file_nodes_option_service_proto_rawDescGZIP(), []int{0}
}

func (x *OptionListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *OptionListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *OptionListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type OptionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *pb.Page     `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count  uint32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Option []*pb.Option `protobuf:"bytes,3,rep,name=option,proto3" json:"option,omitempty"`
}

func (x *OptionListResponse) Reset() {
	*x = OptionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_option_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionListResponse) ProtoMessage() {}

func (x *OptionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_option_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionListResponse.ProtoReflect.Descriptor instead.
func (*OptionListResponse) Descriptor() ([]byte, []int) {
	return file_nodes_option_service_proto_rawDescGZIP(), []int{1}
}

func (x *OptionListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *OptionListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OptionListResponse) GetOption() []*pb.Option {
	if x != nil {
		return x.Option
	}
	return nil
}

type OptionLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OptionLinkRequest) Reset() {
	*x = OptionLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_option_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionLinkRequest) ProtoMessage() {}

func (x *OptionLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_option_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionLinkRequest.ProtoReflect.Descriptor instead.
func (*OptionLinkRequest) Descriptor() ([]byte, []int) {
	return file_nodes_option_service_proto_rawDescGZIP(), []int{2}
}

func (x *OptionLinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OptionLinkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type OptionPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *OptionPullRequest) Reset() {
	*x = OptionPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_option_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionPullRequest) ProtoMessage() {}

func (x *OptionPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_option_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionPullRequest.ProtoReflect.Descriptor instead.
func (*OptionPullRequest) Descriptor() ([]byte, []int) {
	return file_nodes_option_service_proto_rawDescGZIP(), []int{3}
}

func (x *OptionPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *OptionPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *OptionPullRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type OptionPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After  int64        `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit  uint32       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Option []*pb.Option `protobuf:"bytes,3,rep,name=option,proto3" json:"option,omitempty"`
}

func (x *OptionPullResponse) Reset() {
	*x = OptionPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_option_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionPullResponse) ProtoMessage() {}

func (x *OptionPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_option_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionPullResponse.ProtoReflect.Descriptor instead.
func (*OptionPullResponse) Descriptor() ([]byte, []int) {
	return file_nodes_option_service_proto_rawDescGZIP(), []int{4}
}

func (x *OptionPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *OptionPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *OptionPullResponse) GetOption() []*pb.Option {
	if x != nil {
		return x.Option
	}
	return nil
}

var File_nodes_option_service_proto protoreflect.FileDescriptor

var file_nodes_option_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x1a, 0x14, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x59, 0x0a, 0x11, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6c, 0x0a, 0x12, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x11, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x64, 0x0a, 0x12, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xfe, 0x02, 0x0a, 0x0d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x20, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodes_option_service_proto_rawDescOnce sync.Once
	file_nodes_option_service_proto_rawDescData = file_nodes_option_service_proto_rawDesc
)

func file_nodes_option_service_proto_rawDescGZIP() []byte {
	file_nodes_option_service_proto_rawDescOnce.Do(func() {
		file_nodes_option_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodes_option_service_proto_rawDescData)
	})
	return file_nodes_option_service_proto_rawDescData
}

var file_nodes_option_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nodes_option_service_proto_goTypes = []interface{}{
	(*OptionListRequest)(nil),  // 0: nodes.OptionListRequest
	(*OptionListResponse)(nil), // 1: nodes.OptionListResponse
	(*OptionLinkRequest)(nil),  // 2: nodes.OptionLinkRequest
	(*OptionPullRequest)(nil),  // 3: nodes.OptionPullRequest
	(*OptionPullResponse)(nil), // 4: nodes.OptionPullResponse
	(*pb.Page)(nil),            // 5: pb.Page
	(*pb.Option)(nil),          // 6: pb.Option
	(*pb.Id)(nil),              // 7: pb.Id
	(*pb.Name)(nil),            // 8: pb.Name
	(*pb.MyBool)(nil),          // 9: pb.MyBool
}
var file_nodes_option_service_proto_depIdxs = []int32{
	5,  // 0: nodes.OptionListRequest.page:type_name -> pb.Page
	5,  // 1: nodes.OptionListResponse.page:type_name -> pb.Page
	6,  // 2: nodes.OptionListResponse.option:type_name -> pb.Option
	6,  // 3: nodes.OptionPullResponse.option:type_name -> pb.Option
	6,  // 4: nodes.OptionService.Create:input_type -> pb.Option
	6,  // 5: nodes.OptionService.Update:input_type -> pb.Option
	7,  // 6: nodes.OptionService.View:input_type -> pb.Id
	8,  // 7: nodes.OptionService.Name:input_type -> pb.Name
	7,  // 8: nodes.OptionService.Delete:input_type -> pb.Id
	0,  // 9: nodes.OptionService.List:input_type -> nodes.OptionListRequest
	7,  // 10: nodes.OptionService.ViewWithDeleted:input_type -> pb.Id
	3,  // 11: nodes.OptionService.Pull:input_type -> nodes.OptionPullRequest
	6,  // 12: nodes.OptionService.Sync:input_type -> pb.Option
	6,  // 13: nodes.OptionService.Create:output_type -> pb.Option
	6,  // 14: nodes.OptionService.Update:output_type -> pb.Option
	6,  // 15: nodes.OptionService.View:output_type -> pb.Option
	6,  // 16: nodes.OptionService.Name:output_type -> pb.Option
	9,  // 17: nodes.OptionService.Delete:output_type -> pb.MyBool
	1,  // 18: nodes.OptionService.List:output_type -> nodes.OptionListResponse
	6,  // 19: nodes.OptionService.ViewWithDeleted:output_type -> pb.Option
	4,  // 20: nodes.OptionService.Pull:output_type -> nodes.OptionPullResponse
	9,  // 21: nodes.OptionService.Sync:output_type -> pb.MyBool
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_nodes_option_service_proto_init() }
func file_nodes_option_service_proto_init() {
	if File_nodes_option_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodes_option_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionListRequest); i {
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
		file_nodes_option_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionListResponse); i {
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
		file_nodes_option_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionLinkRequest); i {
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
		file_nodes_option_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionPullRequest); i {
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
		file_nodes_option_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionPullResponse); i {
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
			RawDescriptor: file_nodes_option_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodes_option_service_proto_goTypes,
		DependencyIndexes: file_nodes_option_service_proto_depIdxs,
		MessageInfos:      file_nodes_option_service_proto_msgTypes,
	}.Build()
	File_nodes_option_service_proto = out.File
	file_nodes_option_service_proto_rawDesc = nil
	file_nodes_option_service_proto_goTypes = nil
	file_nodes_option_service_proto_depIdxs = nil
}
