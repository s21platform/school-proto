// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: school.proto

package school_proto

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
	SchoolService_Login_FullMethodName = "/SchoolService/Login"
)

// SchoolServiceClient is the client API for SchoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchoolServiceClient interface {
	// Метод для получения токена sberclass, которым будут подписаны запросы к платформе
	Login(ctx context.Context, in *SchoolLoginRequest, opts ...grpc.CallOption) (*SchoolLoginResponse, error)
}

type schoolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchoolServiceClient(cc grpc.ClientConnInterface) SchoolServiceClient {
	return &schoolServiceClient{cc}
}

func (c *schoolServiceClient) Login(ctx context.Context, in *SchoolLoginRequest, opts ...grpc.CallOption) (*SchoolLoginResponse, error) {
	out := new(SchoolLoginResponse)
	err := c.cc.Invoke(ctx, SchoolService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchoolServiceServer is the server API for SchoolService service.
// All implementations must embed UnimplementedSchoolServiceServer
// for forward compatibility
type SchoolServiceServer interface {
	// Метод для получения токена sberclass, которым будут подписаны запросы к платформе
	Login(context.Context, *SchoolLoginRequest) (*SchoolLoginResponse, error)
	mustEmbedUnimplementedSchoolServiceServer()
}

// UnimplementedSchoolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchoolServiceServer struct {
}

func (UnimplementedSchoolServiceServer) Login(context.Context, *SchoolLoginRequest) (*SchoolLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSchoolServiceServer) mustEmbedUnimplementedSchoolServiceServer() {}

// UnsafeSchoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchoolServiceServer will
// result in compilation errors.
type UnsafeSchoolServiceServer interface {
	mustEmbedUnimplementedSchoolServiceServer()
}

func RegisterSchoolServiceServer(s grpc.ServiceRegistrar, srv SchoolServiceServer) {
	s.RegisterService(&SchoolService_ServiceDesc, srv)
}

func _SchoolService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchoolLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchoolService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolServiceServer).Login(ctx, req.(*SchoolLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchoolService_ServiceDesc is the grpc.ServiceDesc for SchoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SchoolService",
	HandlerType: (*SchoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SchoolService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "school.proto",
}
