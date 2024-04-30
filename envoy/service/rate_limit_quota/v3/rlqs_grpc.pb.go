// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: envoy/service/rate_limit_quota/v3/rlqs.proto

package rate_limit_quotav3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RateLimitQuotaService_StreamRateLimitQuotas_FullMethodName = "/envoy.service.rate_limit_quota.v3.RateLimitQuotaService/StreamRateLimitQuotas"
)

// RateLimitQuotaServiceClient is the client API for RateLimitQuotaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateLimitQuotaServiceClient interface {
	// Main communication channel: the data plane sends usage reports to the RLQS server,
	// and the server asynchronously responding with the assignments.
	StreamRateLimitQuotas(ctx context.Context, opts ...grpc.CallOption) (RateLimitQuotaService_StreamRateLimitQuotasClient, error)
}

type rateLimitQuotaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimitQuotaServiceClient(cc grpc.ClientConnInterface) RateLimitQuotaServiceClient {
	return &rateLimitQuotaServiceClient{cc}
}

func (c *rateLimitQuotaServiceClient) StreamRateLimitQuotas(ctx context.Context, opts ...grpc.CallOption) (RateLimitQuotaService_StreamRateLimitQuotasClient, error) {
	stream, err := c.cc.NewStream(ctx, &RateLimitQuotaService_ServiceDesc.Streams[0], RateLimitQuotaService_StreamRateLimitQuotas_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitQuotaServiceStreamRateLimitQuotasClient{stream}
	return x, nil
}

type RateLimitQuotaService_StreamRateLimitQuotasClient interface {
	Send(*RateLimitQuotaUsageReports) error
	Recv() (*RateLimitQuotaResponse, error)
	grpc.ClientStream
}

type rateLimitQuotaServiceStreamRateLimitQuotasClient struct {
	grpc.ClientStream
}

func (x *rateLimitQuotaServiceStreamRateLimitQuotasClient) Send(m *RateLimitQuotaUsageReports) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitQuotaServiceStreamRateLimitQuotasClient) Recv() (*RateLimitQuotaResponse, error) {
	m := new(RateLimitQuotaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RateLimitQuotaServiceServer is the server API for RateLimitQuotaService service.
// All implementations should embed UnimplementedRateLimitQuotaServiceServer
// for forward compatibility
type RateLimitQuotaServiceServer interface {
	// Main communication channel: the data plane sends usage reports to the RLQS server,
	// and the server asynchronously responding with the assignments.
	StreamRateLimitQuotas(RateLimitQuotaService_StreamRateLimitQuotasServer) error
}

// UnimplementedRateLimitQuotaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRateLimitQuotaServiceServer struct {
}

func (UnimplementedRateLimitQuotaServiceServer) StreamRateLimitQuotas(RateLimitQuotaService_StreamRateLimitQuotasServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRateLimitQuotas not implemented")
}

// UnsafeRateLimitQuotaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateLimitQuotaServiceServer will
// result in compilation errors.
type UnsafeRateLimitQuotaServiceServer interface {
	mustEmbedUnimplementedRateLimitQuotaServiceServer()
}

func RegisterRateLimitQuotaServiceServer(s grpc.ServiceRegistrar, srv RateLimitQuotaServiceServer) {
	s.RegisterService(&RateLimitQuotaService_ServiceDesc, srv)
}

func _RateLimitQuotaService_StreamRateLimitQuotas_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitQuotaServiceServer).StreamRateLimitQuotas(&rateLimitQuotaServiceStreamRateLimitQuotasServer{stream})
}

type RateLimitQuotaService_StreamRateLimitQuotasServer interface {
	Send(*RateLimitQuotaResponse) error
	Recv() (*RateLimitQuotaUsageReports, error)
	grpc.ServerStream
}

type rateLimitQuotaServiceStreamRateLimitQuotasServer struct {
	grpc.ServerStream
}

func (x *rateLimitQuotaServiceStreamRateLimitQuotasServer) Send(m *RateLimitQuotaResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitQuotaServiceStreamRateLimitQuotasServer) Recv() (*RateLimitQuotaUsageReports, error) {
	m := new(RateLimitQuotaUsageReports)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RateLimitQuotaService_ServiceDesc is the grpc.ServiceDesc for RateLimitQuotaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateLimitQuotaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.rate_limit_quota.v3.RateLimitQuotaService",
	HandlerType: (*RateLimitQuotaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRateLimitQuotas",
			Handler:       _RateLimitQuotaService_StreamRateLimitQuotas_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/rate_limit_quota/v3/rlqs.proto",
}
