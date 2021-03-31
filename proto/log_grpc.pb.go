// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SendLogsClient is the client API for SendLogs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendLogsClient interface {
	Send(ctx context.Context, in *Logs, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sendLogsClient struct {
	cc grpc.ClientConnInterface
}

func NewSendLogsClient(cc grpc.ClientConnInterface) SendLogsClient {
	return &sendLogsClient{cc}
}

func (c *sendLogsClient) Send(ctx context.Context, in *Logs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.SendLogs/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendLogsServer is the server API for SendLogs service.
// All implementations must embed UnimplementedSendLogsServer
// for forward compatibility
type SendLogsServer interface {
	Send(context.Context, *Logs) (*empty.Empty, error)
	mustEmbedUnimplementedSendLogsServer()
}

// UnimplementedSendLogsServer must be embedded to have forward compatible implementations.
type UnimplementedSendLogsServer struct {
}

func (UnimplementedSendLogsServer) Send(context.Context, *Logs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSendLogsServer) mustEmbedUnimplementedSendLogsServer() {}

// UnsafeSendLogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendLogsServer will
// result in compilation errors.
type UnsafeSendLogsServer interface {
	mustEmbedUnimplementedSendLogsServer()
}

func RegisterSendLogsServer(s grpc.ServiceRegistrar, srv SendLogsServer) {
	s.RegisterService(&SendLogs_ServiceDesc, srv)
}

func _SendLogs_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Logs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendLogsServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SendLogs/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendLogsServer).Send(ctx, req.(*Logs))
	}
	return interceptor(ctx, in, info, handler)
}

// SendLogs_ServiceDesc is the grpc.ServiceDesc for SendLogs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendLogs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendLogs",
	HandlerType: (*SendLogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SendLogs_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
