// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.0
// source: gsm.proto

package gsm

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

// GoPassSecretsManagerClient is the client API for GoPassSecretsManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoPassSecretsManagerClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (GoPassSecretsManager_SyncClient, error)
}

type goPassSecretsManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGoPassSecretsManagerClient(cc grpc.ClientConnInterface) GoPassSecretsManagerClient {
	return &goPassSecretsManagerClient{cc}
}

func (c *goPassSecretsManagerClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (GoPassSecretsManager_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoPassSecretsManager_ServiceDesc.Streams[0], "/gsm.GoPassSecretsManager/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &goPassSecretsManagerSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoPassSecretsManager_SyncClient interface {
	Recv() (*SyncResponse, error)
	grpc.ClientStream
}

type goPassSecretsManagerSyncClient struct {
	grpc.ClientStream
}

func (x *goPassSecretsManagerSyncClient) Recv() (*SyncResponse, error) {
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoPassSecretsManagerServer is the server API for GoPassSecretsManager service.
// All implementations must embed UnimplementedGoPassSecretsManagerServer
// for forward compatibility
type GoPassSecretsManagerServer interface {
	Sync(*SyncRequest, GoPassSecretsManager_SyncServer) error
	mustEmbedUnimplementedGoPassSecretsManagerServer()
}

// UnimplementedGoPassSecretsManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGoPassSecretsManagerServer struct {
}

func (UnimplementedGoPassSecretsManagerServer) Sync(*SyncRequest, GoPassSecretsManager_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedGoPassSecretsManagerServer) mustEmbedUnimplementedGoPassSecretsManagerServer() {}

// UnsafeGoPassSecretsManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoPassSecretsManagerServer will
// result in compilation errors.
type UnsafeGoPassSecretsManagerServer interface {
	mustEmbedUnimplementedGoPassSecretsManagerServer()
}

func RegisterGoPassSecretsManagerServer(s grpc.ServiceRegistrar, srv GoPassSecretsManagerServer) {
	s.RegisterService(&GoPassSecretsManager_ServiceDesc, srv)
}

func _GoPassSecretsManager_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoPassSecretsManagerServer).Sync(m, &goPassSecretsManagerSyncServer{stream})
}

type GoPassSecretsManager_SyncServer interface {
	Send(*SyncResponse) error
	grpc.ServerStream
}

type goPassSecretsManagerSyncServer struct {
	grpc.ServerStream
}

func (x *goPassSecretsManagerSyncServer) Send(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GoPassSecretsManager_ServiceDesc is the grpc.ServiceDesc for GoPassSecretsManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoPassSecretsManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsm.GoPassSecretsManager",
	HandlerType: (*GoPassSecretsManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _GoPassSecretsManager_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gsm.proto",
}

// GoPassSecretsManagerHookClient is the client API for GoPassSecretsManagerHook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoPassSecretsManagerHookClient interface {
	PostHook(ctx context.Context, in *PostHookRequest, opts ...grpc.CallOption) (GoPassSecretsManagerHook_PostHookClient, error)
}

type goPassSecretsManagerHookClient struct {
	cc grpc.ClientConnInterface
}

func NewGoPassSecretsManagerHookClient(cc grpc.ClientConnInterface) GoPassSecretsManagerHookClient {
	return &goPassSecretsManagerHookClient{cc}
}

func (c *goPassSecretsManagerHookClient) PostHook(ctx context.Context, in *PostHookRequest, opts ...grpc.CallOption) (GoPassSecretsManagerHook_PostHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoPassSecretsManagerHook_ServiceDesc.Streams[0], "/gsm.GoPassSecretsManagerHook/PostHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &goPassSecretsManagerHookPostHookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoPassSecretsManagerHook_PostHookClient interface {
	Recv() (*PostHookResponse, error)
	grpc.ClientStream
}

type goPassSecretsManagerHookPostHookClient struct {
	grpc.ClientStream
}

func (x *goPassSecretsManagerHookPostHookClient) Recv() (*PostHookResponse, error) {
	m := new(PostHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoPassSecretsManagerHookServer is the server API for GoPassSecretsManagerHook service.
// All implementations must embed UnimplementedGoPassSecretsManagerHookServer
// for forward compatibility
type GoPassSecretsManagerHookServer interface {
	PostHook(*PostHookRequest, GoPassSecretsManagerHook_PostHookServer) error
	mustEmbedUnimplementedGoPassSecretsManagerHookServer()
}

// UnimplementedGoPassSecretsManagerHookServer must be embedded to have forward compatible implementations.
type UnimplementedGoPassSecretsManagerHookServer struct {
}

func (UnimplementedGoPassSecretsManagerHookServer) PostHook(*PostHookRequest, GoPassSecretsManagerHook_PostHookServer) error {
	return status.Errorf(codes.Unimplemented, "method PostHook not implemented")
}
func (UnimplementedGoPassSecretsManagerHookServer) mustEmbedUnimplementedGoPassSecretsManagerHookServer() {
}

// UnsafeGoPassSecretsManagerHookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoPassSecretsManagerHookServer will
// result in compilation errors.
type UnsafeGoPassSecretsManagerHookServer interface {
	mustEmbedUnimplementedGoPassSecretsManagerHookServer()
}

func RegisterGoPassSecretsManagerHookServer(s grpc.ServiceRegistrar, srv GoPassSecretsManagerHookServer) {
	s.RegisterService(&GoPassSecretsManagerHook_ServiceDesc, srv)
}

func _GoPassSecretsManagerHook_PostHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PostHookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoPassSecretsManagerHookServer).PostHook(m, &goPassSecretsManagerHookPostHookServer{stream})
}

type GoPassSecretsManagerHook_PostHookServer interface {
	Send(*PostHookResponse) error
	grpc.ServerStream
}

type goPassSecretsManagerHookPostHookServer struct {
	grpc.ServerStream
}

func (x *goPassSecretsManagerHookPostHookServer) Send(m *PostHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GoPassSecretsManagerHook_ServiceDesc is the grpc.ServiceDesc for GoPassSecretsManagerHook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoPassSecretsManagerHook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsm.GoPassSecretsManagerHook",
	HandlerType: (*GoPassSecretsManagerHookServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostHook",
			Handler:       _GoPassSecretsManagerHook_PostHook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gsm.proto",
}
