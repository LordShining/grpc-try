// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	Working(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Working(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error) {
	out := new(WorkReply)
	err := c.cc.Invoke(ctx, "/pb.Worker/Working", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	Working(context.Context, *WorkRequest) (*WorkReply, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) Working(context.Context, *WorkRequest) (*WorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Working not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Working_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Working(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Worker/Working",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Working(ctx, req.(*WorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Working",
			Handler:    _Worker_Working_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/grpc_try.proto",
}
