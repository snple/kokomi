// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: edges/sync_service.proto

package edges

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

type SyncUpdated struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// string id = 1;
	Updated       int64 `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SyncUpdated) Reset() {
	*x = SyncUpdated{}
	mi := &file_edges_sync_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SyncUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUpdated) ProtoMessage() {}

func (x *SyncUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_edges_sync_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUpdated.ProtoReflect.Descriptor instead.
func (*SyncUpdated) Descriptor() ([]byte, []int) {
	return file_edges_sync_service_proto_rawDescGZIP(), []int{0}
}

func (x *SyncUpdated) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_edges_sync_service_proto protoreflect.FileDescriptor

var file_edges_sync_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x32, 0xde, 0x0b, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x11, 0x57, 0x61, 0x69, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x32, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x46,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x46, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42,
	0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x13, 0x57, 0x61, 0x69, 0x74, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x73, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edges_sync_service_proto_rawDescOnce sync.Once
	file_edges_sync_service_proto_rawDescData = file_edges_sync_service_proto_rawDesc
)

func file_edges_sync_service_proto_rawDescGZIP() []byte {
	file_edges_sync_service_proto_rawDescOnce.Do(func() {
		file_edges_sync_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_edges_sync_service_proto_rawDescData)
	})
	return file_edges_sync_service_proto_rawDescData
}

var file_edges_sync_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_edges_sync_service_proto_goTypes = []any{
	(*SyncUpdated)(nil), // 0: edges.SyncUpdated
	(*pb.MyEmpty)(nil),  // 1: pb.MyEmpty
	(*pb.MyBool)(nil),   // 2: pb.MyBool
}
var file_edges_sync_service_proto_depIdxs = []int32{
	0,  // 0: edges.SyncService.SetDeviceUpdated:input_type -> edges.SyncUpdated
	1,  // 1: edges.SyncService.GetDeviceUpdated:input_type -> pb.MyEmpty
	1,  // 2: edges.SyncService.WaitDeviceUpdated:input_type -> pb.MyEmpty
	0,  // 3: edges.SyncService.SetSlotUpdated:input_type -> edges.SyncUpdated
	1,  // 4: edges.SyncService.GetSlotUpdated:input_type -> pb.MyEmpty
	0,  // 5: edges.SyncService.SetOptionUpdated:input_type -> edges.SyncUpdated
	1,  // 6: edges.SyncService.GetOptionUpdated:input_type -> pb.MyEmpty
	0,  // 7: edges.SyncService.SetPortUpdated:input_type -> edges.SyncUpdated
	1,  // 8: edges.SyncService.GetPortUpdated:input_type -> pb.MyEmpty
	0,  // 9: edges.SyncService.SetProxyUpdated:input_type -> edges.SyncUpdated
	1,  // 10: edges.SyncService.GetProxyUpdated:input_type -> pb.MyEmpty
	0,  // 11: edges.SyncService.SetSourceUpdated:input_type -> edges.SyncUpdated
	1,  // 12: edges.SyncService.GetSourceUpdated:input_type -> pb.MyEmpty
	0,  // 13: edges.SyncService.SetTagUpdated:input_type -> edges.SyncUpdated
	1,  // 14: edges.SyncService.GetTagUpdated:input_type -> pb.MyEmpty
	0,  // 15: edges.SyncService.SetConstUpdated:input_type -> edges.SyncUpdated
	1,  // 16: edges.SyncService.GetConstUpdated:input_type -> pb.MyEmpty
	0,  // 17: edges.SyncService.SetClassUpdated:input_type -> edges.SyncUpdated
	1,  // 18: edges.SyncService.GetClassUpdated:input_type -> pb.MyEmpty
	0,  // 19: edges.SyncService.SetAttrUpdated:input_type -> edges.SyncUpdated
	1,  // 20: edges.SyncService.GetAttrUpdated:input_type -> pb.MyEmpty
	0,  // 21: edges.SyncService.SetLogicUpdated:input_type -> edges.SyncUpdated
	1,  // 22: edges.SyncService.GetLogicUpdated:input_type -> pb.MyEmpty
	0,  // 23: edges.SyncService.SetFnUpdated:input_type -> edges.SyncUpdated
	1,  // 24: edges.SyncService.GetFnUpdated:input_type -> pb.MyEmpty
	0,  // 25: edges.SyncService.SetTagValueUpdated:input_type -> edges.SyncUpdated
	1,  // 26: edges.SyncService.GetTagValueUpdated:input_type -> pb.MyEmpty
	1,  // 27: edges.SyncService.WaitTagValueUpdated:input_type -> pb.MyEmpty
	2,  // 28: edges.SyncService.SetDeviceUpdated:output_type -> pb.MyBool
	0,  // 29: edges.SyncService.GetDeviceUpdated:output_type -> edges.SyncUpdated
	2,  // 30: edges.SyncService.WaitDeviceUpdated:output_type -> pb.MyBool
	2,  // 31: edges.SyncService.SetSlotUpdated:output_type -> pb.MyBool
	0,  // 32: edges.SyncService.GetSlotUpdated:output_type -> edges.SyncUpdated
	2,  // 33: edges.SyncService.SetOptionUpdated:output_type -> pb.MyBool
	0,  // 34: edges.SyncService.GetOptionUpdated:output_type -> edges.SyncUpdated
	2,  // 35: edges.SyncService.SetPortUpdated:output_type -> pb.MyBool
	0,  // 36: edges.SyncService.GetPortUpdated:output_type -> edges.SyncUpdated
	2,  // 37: edges.SyncService.SetProxyUpdated:output_type -> pb.MyBool
	0,  // 38: edges.SyncService.GetProxyUpdated:output_type -> edges.SyncUpdated
	2,  // 39: edges.SyncService.SetSourceUpdated:output_type -> pb.MyBool
	0,  // 40: edges.SyncService.GetSourceUpdated:output_type -> edges.SyncUpdated
	2,  // 41: edges.SyncService.SetTagUpdated:output_type -> pb.MyBool
	0,  // 42: edges.SyncService.GetTagUpdated:output_type -> edges.SyncUpdated
	2,  // 43: edges.SyncService.SetConstUpdated:output_type -> pb.MyBool
	0,  // 44: edges.SyncService.GetConstUpdated:output_type -> edges.SyncUpdated
	2,  // 45: edges.SyncService.SetClassUpdated:output_type -> pb.MyBool
	0,  // 46: edges.SyncService.GetClassUpdated:output_type -> edges.SyncUpdated
	2,  // 47: edges.SyncService.SetAttrUpdated:output_type -> pb.MyBool
	0,  // 48: edges.SyncService.GetAttrUpdated:output_type -> edges.SyncUpdated
	2,  // 49: edges.SyncService.SetLogicUpdated:output_type -> pb.MyBool
	0,  // 50: edges.SyncService.GetLogicUpdated:output_type -> edges.SyncUpdated
	2,  // 51: edges.SyncService.SetFnUpdated:output_type -> pb.MyBool
	0,  // 52: edges.SyncService.GetFnUpdated:output_type -> edges.SyncUpdated
	2,  // 53: edges.SyncService.SetTagValueUpdated:output_type -> pb.MyBool
	0,  // 54: edges.SyncService.GetTagValueUpdated:output_type -> edges.SyncUpdated
	2,  // 55: edges.SyncService.WaitTagValueUpdated:output_type -> pb.MyBool
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_edges_sync_service_proto_init() }
func file_edges_sync_service_proto_init() {
	if File_edges_sync_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_edges_sync_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edges_sync_service_proto_goTypes,
		DependencyIndexes: file_edges_sync_service_proto_depIdxs,
		MessageInfos:      file_edges_sync_service_proto_msgTypes,
	}.Build()
	File_edges_sync_service_proto = out.File
	file_edges_sync_service_proto_rawDesc = nil
	file_edges_sync_service_proto_goTypes = nil
	file_edges_sync_service_proto_depIdxs = nil
}
