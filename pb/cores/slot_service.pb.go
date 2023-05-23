// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: cores/slot_service.proto

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

type ListSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	DeviceId string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Tags     string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListSlotRequest) Reset() {
	*x = ListSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSlotRequest) ProtoMessage() {}

func (x *ListSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSlotRequest.ProtoReflect.Descriptor instead.
func (*ListSlotRequest) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSlotRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListSlotRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ListSlotRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ListSlotRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Slot  []*pb.Slot `protobuf:"bytes,3,rep,name=slot,proto3" json:"slot,omitempty"`
}

func (x *ListSlotResponse) Reset() {
	*x = ListSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSlotResponse) ProtoMessage() {}

func (x *ListSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSlotResponse.ProtoReflect.Descriptor instead.
func (*ListSlotResponse) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSlotResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListSlotResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListSlotResponse) GetSlot() []*pb.Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

type ViewSlotByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ViewSlotByNameRequest) Reset() {
	*x = ViewSlotByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewSlotByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewSlotByNameRequest) ProtoMessage() {}

func (x *ViewSlotByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewSlotByNameRequest.ProtoReflect.Descriptor instead.
func (*ViewSlotByNameRequest) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{2}
}

func (x *ViewSlotByNameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ViewSlotByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LinkSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LinkSlotRequest) Reset() {
	*x = LinkSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkSlotRequest) ProtoMessage() {}

func (x *LinkSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkSlotRequest.ProtoReflect.Descriptor instead.
func (*LinkSlotRequest) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{3}
}

func (x *LinkSlotRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LinkSlotRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CloneSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *CloneSlotRequest) Reset() {
	*x = CloneSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneSlotRequest) ProtoMessage() {}

func (x *CloneSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneSlotRequest.ProtoReflect.Descriptor instead.
func (*CloneSlotRequest) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{4}
}

func (x *CloneSlotRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloneSlotRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type PullSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After    int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit    uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PullSlotRequest) Reset() {
	*x = PullSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullSlotRequest) ProtoMessage() {}

func (x *PullSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullSlotRequest.ProtoReflect.Descriptor instead.
func (*PullSlotRequest) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{5}
}

func (x *PullSlotRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullSlotRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullSlotRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *PullSlotRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PullSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64      `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Slot  []*pb.Slot `protobuf:"bytes,3,rep,name=slot,proto3" json:"slot,omitempty"`
}

func (x *PullSlotResponse) Reset() {
	*x = PullSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_slot_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullSlotResponse) ProtoMessage() {}

func (x *PullSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_slot_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullSlotResponse.ProtoReflect.Descriptor instead.
func (*PullSlotResponse) Descriptor() ([]byte, []int) {
	return file_cores_slot_service_proto_rawDescGZIP(), []int{6}
}

func (x *PullSlotResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullSlotResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullSlotResponse) GetSlot() []*pb.Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

var File_cores_slot_service_proto protoreflect.FileDescriptor

var file_cores_slot_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x1a, 0x12, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x56, 0x69, 0x65, 0x77,
	0x53, 0x6c, 0x6f, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a,
	0x10, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x6e,
	0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c,
	0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c,
	0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x32, 0x84, 0x04, 0x0a,
	0x0b, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x53, 0x6c, 0x6f, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00,
	0x12, 0x26, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x75,
	0x6c, 0x6c, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x25, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cores_slot_service_proto_rawDescOnce sync.Once
	file_cores_slot_service_proto_rawDescData = file_cores_slot_service_proto_rawDesc
)

func file_cores_slot_service_proto_rawDescGZIP() []byte {
	file_cores_slot_service_proto_rawDescOnce.Do(func() {
		file_cores_slot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cores_slot_service_proto_rawDescData)
	})
	return file_cores_slot_service_proto_rawDescData
}

var file_cores_slot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cores_slot_service_proto_goTypes = []interface{}{
	(*ListSlotRequest)(nil),       // 0: cores.ListSlotRequest
	(*ListSlotResponse)(nil),      // 1: cores.ListSlotResponse
	(*ViewSlotByNameRequest)(nil), // 2: cores.ViewSlotByNameRequest
	(*LinkSlotRequest)(nil),       // 3: cores.LinkSlotRequest
	(*CloneSlotRequest)(nil),      // 4: cores.CloneSlotRequest
	(*PullSlotRequest)(nil),       // 5: cores.PullSlotRequest
	(*PullSlotResponse)(nil),      // 6: cores.PullSlotResponse
	(*pb.Page)(nil),               // 7: pb.Page
	(*pb.Slot)(nil),               // 8: pb.Slot
	(*pb.Id)(nil),                 // 9: pb.Id
	(*pb.Name)(nil),               // 10: pb.Name
	(*pb.MyBool)(nil),             // 11: pb.MyBool
}
var file_cores_slot_service_proto_depIdxs = []int32{
	7,  // 0: cores.ListSlotRequest.page:type_name -> pb.Page
	7,  // 1: cores.ListSlotResponse.page:type_name -> pb.Page
	8,  // 2: cores.ListSlotResponse.slot:type_name -> pb.Slot
	8,  // 3: cores.PullSlotResponse.slot:type_name -> pb.Slot
	8,  // 4: cores.SlotService.Create:input_type -> pb.Slot
	8,  // 5: cores.SlotService.Update:input_type -> pb.Slot
	9,  // 6: cores.SlotService.View:input_type -> pb.Id
	2,  // 7: cores.SlotService.ViewByName:input_type -> cores.ViewSlotByNameRequest
	10, // 8: cores.SlotService.ViewByNameFull:input_type -> pb.Name
	9,  // 9: cores.SlotService.Delete:input_type -> pb.Id
	0,  // 10: cores.SlotService.List:input_type -> cores.ListSlotRequest
	3,  // 11: cores.SlotService.Link:input_type -> cores.LinkSlotRequest
	4,  // 12: cores.SlotService.Clone:input_type -> cores.CloneSlotRequest
	9,  // 13: cores.SlotService.ViewWithDeleted:input_type -> pb.Id
	5,  // 14: cores.SlotService.Pull:input_type -> cores.PullSlotRequest
	8,  // 15: cores.SlotService.Sync:input_type -> pb.Slot
	8,  // 16: cores.SlotService.Create:output_type -> pb.Slot
	8,  // 17: cores.SlotService.Update:output_type -> pb.Slot
	8,  // 18: cores.SlotService.View:output_type -> pb.Slot
	8,  // 19: cores.SlotService.ViewByName:output_type -> pb.Slot
	8,  // 20: cores.SlotService.ViewByNameFull:output_type -> pb.Slot
	11, // 21: cores.SlotService.Delete:output_type -> pb.MyBool
	1,  // 22: cores.SlotService.List:output_type -> cores.ListSlotResponse
	11, // 23: cores.SlotService.Link:output_type -> pb.MyBool
	11, // 24: cores.SlotService.Clone:output_type -> pb.MyBool
	8,  // 25: cores.SlotService.ViewWithDeleted:output_type -> pb.Slot
	6,  // 26: cores.SlotService.Pull:output_type -> cores.PullSlotResponse
	11, // 27: cores.SlotService.Sync:output_type -> pb.MyBool
	16, // [16:28] is the sub-list for method output_type
	4,  // [4:16] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cores_slot_service_proto_init() }
func file_cores_slot_service_proto_init() {
	if File_cores_slot_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cores_slot_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSlotRequest); i {
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
		file_cores_slot_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSlotResponse); i {
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
		file_cores_slot_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewSlotByNameRequest); i {
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
		file_cores_slot_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkSlotRequest); i {
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
		file_cores_slot_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneSlotRequest); i {
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
		file_cores_slot_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullSlotRequest); i {
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
		file_cores_slot_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullSlotResponse); i {
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
			RawDescriptor: file_cores_slot_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cores_slot_service_proto_goTypes,
		DependencyIndexes: file_cores_slot_service_proto_depIdxs,
		MessageInfos:      file_cores_slot_service_proto_msgTypes,
	}.Build()
	File_cores_slot_service_proto = out.File
	file_cores_slot_service_proto_rawDesc = nil
	file_cores_slot_service_proto_goTypes = nil
	file_cores_slot_service_proto_depIdxs = nil
}
