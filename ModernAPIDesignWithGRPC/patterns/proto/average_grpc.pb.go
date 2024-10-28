// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: proto/average.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AverageService_FindAverage_FullMethodName = "/proto.AverageService/FindAverage"
)

// AverageServiceClient is the client API for AverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AverageServiceClient interface {
	FindAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error)
}

type averageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAverageServiceClient(cc grpc.ClientConnInterface) AverageServiceClient {
	return &averageServiceClient{cc}
}

func (c *averageServiceClient) FindAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AverageService_ServiceDesc.Streams[0], AverageService_FindAverage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AverageRequest, AverageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AverageService_FindAverageClient = grpc.ClientStreamingClient[AverageRequest, AverageResponse]

// AverageServiceServer is the server API for AverageService service.
// All implementations must embed UnimplementedAverageServiceServer
// for forward compatibility.
type AverageServiceServer interface {
	FindAverage(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error
	mustEmbedUnimplementedAverageServiceServer()
}

// UnimplementedAverageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAverageServiceServer struct{}

func (UnimplementedAverageServiceServer) FindAverage(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FindAverage not implemented")
}
func (UnimplementedAverageServiceServer) mustEmbedUnimplementedAverageServiceServer() {}
func (UnimplementedAverageServiceServer) testEmbeddedByValue()                        {}

// UnsafeAverageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AverageServiceServer will
// result in compilation errors.
type UnsafeAverageServiceServer interface {
	mustEmbedUnimplementedAverageServiceServer()
}

func RegisterAverageServiceServer(s grpc.ServiceRegistrar, srv AverageServiceServer) {
	// If the following call pancis, it indicates UnimplementedAverageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AverageService_ServiceDesc, srv)
}

func _AverageService_FindAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AverageServiceServer).FindAverage(&grpc.GenericServerStream[AverageRequest, AverageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AverageService_FindAverageServer = grpc.ClientStreamingServer[AverageRequest, AverageResponse]

// AverageService_ServiceDesc is the grpc.ServiceDesc for AverageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AverageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AverageService",
	HandlerType: (*AverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAverage",
			Handler:       _AverageService_FindAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/average.proto",
}
