// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: http.proto

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HttpService_Echo_FullMethodName         = "/example.HttpService/Echo"
	HttpService_EchoProtobuf_FullMethodName = "/example.HttpService/EchoProtobuf"
	HttpService_TestGrpc_FullMethodName     = "/example.HttpService/TestGrpc"
)

// HttpServiceClient is the client API for HttpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpServiceClient interface {
	Echo(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error)
	EchoProtobuf(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error)
	TestGrpc(ctx context.Context, in *TestGrpcReq, opts ...grpc.CallOption) (*TestGrpcRsp, error)
}

type httpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpServiceClient(cc grpc.ClientConnInterface) HttpServiceClient {
	return &httpServiceClient{cc}
}

func (c *httpServiceClient) Echo(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HttpResponse)
	err := c.cc.Invoke(ctx, HttpService_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) EchoProtobuf(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HttpResponse)
	err := c.cc.Invoke(ctx, HttpService_EchoProtobuf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) TestGrpc(ctx context.Context, in *TestGrpcReq, opts ...grpc.CallOption) (*TestGrpcRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestGrpcRsp)
	err := c.cc.Invoke(ctx, HttpService_TestGrpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpServiceServer is the server API for HttpService service.
// All implementations must embed UnimplementedHttpServiceServer
// for forward compatibility.
type HttpServiceServer interface {
	Echo(context.Context, *HttpRequest) (*HttpResponse, error)
	EchoProtobuf(context.Context, *HttpRequest) (*HttpResponse, error)
	TestGrpc(context.Context, *TestGrpcReq) (*TestGrpcRsp, error)
	mustEmbedUnimplementedHttpServiceServer()
}

// UnimplementedHttpServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHttpServiceServer struct{}

func (UnimplementedHttpServiceServer) Echo(context.Context, *HttpRequest) (*HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedHttpServiceServer) EchoProtobuf(context.Context, *HttpRequest) (*HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoProtobuf not implemented")
}
func (UnimplementedHttpServiceServer) TestGrpc(context.Context, *TestGrpcReq) (*TestGrpcRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestGrpc not implemented")
}
func (UnimplementedHttpServiceServer) mustEmbedUnimplementedHttpServiceServer() {}
func (UnimplementedHttpServiceServer) testEmbeddedByValue()                     {}

// UnsafeHttpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpServiceServer will
// result in compilation errors.
type UnsafeHttpServiceServer interface {
	mustEmbedUnimplementedHttpServiceServer()
}

func RegisterHttpServiceServer(s grpc.ServiceRegistrar, srv HttpServiceServer) {
	// If the following call pancis, it indicates UnimplementedHttpServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HttpService_ServiceDesc, srv)
}

func _HttpService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HttpService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).Echo(ctx, req.(*HttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_EchoProtobuf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).EchoProtobuf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HttpService_EchoProtobuf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).EchoProtobuf(ctx, req.(*HttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_TestGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestGrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).TestGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HttpService_TestGrpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).TestGrpc(ctx, req.(*TestGrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpService_ServiceDesc is the grpc.ServiceDesc for HttpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.HttpService",
	HandlerType: (*HttpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _HttpService_Echo_Handler,
		},
		{
			MethodName: "EchoProtobuf",
			Handler:    _HttpService_EchoProtobuf_Handler,
		},
		{
			MethodName: "TestGrpc",
			Handler:    _HttpService_TestGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "http.proto",
}
