// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: source_message.proto

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

type Source struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId        string                 `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc          string                 `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags          string                 `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Source        string                 `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Params        string                 `protobuf:"bytes,7,opt,name=params,proto3" json:"params,omitempty"`
	Config        string                 `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	Link          int32                  `protobuf:"zigzag32,9,opt,name=link,proto3" json:"link,omitempty"`
	Status        int32                  `protobuf:"zigzag32,10,opt,name=status,proto3" json:"status,omitempty"`
	Created       int64                  `protobuf:"varint,11,opt,name=created,proto3" json:"created,omitempty"`
	Updated       int64                  `protobuf:"varint,12,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted       int64                  `protobuf:"varint,13,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Source) Reset() {
	*x = Source{}
	mi := &file_source_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_source_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_source_message_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Source) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Source) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Source) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Source) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Source) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Source) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *Source) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Source) GetLink() int32 {
	if x != nil {
		return x.Link
	}
	return 0
}

func (x *Source) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Source) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Source) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Source) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type Tag struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId        string                 `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	SourceId      string                 `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Desc          string                 `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags          string                 `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	DataType      string                 `protobuf:"bytes,7,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Address       string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Value         string                 `protobuf:"bytes,9,opt,name=value,proto3" json:"value,omitempty"`
	Config        string                 `protobuf:"bytes,10,opt,name=config,proto3" json:"config,omitempty"`
	Status        int32                  `protobuf:"zigzag32,11,opt,name=status,proto3" json:"status,omitempty"`
	Access        int32                  `protobuf:"zigzag32,12,opt,name=access,proto3" json:"access,omitempty"`
	Created       int64                  `protobuf:"varint,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated       int64                  `protobuf:"varint,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted       int64                  `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_source_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_source_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_source_message_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Tag) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Tag) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Tag) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Tag) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Tag) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Tag) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Tag) GetAccess() int32 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *Tag) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Tag) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Tag) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type TagValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Updated       int64                  `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TagValue) Reset() {
	*x = TagValue{}
	mi := &file_source_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagValue) ProtoMessage() {}

func (x *TagValue) ProtoReflect() protoreflect.Message {
	mi := &file_source_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagValue.ProtoReflect.Descriptor instead.
func (*TagValue) Descriptor() ([]byte, []int) {
	return file_source_message_proto_rawDescGZIP(), []int{2}
}

func (x *TagValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TagValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type TagNameValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value         string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Updated       int64                  `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TagNameValue) Reset() {
	*x = TagNameValue{}
	mi := &file_source_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagNameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagNameValue) ProtoMessage() {}

func (x *TagNameValue) ProtoReflect() protoreflect.Message {
	mi := &file_source_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagNameValue.ProtoReflect.Descriptor instead.
func (*TagNameValue) Descriptor() ([]byte, []int) {
	return file_source_message_proto_rawDescGZIP(), []int{3}
}

func (x *TagNameValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagNameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagNameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TagNameValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type TagValueUpdated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId        string                 `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	SourceId      string                 `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Value         string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Updated       int64                  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TagValueUpdated) Reset() {
	*x = TagValueUpdated{}
	mi := &file_source_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagValueUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagValueUpdated) ProtoMessage() {}

func (x *TagValueUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_source_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagValueUpdated.ProtoReflect.Descriptor instead.
func (*TagValueUpdated) Descriptor() ([]byte, []int) {
	return file_source_message_proto_rawDescGZIP(), []int{4}
}

func (x *TagValueUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagValueUpdated) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *TagValueUpdated) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *TagValueUpdated) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TagValueUpdated) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_source_message_proto protoreflect.FileDescriptor

var file_source_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xaf, 0x02, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xea, 0x02, 0x0a,
	0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x08, 0x54, 0x61, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x62, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x54, 0x61,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_message_proto_rawDescOnce sync.Once
	file_source_message_proto_rawDescData = file_source_message_proto_rawDesc
)

func file_source_message_proto_rawDescGZIP() []byte {
	file_source_message_proto_rawDescOnce.Do(func() {
		file_source_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_message_proto_rawDescData)
	})
	return file_source_message_proto_rawDescData
}

var file_source_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_source_message_proto_goTypes = []any{
	(*Source)(nil),          // 0: pb.Source
	(*Tag)(nil),             // 1: pb.Tag
	(*TagValue)(nil),        // 2: pb.TagValue
	(*TagNameValue)(nil),    // 3: pb.TagNameValue
	(*TagValueUpdated)(nil), // 4: pb.TagValueUpdated
}
var file_source_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_source_message_proto_init() }
func file_source_message_proto_init() {
	if File_source_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_source_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_message_proto_goTypes,
		DependencyIndexes: file_source_message_proto_depIdxs,
		MessageInfos:      file_source_message_proto_msgTypes,
	}.Build()
	File_source_message_proto = out.File
	file_source_message_proto_rawDesc = nil
	file_source_message_proto_goTypes = nil
	file_source_message_proto_depIdxs = nil
}
