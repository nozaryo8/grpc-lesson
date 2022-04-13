// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/file.proto

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

// FilseServiseClient is the client API for FilseServise service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilseServiseClient interface {
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FilseServise_DownloadClient, error)
}

type filseServiseClient struct {
	cc grpc.ClientConnInterface
}

func NewFilseServiseClient(cc grpc.ClientConnInterface) FilseServiseClient {
	return &filseServiseClient{cc}
}

func (c *filseServiseClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, "/file.FilseServise/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filseServiseClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FilseServise_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FilseServise_ServiceDesc.Streams[0], "/file.FilseServise/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &filseServiseDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FilseServise_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type filseServiseDownloadClient struct {
	grpc.ClientStream
}

func (x *filseServiseDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilseServiseServer is the server API for FilseServise service.
// All implementations must embed UnimplementedFilseServiseServer
// for forward compatibility
type FilseServiseServer interface {
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	Download(*DownloadRequest, FilseServise_DownloadServer) error
	mustEmbedUnimplementedFilseServiseServer()
}

// UnimplementedFilseServiseServer must be embedded to have forward compatible implementations.
type UnimplementedFilseServiseServer struct {
}

func (UnimplementedFilseServiseServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFilseServiseServer) Download(*DownloadRequest, FilseServise_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFilseServiseServer) mustEmbedUnimplementedFilseServiseServer() {}

// UnsafeFilseServiseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilseServiseServer will
// result in compilation errors.
type UnsafeFilseServiseServer interface {
	mustEmbedUnimplementedFilseServiseServer()
}

func RegisterFilseServiseServer(s grpc.ServiceRegistrar, srv FilseServiseServer) {
	s.RegisterService(&FilseServise_ServiceDesc, srv)
}

func _FilseServise_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilseServiseServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FilseServise/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilseServiseServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilseServise_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilseServiseServer).Download(m, &filseServiseDownloadServer{stream})
}

type FilseServise_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type filseServiseDownloadServer struct {
	grpc.ServerStream
}

func (x *filseServiseDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FilseServise_ServiceDesc is the grpc.ServiceDesc for FilseServise service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilseServise_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FilseServise",
	HandlerType: (*FilseServiseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FilseServise_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _FilseServise_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}