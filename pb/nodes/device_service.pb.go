// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: nodes/device_service.proto

package nodes

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

type DeviceLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Access        string                 `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Secret        string                 `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceLoginRequest) Reset() {
	*x = DeviceLoginRequest{}
	mi := &file_nodes_device_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLoginRequest) ProtoMessage() {}

func (x *DeviceLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_device_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLoginRequest.ProtoReflect.Descriptor instead.
func (*DeviceLoginRequest) Descriptor() ([]byte, []int) {
	return file_nodes_device_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceLoginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceLoginRequest) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *DeviceLoginRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type DeviceLoginReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Device        *pb.Device             `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceLoginReply) Reset() {
	*x = DeviceLoginReply{}
	mi := &file_nodes_device_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLoginReply) ProtoMessage() {}

func (x *DeviceLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_device_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLoginReply.ProtoReflect.Descriptor instead.
func (*DeviceLoginReply) Descriptor() ([]byte, []int) {
	return file_nodes_device_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceLoginReply) GetDevice() *pb.Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *DeviceLoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeviceLinkRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// string id = 1;
	Status        int32 `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceLinkRequest) Reset() {
	*x = DeviceLinkRequest{}
	mi := &file_nodes_device_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLinkRequest) ProtoMessage() {}

func (x *DeviceLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_device_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLinkRequest.ProtoReflect.Descriptor instead.
func (*DeviceLinkRequest) Descriptor() ([]byte, []int) {
	return file_nodes_device_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceLinkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DeviceKeepAliveReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Time          int32                  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceKeepAliveReply) Reset() {
	*x = DeviceKeepAliveReply{}
	mi := &file_nodes_device_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceKeepAliveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceKeepAliveReply) ProtoMessage() {}

func (x *DeviceKeepAliveReply) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_device_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceKeepAliveReply.ProtoReflect.Descriptor instead.
func (*DeviceKeepAliveReply) Descriptor() ([]byte, []int) {
	return file_nodes_device_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceKeepAliveReply) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_nodes_device_service_proto protoreflect.FileDescriptor

var file_nodes_device_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x1a, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x54, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x65, 0x70, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xd0, 0x02,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x22,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x00, 0x12, 0x21, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42,
	0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42,
	0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65,
	0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nodes_device_service_proto_rawDescOnce sync.Once
	file_nodes_device_service_proto_rawDescData = file_nodes_device_service_proto_rawDesc
)

func file_nodes_device_service_proto_rawDescGZIP() []byte {
	file_nodes_device_service_proto_rawDescOnce.Do(func() {
		file_nodes_device_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodes_device_service_proto_rawDescData)
	})
	return file_nodes_device_service_proto_rawDescData
}

var file_nodes_device_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodes_device_service_proto_goTypes = []any{
	(*DeviceLoginRequest)(nil),   // 0: nodes.DeviceLoginRequest
	(*DeviceLoginReply)(nil),     // 1: nodes.DeviceLoginReply
	(*DeviceLinkRequest)(nil),    // 2: nodes.DeviceLinkRequest
	(*DeviceKeepAliveReply)(nil), // 3: nodes.DeviceKeepAliveReply
	(*pb.Device)(nil),            // 4: pb.Device
	(*pb.MyEmpty)(nil),           // 5: pb.MyEmpty
	(*pb.MyBool)(nil),            // 6: pb.MyBool
}
var file_nodes_device_service_proto_depIdxs = []int32{
	4, // 0: nodes.DeviceLoginReply.device:type_name -> pb.Device
	0, // 1: nodes.DeviceService.Login:input_type -> nodes.DeviceLoginRequest
	4, // 2: nodes.DeviceService.Update:input_type -> pb.Device
	5, // 3: nodes.DeviceService.View:input_type -> pb.MyEmpty
	2, // 4: nodes.DeviceService.Link:input_type -> nodes.DeviceLinkRequest
	5, // 5: nodes.DeviceService.ViewWithDeleted:input_type -> pb.MyEmpty
	4, // 6: nodes.DeviceService.Sync:input_type -> pb.Device
	5, // 7: nodes.DeviceService.KeepAlive:input_type -> pb.MyEmpty
	1, // 8: nodes.DeviceService.Login:output_type -> nodes.DeviceLoginReply
	4, // 9: nodes.DeviceService.Update:output_type -> pb.Device
	4, // 10: nodes.DeviceService.View:output_type -> pb.Device
	6, // 11: nodes.DeviceService.Link:output_type -> pb.MyBool
	4, // 12: nodes.DeviceService.ViewWithDeleted:output_type -> pb.Device
	6, // 13: nodes.DeviceService.Sync:output_type -> pb.MyBool
	3, // 14: nodes.DeviceService.KeepAlive:output_type -> nodes.DeviceKeepAliveReply
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nodes_device_service_proto_init() }
func file_nodes_device_service_proto_init() {
	if File_nodes_device_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodes_device_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodes_device_service_proto_goTypes,
		DependencyIndexes: file_nodes_device_service_proto_depIdxs,
		MessageInfos:      file_nodes_device_service_proto_msgTypes,
	}.Build()
	File_nodes_device_service_proto = out.File
	file_nodes_device_service_proto_rawDesc = nil
	file_nodes_device_service_proto_goTypes = nil
	file_nodes_device_service_proto_depIdxs = nil
}
