// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: asset.proto

package asset

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

// AssetClient is the client API for Asset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetClient interface {
	FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error)
	FindWallet(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWalletList, error)
	ResetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AssetResp, error)
	FindTransaction(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberTransactionList, error)
	GetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AddressList, error)
}

type assetClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetClient(cc grpc.ClientConnInterface) AssetClient {
	return &assetClient{cc}
}

func (c *assetClient) FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error) {
	out := new(MemberWallet)
	err := c.cc.Invoke(ctx, "/asset.Asset/findWalletBySymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) FindWallet(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWalletList, error) {
	out := new(MemberWalletList)
	err := c.cc.Invoke(ctx, "/asset.Asset/findWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) ResetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AssetResp, error) {
	out := new(AssetResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/ResetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) FindTransaction(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberTransactionList, error) {
	out := new(MemberTransactionList)
	err := c.cc.Invoke(ctx, "/asset.Asset/FindTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AddressList, error) {
	out := new(AddressList)
	err := c.cc.Invoke(ctx, "/asset.Asset/getAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServer is the server API for Asset service.
// All implementations must embed UnimplementedAssetServer
// for forward compatibility
type AssetServer interface {
	FindWalletBySymbol(context.Context, *AssetReq) (*MemberWallet, error)
	FindWallet(context.Context, *AssetReq) (*MemberWalletList, error)
	ResetAddress(context.Context, *AssetReq) (*AssetResp, error)
	FindTransaction(context.Context, *AssetReq) (*MemberTransactionList, error)
	GetAddress(context.Context, *AssetReq) (*AddressList, error)
	mustEmbedUnimplementedAssetServer()
}

// UnimplementedAssetServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServer struct {
}

func (UnimplementedAssetServer) FindWalletBySymbol(context.Context, *AssetReq) (*MemberWallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWalletBySymbol not implemented")
}
func (UnimplementedAssetServer) FindWallet(context.Context, *AssetReq) (*MemberWalletList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWallet not implemented")
}
func (UnimplementedAssetServer) ResetAddress(context.Context, *AssetReq) (*AssetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAddress not implemented")
}
func (UnimplementedAssetServer) FindTransaction(context.Context, *AssetReq) (*MemberTransactionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTransaction not implemented")
}
func (UnimplementedAssetServer) GetAddress(context.Context, *AssetReq) (*AddressList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedAssetServer) mustEmbedUnimplementedAssetServer() {}

// UnsafeAssetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServer will
// result in compilation errors.
type UnsafeAssetServer interface {
	mustEmbedUnimplementedAssetServer()
}

func RegisterAssetServer(s grpc.ServiceRegistrar, srv AssetServer) {
	s.RegisterService(&Asset_ServiceDesc, srv)
}

func _Asset_FindWalletBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).FindWalletBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/findWalletBySymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).FindWalletBySymbol(ctx, req.(*AssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_FindWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).FindWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/findWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).FindWallet(ctx, req.(*AssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_ResetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).ResetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/ResetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).ResetAddress(ctx, req.(*AssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_FindTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).FindTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/FindTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).FindTransaction(ctx, req.(*AssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/getAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).GetAddress(ctx, req.(*AssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Asset_ServiceDesc is the grpc.ServiceDesc for Asset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Asset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asset.Asset",
	HandlerType: (*AssetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findWalletBySymbol",
			Handler:    _Asset_FindWalletBySymbol_Handler,
		},
		{
			MethodName: "findWallet",
			Handler:    _Asset_FindWallet_Handler,
		},
		{
			MethodName: "ResetAddress",
			Handler:    _Asset_ResetAddress_Handler,
		},
		{
			MethodName: "FindTransaction",
			Handler:    _Asset_FindTransaction_Handler,
		},
		{
			MethodName: "getAddress",
			Handler:    _Asset_GetAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "asset.proto",
}
