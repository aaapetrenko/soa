// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mafia

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

// ReverseClient is the client API for Reverse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReverseClient interface {
	Name(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingResponse, error)
	PingMembers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingResponse, error)
	PingResultDay(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResultVoteDay, error)
	Exit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Vote(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	NightVote(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ListPlayers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ChangeTime(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type reverseClient struct {
	cc grpc.ClientConnInterface
}

func NewReverseClient(cc grpc.ClientConnInterface) ReverseClient {
	return &reverseClient{cc}
}

func (c *reverseClient) Name(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) PingMembers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/PingMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) PingResultDay(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResultVoteDay, error) {
	out := new(ResultVoteDay)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/PingResultDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) Exit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) Vote(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) NightVote(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/NightVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) ListPlayers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/ListPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseClient) ChangeTime(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mafia.Reverse/ChangeTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseServer is the server API for Reverse service.
// All implementations must embed UnimplementedReverseServer
// for forward compatibility
type ReverseServer interface {
	Name(context.Context, *Request) (*Response, error)
	Ping(context.Context, *Request) (*PingResponse, error)
	PingMembers(context.Context, *Request) (*PingResponse, error)
	PingResultDay(context.Context, *Request) (*ResultVoteDay, error)
	Exit(context.Context, *Request) (*Response, error)
	Vote(context.Context, *Request) (*Response, error)
	NightVote(context.Context, *Request) (*Response, error)
	ListPlayers(context.Context, *Request) (*Response, error)
	ChangeTime(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedReverseServer()
}

// UnimplementedReverseServer must be embedded to have forward compatible implementations.
type UnimplementedReverseServer struct {
}

func (UnimplementedReverseServer) Name(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedReverseServer) Ping(context.Context, *Request) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedReverseServer) PingMembers(context.Context, *Request) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingMembers not implemented")
}
func (UnimplementedReverseServer) PingResultDay(context.Context, *Request) (*ResultVoteDay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingResultDay not implemented")
}
func (UnimplementedReverseServer) Exit(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedReverseServer) Vote(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedReverseServer) NightVote(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NightVote not implemented")
}
func (UnimplementedReverseServer) ListPlayers(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlayers not implemented")
}
func (UnimplementedReverseServer) ChangeTime(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTime not implemented")
}
func (UnimplementedReverseServer) mustEmbedUnimplementedReverseServer() {}

// UnsafeReverseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReverseServer will
// result in compilation errors.
type UnsafeReverseServer interface {
	mustEmbedUnimplementedReverseServer()
}

func RegisterReverseServer(s grpc.ServiceRegistrar, srv ReverseServer) {
	s.RegisterService(&Reverse_ServiceDesc, srv)
}

func _Reverse_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).Name(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_PingMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).PingMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/PingMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).PingMembers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_PingResultDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).PingResultDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/PingResultDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).PingResultDay(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).Exit(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).Vote(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_NightVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).NightVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/NightVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).NightVote(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_ListPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).ListPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/ListPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).ListPlayers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reverse_ChangeTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).ChangeTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Reverse/ChangeTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).ChangeTime(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Reverse_ServiceDesc is the grpc.ServiceDesc for Reverse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reverse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mafia.Reverse",
	HandlerType: (*ReverseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _Reverse_Name_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Reverse_Ping_Handler,
		},
		{
			MethodName: "PingMembers",
			Handler:    _Reverse_PingMembers_Handler,
		},
		{
			MethodName: "PingResultDay",
			Handler:    _Reverse_PingResultDay_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _Reverse_Exit_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Reverse_Vote_Handler,
		},
		{
			MethodName: "NightVote",
			Handler:    _Reverse_NightVote_Handler,
		},
		{
			MethodName: "ListPlayers",
			Handler:    _Reverse_ListPlayers_Handler,
		},
		{
			MethodName: "ChangeTime",
			Handler:    _Reverse_ChangeTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mafia/mafia.proto",
}
