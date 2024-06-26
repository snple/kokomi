// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: edges/data_service.proto

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

type DataUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContentType int32  `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Content     []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// string device_id = 4;
	Realtime bool `protobuf:"varint,5,opt,name=realtime,proto3" json:"realtime,omitempty"`
	Save     bool `protobuf:"varint,6,opt,name=save,proto3" json:"save,omitempty"`
}

func (x *DataUploadRequest) Reset() {
	*x = DataUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUploadRequest) ProtoMessage() {}

func (x *DataUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUploadRequest.ProtoReflect.Descriptor instead.
func (*DataUploadRequest) Descriptor() ([]byte, []int) {
	return file_edges_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *DataUploadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataUploadRequest) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *DataUploadRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *DataUploadRequest) GetRealtime() bool {
	if x != nil {
		return x.Realtime
	}
	return false
}

func (x *DataUploadRequest) GetSave() bool {
	if x != nil {
		return x.Save
	}
	return false
}

type DataUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DataUploadResponse) Reset() {
	*x = DataUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUploadResponse) ProtoMessage() {}

func (x *DataUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edges_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUploadResponse.ProtoReflect.Descriptor instead.
func (*DataUploadResponse) Descriptor() ([]byte, []int) {
	return file_edges_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *DataUploadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataUploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DataQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flux string            `protobuf:"bytes,1,opt,name=flux,proto3" json:"flux,omitempty"`
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataQueryRequest) Reset() {
	*x = DataQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_data_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataQueryRequest) ProtoMessage() {}

func (x *DataQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_data_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataQueryRequest.ProtoReflect.Descriptor instead.
func (*DataQueryRequest) Descriptor() ([]byte, []int) {
	return file_edges_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *DataQueryRequest) GetFlux() string {
	if x != nil {
		return x.Flux
	}
	return ""
}

func (x *DataQueryRequest) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

type DataQueryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataQueryByIdRequest) Reset() {
	*x = DataQueryByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_data_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataQueryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataQueryByIdRequest) ProtoMessage() {}

func (x *DataQueryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_data_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataQueryByIdRequest.ProtoReflect.Descriptor instead.
func (*DataQueryByIdRequest) Descriptor() ([]byte, []int) {
	return file_edges_data_service_proto_rawDescGZIP(), []int{3}
}

func (x *DataQueryByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataQueryByIdRequest) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

var File_edges_data_service_proto protoreflect.FileDescriptor

var file_edges_data_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74,
	0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x61, 0x76, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x44,
	0x61, 0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x10,
	0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x75, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x6c, 0x75, 0x78, 0x12, 0x35, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56,
	0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0xef, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70,
	0x62, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x73, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edges_data_service_proto_rawDescOnce sync.Once
	file_edges_data_service_proto_rawDescData = file_edges_data_service_proto_rawDesc
)

func file_edges_data_service_proto_rawDescGZIP() []byte {
	file_edges_data_service_proto_rawDescOnce.Do(func() {
		file_edges_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_edges_data_service_proto_rawDescData)
	})
	return file_edges_data_service_proto_rawDescData
}

var file_edges_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_edges_data_service_proto_goTypes = []interface{}{
	(*DataUploadRequest)(nil),    // 0: edges.DataUploadRequest
	(*DataUploadResponse)(nil),   // 1: edges.DataUploadResponse
	(*DataQueryRequest)(nil),     // 2: edges.DataQueryRequest
	(*DataQueryByIdRequest)(nil), // 3: edges.DataQueryByIdRequest
	nil,                          // 4: edges.DataQueryRequest.VarsEntry
	nil,                          // 5: edges.DataQueryByIdRequest.VarsEntry
	(*pb.Message)(nil),           // 6: pb.Message
}
var file_edges_data_service_proto_depIdxs = []int32{
	4, // 0: edges.DataQueryRequest.vars:type_name -> edges.DataQueryRequest.VarsEntry
	5, // 1: edges.DataQueryByIdRequest.vars:type_name -> edges.DataQueryByIdRequest.VarsEntry
	0, // 2: edges.DataService.Upload:input_type -> edges.DataUploadRequest
	2, // 3: edges.DataService.Compile:input_type -> edges.DataQueryRequest
	2, // 4: edges.DataService.Query:input_type -> edges.DataQueryRequest
	3, // 5: edges.DataService.QueryById:input_type -> edges.DataQueryByIdRequest
	1, // 6: edges.DataService.Upload:output_type -> edges.DataUploadResponse
	6, // 7: edges.DataService.Compile:output_type -> pb.Message
	6, // 8: edges.DataService.Query:output_type -> pb.Message
	6, // 9: edges.DataService.QueryById:output_type -> pb.Message
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_edges_data_service_proto_init() }
func file_edges_data_service_proto_init() {
	if File_edges_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_edges_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUploadRequest); i {
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
		file_edges_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUploadResponse); i {
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
		file_edges_data_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataQueryRequest); i {
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
		file_edges_data_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataQueryByIdRequest); i {
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
			RawDescriptor: file_edges_data_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edges_data_service_proto_goTypes,
		DependencyIndexes: file_edges_data_service_proto_depIdxs,
		MessageInfos:      file_edges_data_service_proto_msgTypes,
	}.Build()
	File_edges_data_service_proto = out.File
	file_edges_data_service_proto_rawDesc = nil
	file_edges_data_service_proto_goTypes = nil
	file_edges_data_service_proto_depIdxs = nil
}
