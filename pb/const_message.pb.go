// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: const_message.proto

package pb

import (
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

type Const struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId      string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc          string                 `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags          string                 `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type          string                 `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	DataType      string                 `protobuf:"bytes,7,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Value         string                 `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	HValue        string                 `protobuf:"bytes,9,opt,name=h_value,json=hValue,proto3" json:"h_value,omitempty"`
	LValue        string                 `protobuf:"bytes,10,opt,name=l_value,json=lValue,proto3" json:"l_value,omitempty"`
	Config        string                 `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	Status        int32                  `protobuf:"zigzag32,12,opt,name=status,proto3" json:"status,omitempty"`
	Created       int64                  `protobuf:"varint,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated       int64                  `protobuf:"varint,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted       int64                  `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Const) Reset() {
	*x = Const{}
	mi := &file_const_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Const) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Const) ProtoMessage() {}

func (x *Const) ProtoReflect() protoreflect.Message {
	mi := &file_const_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Const.ProtoReflect.Descriptor instead.
func (*Const) Descriptor() ([]byte, []int) {
	return file_const_message_proto_rawDescGZIP(), []int{0}
}

func (x *Const) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Const) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Const) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Const) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Const) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Const) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Const) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Const) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Const) GetHValue() string {
	if x != nil {
		return x.HValue
	}
	return ""
}

func (x *Const) GetLValue() string {
	if x != nil {
		return x.LValue
	}
	return ""
}

func (x *Const) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Const) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Const) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Const) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Const) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type ConstValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Updated       int64                  `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConstValue) Reset() {
	*x = ConstValue{}
	mi := &file_const_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConstValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstValue) ProtoMessage() {}

func (x *ConstValue) ProtoReflect() protoreflect.Message {
	mi := &file_const_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstValue.ProtoReflect.Descriptor instead.
func (*ConstValue) Descriptor() ([]byte, []int) {
	return file_const_message_proto_rawDescGZIP(), []int{1}
}

func (x *ConstValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConstValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConstValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type ConstNameValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value         string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Updated       int64                  `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConstNameValue) Reset() {
	*x = ConstNameValue{}
	mi := &file_const_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConstNameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstNameValue) ProtoMessage() {}

func (x *ConstNameValue) ProtoReflect() protoreflect.Message {
	mi := &file_const_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstNameValue.ProtoReflect.Descriptor instead.
func (*ConstNameValue) Descriptor() ([]byte, []int) {
	return file_const_message_proto_rawDescGZIP(), []int{2}
}

func (x *ConstNameValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConstNameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConstNameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConstNameValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_const_message_proto protoreflect.FileDescriptor

var file_const_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe7, 0x02, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x22, 0x4c, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x64, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x63,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_const_message_proto_rawDescOnce sync.Once
	file_const_message_proto_rawDescData = file_const_message_proto_rawDesc
)

func file_const_message_proto_rawDescGZIP() []byte {
	file_const_message_proto_rawDescOnce.Do(func() {
		file_const_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_const_message_proto_rawDescData)
	})
	return file_const_message_proto_rawDescData
}

var file_const_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_const_message_proto_goTypes = []any{
	(*Const)(nil),          // 0: pb.Const
	(*ConstValue)(nil),     // 1: pb.ConstValue
	(*ConstNameValue)(nil), // 2: pb.ConstNameValue
}
var file_const_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_const_message_proto_init() }
func file_const_message_proto_init() {
	if File_const_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_const_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_const_message_proto_goTypes,
		DependencyIndexes: file_const_message_proto_depIdxs,
		MessageInfos:      file_const_message_proto_msgTypes,
	}.Build()
	File_const_message_proto = out.File
	file_const_message_proto_rawDesc = nil
	file_const_message_proto_goTypes = nil
	file_const_message_proto_depIdxs = nil
}
