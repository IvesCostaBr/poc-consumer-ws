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
	ExchangeSteam_Heathcheck_FullMethodName     = "/exchange.ExchangeSteam/Heathcheck"
	ExchangeSteam_RequestAction_FullMethodName  = "/exchange.ExchangeSteam/RequestAction"
	ExchangeSteam_SubscribeEvent_FullMethodName = "/exchange.ExchangeSteam/SubscribeEvent"
)

// ExchangeSteamClient is the client API for ExchangeSteam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeSteamClient interface {
	Heathcheck(ctx context.Context, in *HeathCheck, opts ...grpc.CallOption) (*HeathCheck, error)
	RequestAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataReceiver, error)
	SubscribeEvent(ctx context.Context, in *SubscribeEventRequest, opts ...grpc.CallOption) (ExchangeSteam_SubscribeEventClient, error)
}

type exchangeSteamClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeSteamClient(cc grpc.ClientConnInterface) ExchangeSteamClient {
	return &exchangeSteamClient{cc}
}

func (c *exchangeSteamClient) Heathcheck(ctx context.Context, in *HeathCheck, opts ...grpc.CallOption) (*HeathCheck, error) {
	out := new(HeathCheck)
	err := c.cc.Invoke(ctx, ExchangeSteam_Heathcheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeSteamClient) RequestAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataReceiver, error) {
	out := new(DataReceiver)
	err := c.cc.Invoke(ctx, ExchangeSteam_RequestAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeSteamClient) SubscribeEvent(ctx context.Context, in *SubscribeEventRequest, opts ...grpc.CallOption) (ExchangeSteam_SubscribeEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExchangeSteam_ServiceDesc.Streams[0], ExchangeSteam_SubscribeEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeSteamSubscribeEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExchangeSteam_SubscribeEventClient interface {
	Recv() (*DataReceiver, error)
	grpc.ClientStream
}

type exchangeSteamSubscribeEventClient struct {
	grpc.ClientStream
}

func (x *exchangeSteamSubscribeEventClient) Recv() (*DataReceiver, error) {
	m := new(DataReceiver)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeSteamServer is the server API for ExchangeSteam service.
// All implementations must embed UnimplementedExchangeSteamServer
// for forward compatibility
type ExchangeSteamServer interface {
	Heathcheck(context.Context, *HeathCheck) (*HeathCheck, error)
	RequestAction(context.Context, *ActionRequest) (*DataReceiver, error)
	SubscribeEvent(*SubscribeEventRequest, ExchangeSteam_SubscribeEventServer) error
	mustEmbedUnimplementedExchangeSteamServer()
}

// UnimplementedExchangeSteamServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeSteamServer struct {
}

func (UnimplementedExchangeSteamServer) Heathcheck(context.Context, *HeathCheck) (*HeathCheck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedExchangeSteamServer) RequestAction(context.Context, *ActionRequest) (*DataReceiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAction not implemented")
}
func (UnimplementedExchangeSteamServer) SubscribeEvent(*SubscribeEventRequest, ExchangeSteam_SubscribeEventServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeEvent not implemented")
}
func (UnimplementedExchangeSteamServer) mustEmbedUnimplementedExchangeSteamServer() {}

// UnsafeExchangeSteamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeSteamServer will
// result in compilation errors.
type UnsafeExchangeSteamServer interface {
	mustEmbedUnimplementedExchangeSteamServer()
}

func RegisterExchangeSteamServer(s grpc.ServiceRegistrar, srv ExchangeSteamServer) {
	s.RegisterService(&ExchangeSteam_ServiceDesc, srv)
}

func _ExchangeSteam_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeathCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeSteamServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeSteam_Heathcheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeSteamServer).Heathcheck(ctx, req.(*HeathCheck))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeSteam_RequestAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeSteamServer).RequestAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeSteam_RequestAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeSteamServer).RequestAction(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeSteam_SubscribeEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeSteamServer).SubscribeEvent(m, &exchangeSteamSubscribeEventServer{stream})
}

type ExchangeSteam_SubscribeEventServer interface {
	Send(*DataReceiver) error
	grpc.ServerStream
}

type exchangeSteamSubscribeEventServer struct {
	grpc.ServerStream
}

func (x *exchangeSteamSubscribeEventServer) Send(m *DataReceiver) error {
	return x.ServerStream.SendMsg(m)
}

// ExchangeSteam_ServiceDesc is the grpc.ServiceDesc for ExchangeSteam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeSteam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.ExchangeSteam",
	HandlerType: (*ExchangeSteamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _ExchangeSteam_Heathcheck_Handler,
		},
		{
			MethodName: "RequestAction",
			Handler:    _ExchangeSteam_RequestAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeEvent",
			Handler:       _ExchangeSteam_SubscribeEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/exchange.proto",
}
