// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: openaipb/openai.proto

package openaipb

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

// ChatgptClient is the client API for Chatgpt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatgptClient interface {
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (Chatgpt_SendClient, error)
}

type chatgptClient struct {
	cc grpc.ClientConnInterface
}

func NewChatgptClient(cc grpc.ClientConnInterface) ChatgptClient {
	return &chatgptClient{cc}
}

func (c *chatgptClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (Chatgpt_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatgpt_ServiceDesc.Streams[0], "/openai.Chatgpt/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatgptSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chatgpt_SendClient interface {
	Recv() (*Answer, error)
	grpc.ClientStream
}

type chatgptSendClient struct {
	grpc.ClientStream
}

func (x *chatgptSendClient) Recv() (*Answer, error) {
	m := new(Answer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatgptServer is the server API for Chatgpt service.
// All implementations must embed UnimplementedChatgptServer
// for forward compatibility
type ChatgptServer interface {
	Send(*Message, Chatgpt_SendServer) error
	mustEmbedUnimplementedChatgptServer()
}

// UnimplementedChatgptServer must be embedded to have forward compatible implementations.
type UnimplementedChatgptServer struct {
}

func (UnimplementedChatgptServer) Send(*Message, Chatgpt_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedChatgptServer) mustEmbedUnimplementedChatgptServer() {}

// UnsafeChatgptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatgptServer will
// result in compilation errors.
type UnsafeChatgptServer interface {
	mustEmbedUnimplementedChatgptServer()
}

func RegisterChatgptServer(s grpc.ServiceRegistrar, srv ChatgptServer) {
	s.RegisterService(&Chatgpt_ServiceDesc, srv)
}

func _Chatgpt_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatgptServer).Send(m, &chatgptSendServer{stream})
}

type Chatgpt_SendServer interface {
	Send(*Answer) error
	grpc.ServerStream
}

type chatgptSendServer struct {
	grpc.ServerStream
}

func (x *chatgptSendServer) Send(m *Answer) error {
	return x.ServerStream.SendMsg(m)
}

// Chatgpt_ServiceDesc is the grpc.ServiceDesc for Chatgpt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chatgpt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openai.Chatgpt",
	HandlerType: (*ChatgptServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Chatgpt_Send_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "openaipb/openai.proto",
}
