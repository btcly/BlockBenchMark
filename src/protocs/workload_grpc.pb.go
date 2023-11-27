// proto/hello.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: workload.proto

package proto

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
	WorkLoad_SendWorkLoad_FullMethodName = "/proto.WorkLoad/SendWorkLoad"
)

// WorkLoadClient is the client API for WorkLoad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkLoadClient interface {
	SendWorkLoad(ctx context.Context, in *WorkLoadRequest, opts ...grpc.CallOption) (*WorkLoadResponse, error)
}

type workLoadClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkLoadClient(cc grpc.ClientConnInterface) WorkLoadClient {
	return &workLoadClient{cc}
}

func (c *workLoadClient) SendWorkLoad(ctx context.Context, in *WorkLoadRequest, opts ...grpc.CallOption) (*WorkLoadResponse, error) {
	out := new(WorkLoadResponse)
	err := c.cc.Invoke(ctx, WorkLoad_SendWorkLoad_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkLoadServer is the server API for WorkLoad service.
// All implementations must embed UnimplementedWorkLoadServer
// for forward compatibility
type WorkLoadServer interface {
	SendWorkLoad(context.Context, *WorkLoadRequest) (*WorkLoadResponse, error)
	mustEmbedUnimplementedWorkLoadServer()
}

// UnimplementedWorkLoadServer must be embedded to have forward compatible implementations.
type UnimplementedWorkLoadServer struct {
}

func (UnimplementedWorkLoadServer) SendWorkLoad(context.Context, *WorkLoadRequest) (*WorkLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWorkLoad not implemented")
}
func (UnimplementedWorkLoadServer) mustEmbedUnimplementedWorkLoadServer() {}

// UnsafeWorkLoadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkLoadServer will
// result in compilation errors.
type UnsafeWorkLoadServer interface {
	mustEmbedUnimplementedWorkLoadServer()
}

func RegisterWorkLoadServer(s grpc.ServiceRegistrar, srv WorkLoadServer) {
	s.RegisterService(&WorkLoad_ServiceDesc, srv)
}

func _WorkLoad_SendWorkLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkLoadServer).SendWorkLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkLoad_SendWorkLoad_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkLoadServer).SendWorkLoad(ctx, req.(*WorkLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkLoad_ServiceDesc is the grpc.ServiceDesc for WorkLoad service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkLoad_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WorkLoad",
	HandlerType: (*WorkLoadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendWorkLoad",
			Handler:    _WorkLoad_SendWorkLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workload.proto",
}