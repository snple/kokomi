// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: slots/slot_service.proto

package slots

import (
	pb "github.com/snple/beacon/pb"
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

type LoginSlotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret        string                 `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginSlotRequest) Reset() {
	*x = LoginSlotRequest{}
	mi := &file_slots_slot_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSlotRequest) ProtoMessage() {}

func (x *LoginSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_slot_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSlotRequest.ProtoReflect.Descriptor instead.
func (*LoginSlotRequest) Descriptor() ([]byte, []int) {
	return file_slots_slot_service_proto_rawDescGZIP(), []int{0}
}

func (x *LoginSlotRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginSlotRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type LoginSlotReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginSlotReply) Reset() {
	*x = LoginSlotReply{}
	mi := &file_slots_slot_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginSlotReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSlotReply) ProtoMessage() {}

func (x *LoginSlotReply) ProtoReflect() protoreflect.Message {
	mi := &file_slots_slot_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSlotReply.ProtoReflect.Descriptor instead.
func (*LoginSlotReply) Descriptor() ([]byte, []int) {
	return file_slots_slot_service_proto_rawDescGZIP(), []int{1}
}

func (x *LoginSlotReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginSlotReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SlotLinkRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// string id = 1;
	Status        int32 `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlotLinkRequest) Reset() {
	*x = SlotLinkRequest{}
	mi := &file_slots_slot_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlotLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotLinkRequest) ProtoMessage() {}

func (x *SlotLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_slot_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotLinkRequest.ProtoReflect.Descriptor instead.
func (*SlotLinkRequest) Descriptor() ([]byte, []int) {
	return file_slots_slot_service_proto_rawDescGZIP(), []int{2}
}

func (x *SlotLinkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_slots_slot_service_proto protoreflect.FileDescriptor

var file_slots_slot_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6c, 0x6f, 0x74,
	0x73, 0x1a, 0x12, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x10,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x29, 0x0a, 0x0f, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb7, 0x01, 0x0a, 0x0b,
	0x53, 0x6c, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x16, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42,
	0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x3b, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slots_slot_service_proto_rawDescOnce sync.Once
	file_slots_slot_service_proto_rawDescData = file_slots_slot_service_proto_rawDesc
)

func file_slots_slot_service_proto_rawDescGZIP() []byte {
	file_slots_slot_service_proto_rawDescOnce.Do(func() {
		file_slots_slot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_slots_slot_service_proto_rawDescData)
	})
	return file_slots_slot_service_proto_rawDescData
}

var file_slots_slot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_slots_slot_service_proto_goTypes = []any{
	(*LoginSlotRequest)(nil), // 0: slots.LoginSlotRequest
	(*LoginSlotReply)(nil),   // 1: slots.LoginSlotReply
	(*SlotLinkRequest)(nil),  // 2: slots.SlotLinkRequest
	(*pb.Slot)(nil),          // 3: pb.Slot
	(*pb.MyEmpty)(nil),       // 4: pb.MyEmpty
	(*pb.MyBool)(nil),        // 5: pb.MyBool
}
var file_slots_slot_service_proto_depIdxs = []int32{
	0, // 0: slots.SlotService.Login:input_type -> slots.LoginSlotRequest
	3, // 1: slots.SlotService.Update:input_type -> pb.Slot
	4, // 2: slots.SlotService.View:input_type -> pb.MyEmpty
	2, // 3: slots.SlotService.Link:input_type -> slots.SlotLinkRequest
	1, // 4: slots.SlotService.Login:output_type -> slots.LoginSlotReply
	3, // 5: slots.SlotService.Update:output_type -> pb.Slot
	3, // 6: slots.SlotService.View:output_type -> pb.Slot
	5, // 7: slots.SlotService.Link:output_type -> pb.MyBool
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_slots_slot_service_proto_init() }
func file_slots_slot_service_proto_init() {
	if File_slots_slot_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_slots_slot_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slots_slot_service_proto_goTypes,
		DependencyIndexes: file_slots_slot_service_proto_depIdxs,
		MessageInfos:      file_slots_slot_service_proto_msgTypes,
	}.Build()
	File_slots_slot_service_proto = out.File
	file_slots_slot_service_proto_rawDesc = nil
	file_slots_slot_service_proto_goTypes = nil
	file_slots_slot_service_proto_depIdxs = nil
}
