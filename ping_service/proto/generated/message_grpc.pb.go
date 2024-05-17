// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: message.proto

package generated

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
	Pong_Pong_FullMethodName   = "/message.Pong/Pong"
	Pong_Health_FullMethodName = "/message.Pong/Health"
)

// PongClient is the client API for Pong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PongClient interface {
	Pong(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Health(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type pongClient struct {
	cc grpc.ClientConnInterface
}

func NewPongClient(cc grpc.ClientConnInterface) PongClient {
	return &pongClient{cc}
}

func (c *pongClient) Pong(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Pong_Pong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pongClient) Health(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Pong_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PongServer is the server API for Pong service.
// All implementations must embed UnimplementedPongServer
// for forward compatibility
type PongServer interface {
	Pong(context.Context, *Request) (*Response, error)
	Health(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPongServer()
}

// UnimplementedPongServer must be embedded to have forward compatible implementations.
type UnimplementedPongServer struct {
}

func (UnimplementedPongServer) Pong(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedPongServer) Health(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedPongServer) mustEmbedUnimplementedPongServer() {}

// UnsafePongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PongServer will
// result in compilation errors.
type UnsafePongServer interface {
	mustEmbedUnimplementedPongServer()
}

func RegisterPongServer(s grpc.ServiceRegistrar, srv PongServer) {
	s.RegisterService(&Pong_ServiceDesc, srv)
}

func _Pong_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pong_Pong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongServer).Pong(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pong_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pong_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongServer).Health(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Pong_ServiceDesc is the grpc.ServiceDesc for Pong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Pong",
	HandlerType: (*PongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _Pong_Pong_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Pong_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
