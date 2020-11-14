// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package acl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WatchServiceClient is the client API for WatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchServiceClient interface {
	// Watches and filters for changes in the ACL system.
	WatchRelationTuples(ctx context.Context, in *WatchRelationTuplesRequest, opts ...grpc.CallOption) (WatchService_WatchRelationTuplesClient, error)
}

type watchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchServiceClient(cc grpc.ClientConnInterface) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) WatchRelationTuples(ctx context.Context, in *WatchRelationTuplesRequest, opts ...grpc.CallOption) (WatchService_WatchRelationTuplesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WatchService_serviceDesc.Streams[0], "/keto.acl.v1.WatchService/WatchRelationTuples", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchRelationTuplesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchService_WatchRelationTuplesClient interface {
	Recv() (*WatchRelationTuplesResponse, error)
	grpc.ClientStream
}

type watchServiceWatchRelationTuplesClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchRelationTuplesClient) Recv() (*WatchRelationTuplesResponse, error) {
	m := new(WatchRelationTuplesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServiceServer is the server API for WatchService service.
// All implementations should embed UnimplementedWatchServiceServer
// for forward compatibility
type WatchServiceServer interface {
	// Watches and filters for changes in the ACL system.
	WatchRelationTuples(*WatchRelationTuplesRequest, WatchService_WatchRelationTuplesServer) error
}

// UnimplementedWatchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWatchServiceServer struct {
}

func (UnimplementedWatchServiceServer) WatchRelationTuples(*WatchRelationTuplesRequest, WatchService_WatchRelationTuplesServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRelationTuples not implemented")
}

// UnsafeWatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServiceServer will
// result in compilation errors.
type UnsafeWatchServiceServer interface {
	mustEmbedUnimplementedWatchServiceServer()
}

func RegisterWatchServiceServer(s grpc.ServiceRegistrar, srv WatchServiceServer) {
	s.RegisterService(&_WatchService_serviceDesc, srv)
}

func _WatchService_WatchRelationTuples_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRelationTuplesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchServiceServer).WatchRelationTuples(m, &watchServiceWatchRelationTuplesServer{stream})
}

type WatchService_WatchRelationTuplesServer interface {
	Send(*WatchRelationTuplesResponse) error
	grpc.ServerStream
}

type watchServiceWatchRelationTuplesServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchRelationTuplesServer) Send(m *WatchRelationTuplesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _WatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keto.acl.v1.WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchRelationTuples",
			Handler:       _WatchService_WatchRelationTuples_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "keto/acl/v1/watch_service.proto",
}
