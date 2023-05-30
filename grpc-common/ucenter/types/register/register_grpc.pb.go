// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: register.proto

package register

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
	Register_RegisterByPhone_FullMethodName = "/register.Register/registerByPhone"
	Register_SendCode_FullMethodName        = "/register.Register/sendCode"
)

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterClient interface {
	RegisterByPhone(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRes, error)
	SendCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*NoRes, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) RegisterByPhone(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRes, error) {
	out := new(RegRes)
	err := c.cc.Invoke(ctx, Register_RegisterByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) SendCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*NoRes, error) {
	out := new(NoRes)
	err := c.cc.Invoke(ctx, Register_SendCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
// All implementations must embed UnimplementedRegisterServer
// for forward compatibility
type RegisterServer interface {
	RegisterByPhone(context.Context, *RegReq) (*RegRes, error)
	SendCode(context.Context, *CodeReq) (*NoRes, error)
	mustEmbedUnimplementedRegisterServer()
}

// UnimplementedRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (UnimplementedRegisterServer) RegisterByPhone(context.Context, *RegReq) (*RegRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterByPhone not implemented")
}
func (UnimplementedRegisterServer) SendCode(context.Context, *CodeReq) (*NoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedRegisterServer) mustEmbedUnimplementedRegisterServer() {}

// UnsafeRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServer will
// result in compilation errors.
type UnsafeRegisterServer interface {
	mustEmbedUnimplementedRegisterServer()
}

func RegisterRegisterServer(s grpc.ServiceRegistrar, srv RegisterServer) {
	s.RegisterService(&Register_ServiceDesc, srv)
}

func _Register_RegisterByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).RegisterByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_RegisterByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).RegisterByPhone(ctx, req.(*RegReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_SendCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).SendCode(ctx, req.(*CodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Register_ServiceDesc is the grpc.ServiceDesc for Register service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Register_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "register.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerByPhone",
			Handler:    _Register_RegisterByPhone_Handler,
		},
		{
			MethodName: "sendCode",
			Handler:    _Register_SendCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
