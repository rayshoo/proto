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

// GoPassSecretManagerClient is the client API for GoPassSecretManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoPassSecretManagerClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (GoPassSecretManager_SyncClient, error)
}

type goPassSecretManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGoPassSecretManagerClient(cc grpc.ClientConnInterface) GoPassSecretManagerClient {
	return &goPassSecretManagerClient{cc}
}

func (c *goPassSecretManagerClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (GoPassSecretManager_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoPassSecretManager_ServiceDesc.Streams[0], "/gsm.GoPassSecretManager/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &goPassSecretManagerSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoPassSecretManager_SyncClient interface {
	Recv() (*SyncResponse, error)
	grpc.ClientStream
}

type goPassSecretManagerSyncClient struct {
	grpc.ClientStream
}

func (x *goPassSecretManagerSyncClient) Recv() (*SyncResponse, error) {
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoPassSecretManagerServer is the server API for GoPassSecretManager service.
// All implementations must embed UnimplementedGoPassSecretManagerServer
// for forward compatibility
type GoPassSecretManagerServer interface {
	Sync(*SyncRequest, GoPassSecretManager_SyncServer) error
	mustEmbedUnimplementedGoPassSecretManagerServer()
}

// UnimplementedGoPassSecretManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGoPassSecretManagerServer struct {
}

func (UnimplementedGoPassSecretManagerServer) Sync(*SyncRequest, GoPassSecretManager_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedGoPassSecretManagerServer) mustEmbedUnimplementedGoPassSecretManagerServer() {}

// UnsafeGoPassSecretManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoPassSecretManagerServer will
// result in compilation errors.
type UnsafeGoPassSecretManagerServer interface {
	mustEmbedUnimplementedGoPassSecretManagerServer()
}

func RegisterGoPassSecretManagerServer(s grpc.ServiceRegistrar, srv GoPassSecretManagerServer) {
	s.RegisterService(&GoPassSecretManager_ServiceDesc, srv)
}

func _GoPassSecretManager_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoPassSecretManagerServer).Sync(m, &goPassSecretManagerSyncServer{stream})
}

type GoPassSecretManager_SyncServer interface {
	Send(*SyncResponse) error
	grpc.ServerStream
}

type goPassSecretManagerSyncServer struct {
	grpc.ServerStream
}

func (x *goPassSecretManagerSyncServer) Send(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GoPassSecretManager_ServiceDesc is the grpc.ServiceDesc for GoPassSecretManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoPassSecretManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsm.GoPassSecretManager",
	HandlerType: (*GoPassSecretManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _GoPassSecretManager_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gsm.proto",
}
