// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: console.proto

package console

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
	Console_ArchiveResourcesForCluster_FullMethodName = "/Console/ArchiveResourcesForCluster"
)

// ConsoleClient is the client API for Console service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsoleClient interface {
	ArchiveResourcesForCluster(ctx context.Context, in *ArchiveResourcesForClusterIn, opts ...grpc.CallOption) (*ArchiveResourcesForClusterOut, error)
}

type consoleClient struct {
	cc grpc.ClientConnInterface
}

func NewConsoleClient(cc grpc.ClientConnInterface) ConsoleClient {
	return &consoleClient{cc}
}

func (c *consoleClient) ArchiveResourcesForCluster(ctx context.Context, in *ArchiveResourcesForClusterIn, opts ...grpc.CallOption) (*ArchiveResourcesForClusterOut, error) {
	out := new(ArchiveResourcesForClusterOut)
	err := c.cc.Invoke(ctx, Console_ArchiveResourcesForCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsoleServer is the server API for Console service.
// All implementations must embed UnimplementedConsoleServer
// for forward compatibility
type ConsoleServer interface {
	ArchiveResourcesForCluster(context.Context, *ArchiveResourcesForClusterIn) (*ArchiveResourcesForClusterOut, error)
	mustEmbedUnimplementedConsoleServer()
}

// UnimplementedConsoleServer must be embedded to have forward compatible implementations.
type UnimplementedConsoleServer struct {
}

func (UnimplementedConsoleServer) ArchiveResourcesForCluster(context.Context, *ArchiveResourcesForClusterIn) (*ArchiveResourcesForClusterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveResourcesForCluster not implemented")
}
func (UnimplementedConsoleServer) mustEmbedUnimplementedConsoleServer() {}

// UnsafeConsoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsoleServer will
// result in compilation errors.
type UnsafeConsoleServer interface {
	mustEmbedUnimplementedConsoleServer()
}

func RegisterConsoleServer(s grpc.ServiceRegistrar, srv ConsoleServer) {
	s.RegisterService(&Console_ServiceDesc, srv)
}

func _Console_ArchiveResourcesForCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveResourcesForClusterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServer).ArchiveResourcesForCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Console_ArchiveResourcesForCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServer).ArchiveResourcesForCluster(ctx, req.(*ArchiveResourcesForClusterIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Console_ServiceDesc is the grpc.ServiceDesc for Console service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Console_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Console",
	HandlerType: (*ConsoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ArchiveResourcesForCluster",
			Handler:    _Console_ArchiveResourcesForCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "console.proto",
}
