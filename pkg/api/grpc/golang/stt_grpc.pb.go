// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: stt.proto

package stt

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SttService_NotifyInText_FullMethodName = "/SttService/NotifyInText"
)

// SttServiceClient is the client API for SttService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SttServiceClient interface {
	NotifyInText(ctx context.Context, in *NotifyInTextRq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type sttServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSttServiceClient(cc grpc.ClientConnInterface) SttServiceClient {
	return &sttServiceClient{cc}
}

func (c *sttServiceClient) NotifyInText(ctx context.Context, in *NotifyInTextRq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SttService_NotifyInText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SttServiceServer is the server API for SttService service.
// All implementations must embed UnimplementedSttServiceServer
// for forward compatibility
type SttServiceServer interface {
	NotifyInText(context.Context, *NotifyInTextRq) (*emptypb.Empty, error)
	mustEmbedUnimplementedSttServiceServer()
}

// UnimplementedSttServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSttServiceServer struct {
}

func (UnimplementedSttServiceServer) NotifyInText(context.Context, *NotifyInTextRq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyInText not implemented")
}
func (UnimplementedSttServiceServer) mustEmbedUnimplementedSttServiceServer() {}

// UnsafeSttServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SttServiceServer will
// result in compilation errors.
type UnsafeSttServiceServer interface {
	mustEmbedUnimplementedSttServiceServer()
}

func RegisterSttServiceServer(s grpc.ServiceRegistrar, srv SttServiceServer) {
	s.RegisterService(&SttService_ServiceDesc, srv)
}

func _SttService_NotifyInText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyInTextRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).NotifyInText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_NotifyInText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).NotifyInText(ctx, req.(*NotifyInTextRq))
	}
	return interceptor(ctx, in, info, handler)
}

// SttService_ServiceDesc is the grpc.ServiceDesc for SttService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SttService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SttService",
	HandlerType: (*SttServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyInText",
			Handler:    _SttService_NotifyInText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stt.proto",
}
