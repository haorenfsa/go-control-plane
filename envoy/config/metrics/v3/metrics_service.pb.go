// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: envoy/config/metrics/v3/metrics_service.proto

package metricsv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// HistogramEmitMode is used to configure which metric types should be emitted for histograms.
type HistogramEmitMode int32

const (
	// Emit Histogram and Summary metric types.
	HistogramEmitMode_SUMMARY_AND_HISTOGRAM HistogramEmitMode = 0
	// Emit only Summary metric types.
	HistogramEmitMode_SUMMARY HistogramEmitMode = 1
	// Emit only Histogram metric types.
	HistogramEmitMode_HISTOGRAM HistogramEmitMode = 2
)

// Enum value maps for HistogramEmitMode.
var (
	HistogramEmitMode_name = map[int32]string{
		0: "SUMMARY_AND_HISTOGRAM",
		1: "SUMMARY",
		2: "HISTOGRAM",
	}
	HistogramEmitMode_value = map[string]int32{
		"SUMMARY_AND_HISTOGRAM": 0,
		"SUMMARY":               1,
		"HISTOGRAM":             2,
	}
)

func (x HistogramEmitMode) Enum() *HistogramEmitMode {
	p := new(HistogramEmitMode)
	*p = x
	return p
}

func (x HistogramEmitMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HistogramEmitMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_metrics_v3_metrics_service_proto_enumTypes[0].Descriptor()
}

func (HistogramEmitMode) Type() protoreflect.EnumType {
	return &file_envoy_config_metrics_v3_metrics_service_proto_enumTypes[0]
}

func (x HistogramEmitMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HistogramEmitMode.Descriptor instead.
func (HistogramEmitMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v3_metrics_service_proto_rawDescGZIP(), []int{0}
}

// Metrics Service is configured as a built-in “envoy.stat_sinks.metrics_service“ :ref:`StatsSink
// <envoy_v3_api_msg_config.metrics.v3.StatsSink>`. This opaque configuration will be used to create
// Metrics Service.
//
// Example:
//
// .. code-block:: yaml
//
//	stats_sinks:
//	  - name: envoy.stat_sinks.metrics_service
//	    typed_config:
//	      "@type": type.googleapis.com/envoy.config.metrics.v3.MetricsServiceConfig
//	      transport_api_version: V3
//
// [#extension: envoy.stat_sinks.metrics_service]
// [#next-free-field: 6]
type MetricsServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// API version for metric service transport protocol. This describes the metric service gRPC
	// endpoint and version of messages used on the wire.
	TransportApiVersion v3.ApiVersion `protobuf:"varint,3,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.config.core.v3.ApiVersion" json:"transport_api_version,omitempty"`
	// If true, counters are reported as the delta between flushing intervals. Otherwise, the current
	// counter value is reported. Defaults to false.
	// Eventually (https://github.com/envoyproxy/envoy/issues/10968) if this value is not set, the
	// sink will take updates from the :ref:`MetricsResponse <envoy_v3_api_msg_service.metrics.v3.StreamMetricsResponse>`.
	ReportCountersAsDeltas *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=report_counters_as_deltas,json=reportCountersAsDeltas,proto3" json:"report_counters_as_deltas,omitempty"`
	// If true, metrics will have their tags emitted as labels on the metrics objects sent to the MetricsService,
	// and the tag extracted name will be used instead of the full name, which may contain values used by the tag
	// extractor or additional tags added during stats creation.
	EmitTagsAsLabels bool `protobuf:"varint,4,opt,name=emit_tags_as_labels,json=emitTagsAsLabels,proto3" json:"emit_tags_as_labels,omitempty"`
	// Specify which metrics types to emit for histograms. Defaults to SUMMARY_AND_HISTOGRAM.
	HistogramEmitMode HistogramEmitMode `protobuf:"varint,5,opt,name=histogram_emit_mode,json=histogramEmitMode,proto3,enum=envoy.config.metrics.v3.HistogramEmitMode" json:"histogram_emit_mode,omitempty"`
}

func (x *MetricsServiceConfig) Reset() {
	*x = MetricsServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v3_metrics_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsServiceConfig) ProtoMessage() {}

func (x *MetricsServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v3_metrics_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsServiceConfig.ProtoReflect.Descriptor instead.
func (*MetricsServiceConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v3_metrics_service_proto_rawDescGZIP(), []int{0}
}

func (x *MetricsServiceConfig) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *MetricsServiceConfig) GetTransportApiVersion() v3.ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return v3.ApiVersion(0)
}

func (x *MetricsServiceConfig) GetReportCountersAsDeltas() *wrapperspb.BoolValue {
	if x != nil {
		return x.ReportCountersAsDeltas
	}
	return nil
}

func (x *MetricsServiceConfig) GetEmitTagsAsLabels() bool {
	if x != nil {
		return x.EmitTagsAsLabels
	}
	return false
}

func (x *MetricsServiceConfig) GetHistogramEmitMode() HistogramEmitMode {
	if x != nil {
		return x.HistogramEmitMode
	}
	return HistogramEmitMode_SUMMARY_AND_HISTOGRAM
}

var File_envoy_config_metrics_v3_metrics_service_proto protoreflect.FileDescriptor

var file_envoy_config_metrics_v3_metrics_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x03, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x4e, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x55, 0x0a, 0x19, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x41, 0x73,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x65, 0x6d, 0x69, 0x74, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x67, 0x73, 0x41, 0x73, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x64, 0x0a, 0x13, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x5f, 0x65, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6d, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x11, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x45, 0x6d, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x33, 0x9a, 0xc5, 0x88,
	0x1e, 0x2e, 0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2a, 0x4a, 0x0a, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6d, 0x69,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x42, 0x90, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x13,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_metrics_v3_metrics_service_proto_rawDescOnce sync.Once
	file_envoy_config_metrics_v3_metrics_service_proto_rawDescData = file_envoy_config_metrics_v3_metrics_service_proto_rawDesc
)

func file_envoy_config_metrics_v3_metrics_service_proto_rawDescGZIP() []byte {
	file_envoy_config_metrics_v3_metrics_service_proto_rawDescOnce.Do(func() {
		file_envoy_config_metrics_v3_metrics_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_metrics_v3_metrics_service_proto_rawDescData)
	})
	return file_envoy_config_metrics_v3_metrics_service_proto_rawDescData
}

var file_envoy_config_metrics_v3_metrics_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_metrics_v3_metrics_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_metrics_v3_metrics_service_proto_goTypes = []interface{}{
	(HistogramEmitMode)(0),       // 0: envoy.config.metrics.v3.HistogramEmitMode
	(*MetricsServiceConfig)(nil), // 1: envoy.config.metrics.v3.MetricsServiceConfig
	(*v3.GrpcService)(nil),       // 2: envoy.config.core.v3.GrpcService
	(v3.ApiVersion)(0),           // 3: envoy.config.core.v3.ApiVersion
	(*wrapperspb.BoolValue)(nil), // 4: google.protobuf.BoolValue
}
var file_envoy_config_metrics_v3_metrics_service_proto_depIdxs = []int32{
	2, // 0: envoy.config.metrics.v3.MetricsServiceConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	3, // 1: envoy.config.metrics.v3.MetricsServiceConfig.transport_api_version:type_name -> envoy.config.core.v3.ApiVersion
	4, // 2: envoy.config.metrics.v3.MetricsServiceConfig.report_counters_as_deltas:type_name -> google.protobuf.BoolValue
	0, // 3: envoy.config.metrics.v3.MetricsServiceConfig.histogram_emit_mode:type_name -> envoy.config.metrics.v3.HistogramEmitMode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_metrics_v3_metrics_service_proto_init() }
func file_envoy_config_metrics_v3_metrics_service_proto_init() {
	if File_envoy_config_metrics_v3_metrics_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_metrics_v3_metrics_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsServiceConfig); i {
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
			RawDescriptor: file_envoy_config_metrics_v3_metrics_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_metrics_v3_metrics_service_proto_goTypes,
		DependencyIndexes: file_envoy_config_metrics_v3_metrics_service_proto_depIdxs,
		EnumInfos:         file_envoy_config_metrics_v3_metrics_service_proto_enumTypes,
		MessageInfos:      file_envoy_config_metrics_v3_metrics_service_proto_msgTypes,
	}.Build()
	File_envoy_config_metrics_v3_metrics_service_proto = out.File
	file_envoy_config_metrics_v3_metrics_service_proto_rawDesc = nil
	file_envoy_config_metrics_v3_metrics_service_proto_goTypes = nil
	file_envoy_config_metrics_v3_metrics_service_proto_depIdxs = nil
}
