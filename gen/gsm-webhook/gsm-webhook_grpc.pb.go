// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.0
// source: gsm-webhook.proto

package gsm_webhook

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

// GoPassSecretManagerWebhookClient is the client API for GoPassSecretManagerWebhook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoPassSecretManagerWebhookClient interface {
	Webhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (GoPassSecretManagerWebhook_WebhookClient, error)
}

type goPassSecretManagerWebhookClient struct {
	cc grpc.ClientConnInterface
}

func NewGoPassSecretManagerWebhookClient(cc grpc.ClientConnInterface) GoPassSecretManagerWebhookClient {
	return &goPassSecretManagerWebhookClient{cc}
}

func (c *goPassSecretManagerWebhookClient) Webhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (GoPassSecretManagerWebhook_WebhookClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoPassSecretManagerWebhook_ServiceDesc.Streams[0], "/gsm.GoPassSecretManagerWebhook/Webhook", opts...)
	if err != nil {
		return nil, err
	}
	x := &goPassSecretManagerWebhookWebhookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoPassSecretManagerWebhook_WebhookClient interface {
	Recv() (*WebhookResponse, error)
	grpc.ClientStream
}

type goPassSecretManagerWebhookWebhookClient struct {
	grpc.ClientStream
}

func (x *goPassSecretManagerWebhookWebhookClient) Recv() (*WebhookResponse, error) {
	m := new(WebhookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoPassSecretManagerWebhookServer is the server API for GoPassSecretManagerWebhook service.
// All implementations must embed UnimplementedGoPassSecretManagerWebhookServer
// for forward compatibility
type GoPassSecretManagerWebhookServer interface {
	Webhook(*WebhookRequest, GoPassSecretManagerWebhook_WebhookServer) error
	mustEmbedUnimplementedGoPassSecretManagerWebhookServer()
}

// UnimplementedGoPassSecretManagerWebhookServer must be embedded to have forward compatible implementations.
type UnimplementedGoPassSecretManagerWebhookServer struct {
}

func (UnimplementedGoPassSecretManagerWebhookServer) Webhook(*WebhookRequest, GoPassSecretManagerWebhook_WebhookServer) error {
	return status.Errorf(codes.Unimplemented, "method Webhook not implemented")
}
func (UnimplementedGoPassSecretManagerWebhookServer) mustEmbedUnimplementedGoPassSecretManagerWebhookServer() {
}

// UnsafeGoPassSecretManagerWebhookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoPassSecretManagerWebhookServer will
// result in compilation errors.
type UnsafeGoPassSecretManagerWebhookServer interface {
	mustEmbedUnimplementedGoPassSecretManagerWebhookServer()
}

func RegisterGoPassSecretManagerWebhookServer(s grpc.ServiceRegistrar, srv GoPassSecretManagerWebhookServer) {
	s.RegisterService(&GoPassSecretManagerWebhook_ServiceDesc, srv)
}

func _GoPassSecretManagerWebhook_Webhook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WebhookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoPassSecretManagerWebhookServer).Webhook(m, &goPassSecretManagerWebhookWebhookServer{stream})
}

type GoPassSecretManagerWebhook_WebhookServer interface {
	Send(*WebhookResponse) error
	grpc.ServerStream
}

type goPassSecretManagerWebhookWebhookServer struct {
	grpc.ServerStream
}

func (x *goPassSecretManagerWebhookWebhookServer) Send(m *WebhookResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GoPassSecretManagerWebhook_ServiceDesc is the grpc.ServiceDesc for GoPassSecretManagerWebhook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoPassSecretManagerWebhook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsm.GoPassSecretManagerWebhook",
	HandlerType: (*GoPassSecretManagerWebhookServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Webhook",
			Handler:       _GoPassSecretManagerWebhook_Webhook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gsm-webhook.proto",
}
