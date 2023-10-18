// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/resiliency/resiliency.proto

package resiliency

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
	ResiliencyService_UnaryResiliency_FullMethodName           = "/resiliency.ResiliencyService/UnaryResiliency"
	ResiliencyService_ServerStreamingResiliency_FullMethodName = "/resiliency.ResiliencyService/ServerStreamingResiliency"
	ResiliencyService_ClientStreamingResiliency_FullMethodName = "/resiliency.ResiliencyService/ClientStreamingResiliency"
	ResiliencyService_BiDirectionalResiliency_FullMethodName   = "/resiliency.ResiliencyService/BiDirectionalResiliency"
)

// ResiliencyServiceClient is the client API for ResiliencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResiliencyServiceClient interface {
	UnaryResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error)
	ServerStreamingResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (ResiliencyService_ServerStreamingResiliencyClient, error)
	ClientStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (ResiliencyService_ClientStreamingResiliencyClient, error)
	BiDirectionalResiliency(ctx context.Context, opts ...grpc.CallOption) (ResiliencyService_BiDirectionalResiliencyClient, error)
}

type resiliencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResiliencyServiceClient(cc grpc.ClientConnInterface) ResiliencyServiceClient {
	return &resiliencyServiceClient{cc}
}

func (c *resiliencyServiceClient) UnaryResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error) {
	out := new(ResiliencyResponse)
	err := c.cc.Invoke(ctx, ResiliencyService_UnaryResiliency_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resiliencyServiceClient) ServerStreamingResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (ResiliencyService_ServerStreamingResiliencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[0], ResiliencyService_ServerStreamingResiliency_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &resiliencyServiceServerStreamingResiliencyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResiliencyService_ServerStreamingResiliencyClient interface {
	Recv() (*ResiliencyResponse, error)
	grpc.ClientStream
}

type resiliencyServiceServerStreamingResiliencyClient struct {
	grpc.ClientStream
}

func (x *resiliencyServiceServerStreamingResiliencyClient) Recv() (*ResiliencyResponse, error) {
	m := new(ResiliencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resiliencyServiceClient) ClientStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (ResiliencyService_ClientStreamingResiliencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[1], ResiliencyService_ClientStreamingResiliency_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &resiliencyServiceClientStreamingResiliencyClient{stream}
	return x, nil
}

type ResiliencyService_ClientStreamingResiliencyClient interface {
	Send(*ResiliencyRequest) error
	CloseAndRecv() (*ResiliencyResponse, error)
	grpc.ClientStream
}

type resiliencyServiceClientStreamingResiliencyClient struct {
	grpc.ClientStream
}

func (x *resiliencyServiceClientStreamingResiliencyClient) Send(m *ResiliencyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resiliencyServiceClientStreamingResiliencyClient) CloseAndRecv() (*ResiliencyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResiliencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resiliencyServiceClient) BiDirectionalResiliency(ctx context.Context, opts ...grpc.CallOption) (ResiliencyService_BiDirectionalResiliencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[2], ResiliencyService_BiDirectionalResiliency_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &resiliencyServiceBiDirectionalResiliencyClient{stream}
	return x, nil
}

type ResiliencyService_BiDirectionalResiliencyClient interface {
	Send(*ResiliencyRequest) error
	Recv() (*ResiliencyResponse, error)
	grpc.ClientStream
}

type resiliencyServiceBiDirectionalResiliencyClient struct {
	grpc.ClientStream
}

func (x *resiliencyServiceBiDirectionalResiliencyClient) Send(m *ResiliencyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resiliencyServiceBiDirectionalResiliencyClient) Recv() (*ResiliencyResponse, error) {
	m := new(ResiliencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResiliencyServiceServer is the server API for ResiliencyService service.
// All implementations must embed UnimplementedResiliencyServiceServer
// for forward compatibility
type ResiliencyServiceServer interface {
	UnaryResiliency(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error)
	ServerStreamingResiliency(*ResiliencyRequest, ResiliencyService_ServerStreamingResiliencyServer) error
	ClientStreamingResiliency(ResiliencyService_ClientStreamingResiliencyServer) error
	BiDirectionalResiliency(ResiliencyService_BiDirectionalResiliencyServer) error
	mustEmbedUnimplementedResiliencyServiceServer()
}

// UnimplementedResiliencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResiliencyServiceServer struct {
}

func (UnimplementedResiliencyServiceServer) UnaryResiliency(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) ServerStreamingResiliency(*ResiliencyRequest, ResiliencyService_ServerStreamingResiliencyServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) ClientStreamingResiliency(ResiliencyService_ClientStreamingResiliencyServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) BiDirectionalResiliency(ResiliencyService_BiDirectionalResiliencyServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) mustEmbedUnimplementedResiliencyServiceServer() {}

// UnsafeResiliencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResiliencyServiceServer will
// result in compilation errors.
type UnsafeResiliencyServiceServer interface {
	mustEmbedUnimplementedResiliencyServiceServer()
}

func RegisterResiliencyServiceServer(s grpc.ServiceRegistrar, srv ResiliencyServiceServer) {
	s.RegisterService(&ResiliencyService_ServiceDesc, srv)
}

func _ResiliencyService_UnaryResiliency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResiliencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResiliencyServiceServer).UnaryResiliency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResiliencyService_UnaryResiliency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResiliencyServiceServer).UnaryResiliency(ctx, req.(*ResiliencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResiliencyService_ServerStreamingResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResiliencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResiliencyServiceServer).ServerStreamingResiliency(m, &resiliencyServiceServerStreamingResiliencyServer{stream})
}

type ResiliencyService_ServerStreamingResiliencyServer interface {
	Send(*ResiliencyResponse) error
	grpc.ServerStream
}

type resiliencyServiceServerStreamingResiliencyServer struct {
	grpc.ServerStream
}

func (x *resiliencyServiceServerStreamingResiliencyServer) Send(m *ResiliencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ResiliencyService_ClientStreamingResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyServiceServer).ClientStreamingResiliency(&resiliencyServiceClientStreamingResiliencyServer{stream})
}

type ResiliencyService_ClientStreamingResiliencyServer interface {
	SendAndClose(*ResiliencyResponse) error
	Recv() (*ResiliencyRequest, error)
	grpc.ServerStream
}

type resiliencyServiceClientStreamingResiliencyServer struct {
	grpc.ServerStream
}

func (x *resiliencyServiceClientStreamingResiliencyServer) SendAndClose(m *ResiliencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resiliencyServiceClientStreamingResiliencyServer) Recv() (*ResiliencyRequest, error) {
	m := new(ResiliencyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ResiliencyService_BiDirectionalResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyServiceServer).BiDirectionalResiliency(&resiliencyServiceBiDirectionalResiliencyServer{stream})
}

type ResiliencyService_BiDirectionalResiliencyServer interface {
	Send(*ResiliencyResponse) error
	Recv() (*ResiliencyRequest, error)
	grpc.ServerStream
}

type resiliencyServiceBiDirectionalResiliencyServer struct {
	grpc.ServerStream
}

func (x *resiliencyServiceBiDirectionalResiliencyServer) Send(m *ResiliencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resiliencyServiceBiDirectionalResiliencyServer) Recv() (*ResiliencyRequest, error) {
	m := new(ResiliencyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResiliencyService_ServiceDesc is the grpc.ServiceDesc for ResiliencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResiliencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resiliency.ResiliencyService",
	HandlerType: (*ResiliencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryResiliency",
			Handler:    _ResiliencyService_UnaryResiliency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingResiliency",
			Handler:       _ResiliencyService_ServerStreamingResiliency_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingResiliency",
			Handler:       _ResiliencyService_ClientStreamingResiliency_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDirectionalResiliency",
			Handler:       _ResiliencyService_BiDirectionalResiliency_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/resiliency/resiliency.proto",
}
