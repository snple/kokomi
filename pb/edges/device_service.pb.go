// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: edges/device_service.proto

package edges

import (
	pb "github.com/snple/kokomi/pb"
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

var File_edges_device_service_proto protoreflect.FileDescriptor

var file_edges_device_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x1a, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xcc, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x07, 0x44, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x20, 0x0a,
	0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e,
	0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_edges_device_service_proto_goTypes = []interface{}{
	(*pb.Device)(nil),  // 0: pb.Device
	(*pb.MyEmpty)(nil), // 1: pb.MyEmpty
	(*pb.MyBool)(nil),  // 2: pb.MyBool
}
var file_edges_device_service_proto_depIdxs = []int32{
	0, // 0: edges.DeviceService.Update:input_type -> pb.Device
	1, // 1: edges.DeviceService.View:input_type -> pb.MyEmpty
	1, // 2: edges.DeviceService.Destory:input_type -> pb.MyEmpty
	1, // 3: edges.DeviceService.ViewWithDeleted:input_type -> pb.MyEmpty
	0, // 4: edges.DeviceService.Sync:input_type -> pb.Device
	0, // 5: edges.DeviceService.Update:output_type -> pb.Device
	0, // 6: edges.DeviceService.View:output_type -> pb.Device
	2, // 7: edges.DeviceService.Destory:output_type -> pb.MyBool
	0, // 8: edges.DeviceService.ViewWithDeleted:output_type -> pb.Device
	2, // 9: edges.DeviceService.Sync:output_type -> pb.MyBool
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_edges_device_service_proto_init() }
func file_edges_device_service_proto_init() {
	if File_edges_device_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_edges_device_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edges_device_service_proto_goTypes,
		DependencyIndexes: file_edges_device_service_proto_depIdxs,
	}.Build()
	File_edges_device_service_proto = out.File
	file_edges_device_service_proto_rawDesc = nil
	file_edges_device_service_proto_goTypes = nil
	file_edges_device_service_proto_depIdxs = nil
}
