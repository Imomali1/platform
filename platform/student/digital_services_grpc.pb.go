// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: platform/student/digital_services.proto

package student

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

// DigitalServicesServiceClient is the client API for DigitalServicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DigitalServicesServiceClient interface {
	SomeFunc(ctx context.Context, in *Request2, opts ...grpc.CallOption) (*Response2, error)
}

type digitalServicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDigitalServicesServiceClient(cc grpc.ClientConnInterface) DigitalServicesServiceClient {
	return &digitalServicesServiceClient{cc}
}

func (c *digitalServicesServiceClient) SomeFunc(ctx context.Context, in *Request2, opts ...grpc.CallOption) (*Response2, error) {
	out := new(Response2)
	err := c.cc.Invoke(ctx, "/DigitalServicesService/SomeFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigitalServicesServiceServer is the server API for DigitalServicesService service.
// All implementations must embed UnimplementedDigitalServicesServiceServer
// for forward compatibility
type DigitalServicesServiceServer interface {
	SomeFunc(context.Context, *Request2) (*Response2, error)
	mustEmbedUnimplementedDigitalServicesServiceServer()
}

// UnimplementedDigitalServicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDigitalServicesServiceServer struct {
}

func (UnimplementedDigitalServicesServiceServer) SomeFunc(context.Context, *Request2) (*Response2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SomeFunc not implemented")
}
func (UnimplementedDigitalServicesServiceServer) mustEmbedUnimplementedDigitalServicesServiceServer() {
}

// UnsafeDigitalServicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DigitalServicesServiceServer will
// result in compilation errors.
type UnsafeDigitalServicesServiceServer interface {
	mustEmbedUnimplementedDigitalServicesServiceServer()
}

func RegisterDigitalServicesServiceServer(s grpc.ServiceRegistrar, srv DigitalServicesServiceServer) {
	s.RegisterService(&DigitalServicesService_ServiceDesc, srv)
}

func _DigitalServicesService_SomeFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigitalServicesServiceServer).SomeFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DigitalServicesService/SomeFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigitalServicesServiceServer).SomeFunc(ctx, req.(*Request2))
	}
	return interceptor(ctx, in, info, handler)
}

// DigitalServicesService_ServiceDesc is the grpc.ServiceDesc for DigitalServicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DigitalServicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DigitalServicesService",
	HandlerType: (*DigitalServicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SomeFunc",
			Handler:    _DigitalServicesService_SomeFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform/student/digital_services.proto",
}
