// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: protos/exchange.proto

package pb

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
	StreamExample_RequestAction_FullMethodName  = "/exchange.StreamExample/RequestAction"
	StreamExample_SubscribeEvent_FullMethodName = "/exchange.StreamExample/SubscribeEvent"
)

// StreamExampleClient is the client API for StreamExample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamExampleClient interface {
	RequestAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataReceiver, error)
	SubscribeEvent(ctx context.Context, in *SubscribeEventRequest, opts ...grpc.CallOption) (StreamExample_SubscribeEventClient, error)
}

type streamExampleClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamExampleClient(cc grpc.ClientConnInterface) StreamExampleClient {
	return &streamExampleClient{cc}
}

func (c *streamExampleClient) RequestAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataReceiver, error) {
	out := new(DataReceiver)
	err := c.cc.Invoke(ctx, StreamExample_RequestAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamExampleClient) SubscribeEvent(ctx context.Context, in *SubscribeEventRequest, opts ...grpc.CallOption) (StreamExample_SubscribeEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamExample_ServiceDesc.Streams[0], StreamExample_SubscribeEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamExampleSubscribeEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamExample_SubscribeEventClient interface {
	Recv() (*DataReceiver, error)
	grpc.ClientStream
}

type streamExampleSubscribeEventClient struct {
	grpc.ClientStream
}

func (x *streamExampleSubscribeEventClient) Recv() (*DataReceiver, error) {
	m := new(DataReceiver)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamExampleServer is the server API for StreamExample service.
// All implementations must embed UnimplementedStreamExampleServer
// for forward compatibility
type StreamExampleServer interface {
	RequestAction(context.Context, *ActionRequest) (*DataReceiver, error)
	SubscribeEvent(*SubscribeEventRequest, StreamExample_SubscribeEventServer) error
	mustEmbedUnimplementedStreamExampleServer()
}

// UnimplementedStreamExampleServer must be embedded to have forward compatible implementations.
type UnimplementedStreamExampleServer struct {
}

func (UnimplementedStreamExampleServer) RequestAction(context.Context, *ActionRequest) (*DataReceiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAction not implemented")
}
func (UnimplementedStreamExampleServer) SubscribeEvent(*SubscribeEventRequest, StreamExample_SubscribeEventServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeEvent not implemented")
}
func (UnimplementedStreamExampleServer) mustEmbedUnimplementedStreamExampleServer() {}

// UnsafeStreamExampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamExampleServer will
// result in compilation errors.
type UnsafeStreamExampleServer interface {
	mustEmbedUnimplementedStreamExampleServer()
}

func RegisterStreamExampleServer(s grpc.ServiceRegistrar, srv StreamExampleServer) {
	s.RegisterService(&StreamExample_ServiceDesc, srv)
}

func _StreamExample_RequestAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamExampleServer).RequestAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamExample_RequestAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamExampleServer).RequestAction(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamExample_SubscribeEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamExampleServer).SubscribeEvent(m, &streamExampleSubscribeEventServer{stream})
}

type StreamExample_SubscribeEventServer interface {
	Send(*DataReceiver) error
	grpc.ServerStream
}

type streamExampleSubscribeEventServer struct {
	grpc.ServerStream
}

func (x *streamExampleSubscribeEventServer) Send(m *DataReceiver) error {
	return x.ServerStream.SendMsg(m)
}

// StreamExample_ServiceDesc is the grpc.ServiceDesc for StreamExample service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamExample_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.StreamExample",
	HandlerType: (*StreamExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAction",
			Handler:    _StreamExample_RequestAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeEvent",
			Handler:       _StreamExample_SubscribeEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/exchange.proto",
}