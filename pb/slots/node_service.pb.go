// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: slots/node_service.proto

package slots

import (
	pb "github.com/snple/beacon/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_slots_node_service_proto protoreflect.FileDescriptor

var file_slots_node_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6c, 0x6f, 0x74,
	0x73, 0x1a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4e, 0x0a, 0x0b,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65,
	0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x73,
	0x3b, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_slots_node_service_proto_goTypes = []any{
	(*pb.Node)(nil),    // 0: pb.Node
	(*pb.MyEmpty)(nil), // 1: pb.MyEmpty
}
var file_slots_node_service_proto_depIdxs = []int32{
	0, // 0: slots.NodeService.Update:input_type -> pb.Node
	1, // 1: slots.NodeService.View:input_type -> pb.MyEmpty
	0, // 2: slots.NodeService.Update:output_type -> pb.Node
	0, // 3: slots.NodeService.View:output_type -> pb.Node
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_slots_node_service_proto_init() }
func file_slots_node_service_proto_init() {
	if File_slots_node_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_slots_node_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slots_node_service_proto_goTypes,
		DependencyIndexes: file_slots_node_service_proto_depIdxs,
	}.Build()
	File_slots_node_service_proto = out.File
	file_slots_node_service_proto_rawDesc = nil
	file_slots_node_service_proto_goTypes = nil
	file_slots_node_service_proto_depIdxs = nil
}
