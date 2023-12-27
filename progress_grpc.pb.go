// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: progress.proto

package progrock

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
	ProgressService_WriteUpdates_FullMethodName = "/progrock.ProgressService/WriteUpdates"
)

// ProgressServiceClient is the client API for ProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgressServiceClient interface {
	WriteUpdates(ctx context.Context, opts ...grpc.CallOption) (ProgressService_WriteUpdatesClient, error)
}

type progressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgressServiceClient(cc grpc.ClientConnInterface) ProgressServiceClient {
	return &progressServiceClient{cc}
}

func (c *progressServiceClient) WriteUpdates(ctx context.Context, opts ...grpc.CallOption) (ProgressService_WriteUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProgressService_ServiceDesc.Streams[0], ProgressService_WriteUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &progressServiceWriteUpdatesClient{stream}
	return x, nil
}

type ProgressService_WriteUpdatesClient interface {
	Send(*StatusUpdate) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type progressServiceWriteUpdatesClient struct {
	grpc.ClientStream
}

func (x *progressServiceWriteUpdatesClient) Send(m *StatusUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *progressServiceWriteUpdatesClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProgressServiceServer is the server API for ProgressService service.
// All implementations must embed UnimplementedProgressServiceServer
// for forward compatibility
type ProgressServiceServer interface {
	WriteUpdates(ProgressService_WriteUpdatesServer) error
	mustEmbedUnimplementedProgressServiceServer()
}

// UnimplementedProgressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgressServiceServer struct {
}

func (UnimplementedProgressServiceServer) WriteUpdates(ProgressService_WriteUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteUpdates not implemented")
}
func (UnimplementedProgressServiceServer) mustEmbedUnimplementedProgressServiceServer() {}

// UnsafeProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgressServiceServer will
// result in compilation errors.
type UnsafeProgressServiceServer interface {
	mustEmbedUnimplementedProgressServiceServer()
}

func RegisterProgressServiceServer(s grpc.ServiceRegistrar, srv ProgressServiceServer) {
	s.RegisterService(&ProgressService_ServiceDesc, srv)
}

func _ProgressService_WriteUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProgressServiceServer).WriteUpdates(&progressServiceWriteUpdatesServer{stream})
}

type ProgressService_WriteUpdatesServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*StatusUpdate, error)
	grpc.ServerStream
}

type progressServiceWriteUpdatesServer struct {
	grpc.ServerStream
}

func (x *progressServiceWriteUpdatesServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *progressServiceWriteUpdatesServer) Recv() (*StatusUpdate, error) {
	m := new(StatusUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProgressService_ServiceDesc is the grpc.ServiceDesc for ProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "progrock.ProgressService",
	HandlerType: (*ProgressServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WriteUpdates",
			Handler:       _ProgressService_WriteUpdates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "progress.proto",
}
