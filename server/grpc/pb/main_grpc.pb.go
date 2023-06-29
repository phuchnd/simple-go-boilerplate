// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: main.proto

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
	SimpleGoBoilerplateService_ListBooks_FullMethodName = "/simple_go_boilerplate.SimpleGoBoilerplateService/ListBooks"
)

// SimpleGoBoilerplateServiceClient is the client API for SimpleGoBoilerplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleGoBoilerplateServiceClient interface {
	ListBooks(ctx context.Context, in *ListBookRequest, opts ...grpc.CallOption) (*ListBookResponse, error)
}

type simpleGoBoilerplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleGoBoilerplateServiceClient(cc grpc.ClientConnInterface) SimpleGoBoilerplateServiceClient {
	return &simpleGoBoilerplateServiceClient{cc}
}

func (c *simpleGoBoilerplateServiceClient) ListBooks(ctx context.Context, in *ListBookRequest, opts ...grpc.CallOption) (*ListBookResponse, error) {
	out := new(ListBookResponse)
	err := c.cc.Invoke(ctx, SimpleGoBoilerplateService_ListBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleGoBoilerplateServiceServer is the server API for SimpleGoBoilerplateService service.
// All implementations must embed UnimplementedSimpleGoBoilerplateServiceServer
// for forward compatibility
type SimpleGoBoilerplateServiceServer interface {
	ListBooks(context.Context, *ListBookRequest) (*ListBookResponse, error)
	mustEmbedUnimplementedSimpleGoBoilerplateServiceServer()
}

// UnimplementedSimpleGoBoilerplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleGoBoilerplateServiceServer struct {
}

func (UnimplementedSimpleGoBoilerplateServiceServer) ListBooks(context.Context, *ListBookRequest) (*ListBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedSimpleGoBoilerplateServiceServer) mustEmbedUnimplementedSimpleGoBoilerplateServiceServer() {
}

// UnsafeSimpleGoBoilerplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleGoBoilerplateServiceServer will
// result in compilation errors.
type UnsafeSimpleGoBoilerplateServiceServer interface {
	mustEmbedUnimplementedSimpleGoBoilerplateServiceServer()
}

func RegisterSimpleGoBoilerplateServiceServer(s grpc.ServiceRegistrar, srv SimpleGoBoilerplateServiceServer) {
	s.RegisterService(&SimpleGoBoilerplateService_ServiceDesc, srv)
}

func _SimpleGoBoilerplateService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleGoBoilerplateServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleGoBoilerplateService_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleGoBoilerplateServiceServer).ListBooks(ctx, req.(*ListBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleGoBoilerplateService_ServiceDesc is the grpc.ServiceDesc for SimpleGoBoilerplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleGoBoilerplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simple_go_boilerplate.SimpleGoBoilerplateService",
	HandlerType: (*SimpleGoBoilerplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _SimpleGoBoilerplateService_ListBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}