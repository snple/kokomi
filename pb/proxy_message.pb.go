// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: proxy_message.proto

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

type Proxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Network  string `protobuf:"bytes,7,opt,name=network,proto3" json:"network,omitempty"`
	Address  string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Target   string `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty"`
	Config   string `protobuf:"bytes,10,opt,name=config,proto3" json:"config,omitempty"`
	Status   int32  `protobuf:"zigzag32,11,opt,name=status,proto3" json:"status,omitempty"`
	Link     int32  `protobuf:"zigzag32,12,opt,name=link,proto3" json:"link,omitempty"`
	Created  int64  `protobuf:"varint,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Proxy) Reset() {
	*x = Proxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proxy) ProtoMessage() {}

func (x *Proxy) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proxy.ProtoReflect.Descriptor instead.
func (*Proxy) Descriptor() ([]byte, []int) {
	return file_proxy_message_proto_rawDescGZIP(), []int{0}
}

func (x *Proxy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Proxy) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Proxy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Proxy) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Proxy) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Proxy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Proxy) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Proxy) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Proxy) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Proxy) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Proxy) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Proxy) GetLink() int32 {
	if x != nil {
		return x.Link
	}
	return 0
}

func (x *Proxy) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Proxy) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Proxy) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_proxy_message_proto protoreflect.FileDescriptor

var file_proxy_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe2, 0x02, 0x0a, 0x05, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70,
	0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_message_proto_rawDescOnce sync.Once
	file_proxy_message_proto_rawDescData = file_proxy_message_proto_rawDesc
)

func file_proxy_message_proto_rawDescGZIP() []byte {
	file_proxy_message_proto_rawDescOnce.Do(func() {
		file_proxy_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_message_proto_rawDescData)
	})
	return file_proxy_message_proto_rawDescData
}

var file_proxy_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_message_proto_goTypes = []interface{}{
	(*Proxy)(nil), // 0: pb.Proxy
}
var file_proxy_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_message_proto_init() }
func file_proxy_message_proto_init() {
	if File_proxy_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proxy); i {
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
			RawDescriptor: file_proxy_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_message_proto_goTypes,
		DependencyIndexes: file_proxy_message_proto_depIdxs,
		MessageInfos:      file_proxy_message_proto_msgTypes,
	}.Build()
	File_proxy_message_proto = out.File
	file_proxy_message_proto_rawDesc = nil
	file_proxy_message_proto_goTypes = nil
	file_proxy_message_proto_depIdxs = nil
}
