// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: movieapi/movieapi.proto

package movieapi

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

// MovieInfoClient is the client API for MovieInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieInfoClient interface {
	// Sends a requeest for movie info
	GetMovieInfo(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error)
	SetMovieInfo(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error)
}

type movieInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieInfoClient(cc grpc.ClientConnInterface) MovieInfoClient {
	return &movieInfoClient{cc}
}

func (c *movieInfoClient) GetMovieInfo(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error) {
	out := new(MovieReply)
	err := c.cc.Invoke(ctx, "/movieapi.MovieInfo/GetMovieInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieInfoClient) SetMovieInfo(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error) {
	out := new(MovieReply)
	err := c.cc.Invoke(ctx, "/movieapi.MovieInfo/SetMovieInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieInfoServer is the server API for MovieInfo service.
// All implementations must embed UnimplementedMovieInfoServer
// for forward compatibility
type MovieInfoServer interface {
	// Sends a requeest for movie info
	GetMovieInfo(context.Context, *MovieRequest) (*MovieReply, error)
	SetMovieInfo(context.Context, *MovieRequest) (*MovieReply, error)
	mustEmbedUnimplementedMovieInfoServer()
}

// UnimplementedMovieInfoServer must be embedded to have forward compatible implementations.
type UnimplementedMovieInfoServer struct {
}

func (UnimplementedMovieInfoServer) GetMovieInfo(context.Context, *MovieRequest) (*MovieReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieInfo not implemented")
}
func (UnimplementedMovieInfoServer) SetMovieInfo(context.Context, *MovieRequest) (*MovieReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMovieInfo not implemented")
}
func (UnimplementedMovieInfoServer) mustEmbedUnimplementedMovieInfoServer() {}

// UnsafeMovieInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieInfoServer will
// result in compilation errors.
type UnsafeMovieInfoServer interface {
	mustEmbedUnimplementedMovieInfoServer()
}

func RegisterMovieInfoServer(s grpc.ServiceRegistrar, srv MovieInfoServer) {
	s.RegisterService(&MovieInfo_ServiceDesc, srv)
}

func _MovieInfo_GetMovieInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieInfoServer).GetMovieInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieapi.MovieInfo/GetMovieInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieInfoServer).GetMovieInfo(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieInfo_SetMovieInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieInfoServer).SetMovieInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieapi.MovieInfo/SetMovieInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieInfoServer).SetMovieInfo(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieInfo_ServiceDesc is the grpc.ServiceDesc for MovieInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movieapi.MovieInfo",
	HandlerType: (*MovieInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieInfo",
			Handler:    _MovieInfo_GetMovieInfo_Handler,
		},
		{
			MethodName: "SetMovieInfo",
			Handler:    _MovieInfo_SetMovieInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movieapi/movieapi.proto",
}
