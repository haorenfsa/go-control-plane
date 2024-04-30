// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: envoy/extensions/filters/network/thrift_proxy/router/v3/router.proto

package routerv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Close downstream connection in case of routing or upstream connection problem. Default: true
	CloseDownstreamOnUpstreamError *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=close_downstream_on_upstream_error,json=closeDownstreamOnUpstreamError,proto3" json:"close_downstream_on_upstream_error,omitempty"`
}

func (x *Router) Reset() {
	*x = Router{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescGZIP(), []int{0}
}

func (x *Router) GetCloseDownstreamOnUpstreamError() *wrapperspb.BoolValue {
	if x != nil {
		return x.CloseDownstreamOnUpstreamError
	}
	return nil
}

var File_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDesc = []byte{
	0x0a, 0x44, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x22,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x1e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x6e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x3a, 0x38, 0x9a, 0xc5, 0x88, 0x1e, 0x33, 0x0a, 0x31, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x42, 0xc7,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x45, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescData = file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDesc
)

func file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDescData
}

var file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_goTypes = []interface{}{
	(*Router)(nil),               // 0: envoy.extensions.filters.network.thrift_proxy.router.v3.Router
	(*wrapperspb.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.thrift_proxy.router.v3.Router.close_downstream_on_upstream_error:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_init() }
func file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_init() {
	if File_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Router); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto = out.File
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_rawDesc = nil
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_goTypes = nil
	file_envoy_extensions_filters_network_thrift_proxy_router_v3_router_proto_depIdxs = nil
}
