// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: messages.proto

package messages

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
	MessageDispatchService_ValidateAccessToken_FullMethodName             = "/MessageDispatchService/ValidateAccessToken"
	MessageDispatchService_GetAccessToken_FullMethodName                  = "/MessageDispatchService/GetAccessToken"
	MessageDispatchService_SendActions_FullMethodName                     = "/MessageDispatchService/SendActions"
	MessageDispatchService_SendClusterGatewayResource_FullMethodName      = "/MessageDispatchService/SendClusterGatewayResource"
	MessageDispatchService_ReceiveError_FullMethodName                    = "/MessageDispatchService/ReceiveError"
	MessageDispatchService_ReceiveConsoleResourceUpdate_FullMethodName    = "/MessageDispatchService/ReceiveConsoleResourceUpdate"
	MessageDispatchService_ReceiveIotConsoleResourceUpdate_FullMethodName = "/MessageDispatchService/ReceiveIotConsoleResourceUpdate"
	MessageDispatchService_ReceiveInfraResourceUpdate_FullMethodName      = "/MessageDispatchService/ReceiveInfraResourceUpdate"
	MessageDispatchService_ReceiveContainerRegistryUpdate_FullMethodName  = "/MessageDispatchService/ReceiveContainerRegistryUpdate"
	MessageDispatchService_Ping_FullMethodName                            = "/MessageDispatchService/Ping"
)

// MessageDispatchServiceClient is the client API for MessageDispatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageDispatchServiceClient interface {
	ValidateAccessToken(ctx context.Context, in *ValidateAccessTokenIn, opts ...grpc.CallOption) (*ValidateAccessTokenOut, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenIn, opts ...grpc.CallOption) (*GetAccessTokenOut, error)
	SendActions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MessageDispatchService_SendActionsClient, error)
	SendClusterGatewayResource(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GatewayResource, error)
	ReceiveError(ctx context.Context, in *ErrorData, opts ...grpc.CallOption) (*Empty, error)
	ReceiveConsoleResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error)
	ReceiveIotConsoleResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error)
	ReceiveInfraResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error)
	ReceiveContainerRegistryUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingOutput, error)
}

type messageDispatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageDispatchServiceClient(cc grpc.ClientConnInterface) MessageDispatchServiceClient {
	return &messageDispatchServiceClient{cc}
}

func (c *messageDispatchServiceClient) ValidateAccessToken(ctx context.Context, in *ValidateAccessTokenIn, opts ...grpc.CallOption) (*ValidateAccessTokenOut, error) {
	out := new(ValidateAccessTokenOut)
	err := c.cc.Invoke(ctx, MessageDispatchService_ValidateAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) GetAccessToken(ctx context.Context, in *GetAccessTokenIn, opts ...grpc.CallOption) (*GetAccessTokenOut, error) {
	out := new(GetAccessTokenOut)
	err := c.cc.Invoke(ctx, MessageDispatchService_GetAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) SendActions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MessageDispatchService_SendActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[0], MessageDispatchService_SendActions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceSendActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageDispatchService_SendActionsClient interface {
	Recv() (*Action, error)
	grpc.ClientStream
}

type messageDispatchServiceSendActionsClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceSendActionsClient) Recv() (*Action, error) {
	m := new(Action)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageDispatchServiceClient) SendClusterGatewayResource(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GatewayResource, error) {
	out := new(GatewayResource)
	err := c.cc.Invoke(ctx, MessageDispatchService_SendClusterGatewayResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) ReceiveError(ctx context.Context, in *ErrorData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MessageDispatchService_ReceiveError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) ReceiveConsoleResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MessageDispatchService_ReceiveConsoleResourceUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) ReceiveIotConsoleResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MessageDispatchService_ReceiveIotConsoleResourceUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) ReceiveInfraResourceUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MessageDispatchService_ReceiveInfraResourceUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) ReceiveContainerRegistryUpdate(ctx context.Context, in *ResourceUpdate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MessageDispatchService_ReceiveContainerRegistryUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingOutput, error) {
	out := new(PingOutput)
	err := c.cc.Invoke(ctx, MessageDispatchService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageDispatchServiceServer is the server API for MessageDispatchService service.
// All implementations must embed UnimplementedMessageDispatchServiceServer
// for forward compatibility
type MessageDispatchServiceServer interface {
	ValidateAccessToken(context.Context, *ValidateAccessTokenIn) (*ValidateAccessTokenOut, error)
	GetAccessToken(context.Context, *GetAccessTokenIn) (*GetAccessTokenOut, error)
	SendActions(*Empty, MessageDispatchService_SendActionsServer) error
	SendClusterGatewayResource(context.Context, *Empty) (*GatewayResource, error)
	ReceiveError(context.Context, *ErrorData) (*Empty, error)
	ReceiveConsoleResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error)
	ReceiveIotConsoleResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error)
	ReceiveInfraResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error)
	ReceiveContainerRegistryUpdate(context.Context, *ResourceUpdate) (*Empty, error)
	Ping(context.Context, *Empty) (*PingOutput, error)
	mustEmbedUnimplementedMessageDispatchServiceServer()
}

// UnimplementedMessageDispatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageDispatchServiceServer struct {
}

func (UnimplementedMessageDispatchServiceServer) ValidateAccessToken(context.Context, *ValidateAccessTokenIn) (*ValidateAccessTokenOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAccessToken not implemented")
}
func (UnimplementedMessageDispatchServiceServer) GetAccessToken(context.Context, *GetAccessTokenIn) (*GetAccessTokenOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedMessageDispatchServiceServer) SendActions(*Empty, MessageDispatchService_SendActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendActions not implemented")
}
func (UnimplementedMessageDispatchServiceServer) SendClusterGatewayResource(context.Context, *Empty) (*GatewayResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClusterGatewayResource not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveError(context.Context, *ErrorData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveError not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveConsoleResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveConsoleResourceUpdate not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveIotConsoleResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveIotConsoleResourceUpdate not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveInfraResourceUpdate(context.Context, *ResourceUpdate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveInfraResourceUpdate not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveContainerRegistryUpdate(context.Context, *ResourceUpdate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveContainerRegistryUpdate not implemented")
}
func (UnimplementedMessageDispatchServiceServer) Ping(context.Context, *Empty) (*PingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMessageDispatchServiceServer) mustEmbedUnimplementedMessageDispatchServiceServer() {
}

// UnsafeMessageDispatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageDispatchServiceServer will
// result in compilation errors.
type UnsafeMessageDispatchServiceServer interface {
	mustEmbedUnimplementedMessageDispatchServiceServer()
}

func RegisterMessageDispatchServiceServer(s grpc.ServiceRegistrar, srv MessageDispatchServiceServer) {
	s.RegisterService(&MessageDispatchService_ServiceDesc, srv)
}

func _MessageDispatchService_ValidateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAccessTokenIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ValidateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ValidateAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ValidateAccessToken(ctx, req.(*ValidateAccessTokenIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_GetAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).GetAccessToken(ctx, req.(*GetAccessTokenIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_SendActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageDispatchServiceServer).SendActions(m, &messageDispatchServiceSendActionsServer{stream})
}

type MessageDispatchService_SendActionsServer interface {
	Send(*Action) error
	grpc.ServerStream
}

type messageDispatchServiceSendActionsServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceSendActionsServer) Send(m *Action) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageDispatchService_SendClusterGatewayResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).SendClusterGatewayResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_SendClusterGatewayResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).SendClusterGatewayResource(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_ReceiveError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ReceiveError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ReceiveError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ReceiveError(ctx, req.(*ErrorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_ReceiveConsoleResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ReceiveConsoleResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ReceiveConsoleResourceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ReceiveConsoleResourceUpdate(ctx, req.(*ResourceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_ReceiveIotConsoleResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ReceiveIotConsoleResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ReceiveIotConsoleResourceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ReceiveIotConsoleResourceUpdate(ctx, req.(*ResourceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_ReceiveInfraResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ReceiveInfraResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ReceiveInfraResourceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ReceiveInfraResourceUpdate(ctx, req.(*ResourceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_ReceiveContainerRegistryUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).ReceiveContainerRegistryUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_ReceiveContainerRegistryUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).ReceiveContainerRegistryUpdate(ctx, req.(*ResourceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageDispatchService_ServiceDesc is the grpc.ServiceDesc for MessageDispatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageDispatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageDispatchService",
	HandlerType: (*MessageDispatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateAccessToken",
			Handler:    _MessageDispatchService_ValidateAccessToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _MessageDispatchService_GetAccessToken_Handler,
		},
		{
			MethodName: "SendClusterGatewayResource",
			Handler:    _MessageDispatchService_SendClusterGatewayResource_Handler,
		},
		{
			MethodName: "ReceiveError",
			Handler:    _MessageDispatchService_ReceiveError_Handler,
		},
		{
			MethodName: "ReceiveConsoleResourceUpdate",
			Handler:    _MessageDispatchService_ReceiveConsoleResourceUpdate_Handler,
		},
		{
			MethodName: "ReceiveIotConsoleResourceUpdate",
			Handler:    _MessageDispatchService_ReceiveIotConsoleResourceUpdate_Handler,
		},
		{
			MethodName: "ReceiveInfraResourceUpdate",
			Handler:    _MessageDispatchService_ReceiveInfraResourceUpdate_Handler,
		},
		{
			MethodName: "ReceiveContainerRegistryUpdate",
			Handler:    _MessageDispatchService_ReceiveContainerRegistryUpdate_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MessageDispatchService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendActions",
			Handler:       _MessageDispatchService_SendActions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messages.proto",
}
