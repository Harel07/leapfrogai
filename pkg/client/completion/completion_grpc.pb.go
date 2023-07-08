// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: completion/completion.proto

package completion

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

// CompletionServiceClient is the client API for CompletionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompletionServiceClient interface {
	Complete(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error)
}

type completionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionServiceClient(cc grpc.ClientConnInterface) CompletionServiceClient {
	return &completionServiceClient{cc}
}

func (c *completionServiceClient) Complete(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error) {
	out := new(CompletionResponse)
	err := c.cc.Invoke(ctx, "/completion.CompletionService/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServiceServer is the server API for CompletionService service.
// All implementations must embed UnimplementedCompletionServiceServer
// for forward compatibility
type CompletionServiceServer interface {
	Complete(context.Context, *CompletionRequest) (*CompletionResponse, error)
	mustEmbedUnimplementedCompletionServiceServer()
}

// UnimplementedCompletionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompletionServiceServer struct {
}

func (UnimplementedCompletionServiceServer) Complete(context.Context, *CompletionRequest) (*CompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedCompletionServiceServer) mustEmbedUnimplementedCompletionServiceServer() {}

// UnsafeCompletionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompletionServiceServer will
// result in compilation errors.
type UnsafeCompletionServiceServer interface {
	mustEmbedUnimplementedCompletionServiceServer()
}

func RegisterCompletionServiceServer(s grpc.ServiceRegistrar, srv CompletionServiceServer) {
	s.RegisterService(&CompletionService_ServiceDesc, srv)
}

func _CompletionService_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/completion.CompletionService/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).Complete(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompletionService_ServiceDesc is the grpc.ServiceDesc for CompletionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompletionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "completion.CompletionService",
	HandlerType: (*CompletionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Complete",
			Handler:    _CompletionService_Complete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "completion/completion.proto",
}

// CompletionStreamServiceClient is the client API for CompletionStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompletionStreamServiceClient interface {
	CompleteStream(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (CompletionStreamService_CompleteStreamClient, error)
}

type completionStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionStreamServiceClient(cc grpc.ClientConnInterface) CompletionStreamServiceClient {
	return &completionStreamServiceClient{cc}
}

func (c *completionStreamServiceClient) CompleteStream(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (CompletionStreamService_CompleteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompletionStreamService_ServiceDesc.Streams[0], "/completion.CompletionStreamService/CompleteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &completionStreamServiceCompleteStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompletionStreamService_CompleteStreamClient interface {
	Recv() (*CompletionResponse, error)
	grpc.ClientStream
}

type completionStreamServiceCompleteStreamClient struct {
	grpc.ClientStream
}

func (x *completionStreamServiceCompleteStreamClient) Recv() (*CompletionResponse, error) {
	m := new(CompletionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompletionStreamServiceServer is the server API for CompletionStreamService service.
// All implementations must embed UnimplementedCompletionStreamServiceServer
// for forward compatibility
type CompletionStreamServiceServer interface {
	CompleteStream(*CompletionRequest, CompletionStreamService_CompleteStreamServer) error
	mustEmbedUnimplementedCompletionStreamServiceServer()
}

// UnimplementedCompletionStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompletionStreamServiceServer struct {
}

func (UnimplementedCompletionStreamServiceServer) CompleteStream(*CompletionRequest, CompletionStreamService_CompleteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CompleteStream not implemented")
}
func (UnimplementedCompletionStreamServiceServer) mustEmbedUnimplementedCompletionStreamServiceServer() {
}

// UnsafeCompletionStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompletionStreamServiceServer will
// result in compilation errors.
type UnsafeCompletionStreamServiceServer interface {
	mustEmbedUnimplementedCompletionStreamServiceServer()
}

func RegisterCompletionStreamServiceServer(s grpc.ServiceRegistrar, srv CompletionStreamServiceServer) {
	s.RegisterService(&CompletionStreamService_ServiceDesc, srv)
}

func _CompletionStreamService_CompleteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompletionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompletionStreamServiceServer).CompleteStream(m, &completionStreamServiceCompleteStreamServer{stream})
}

type CompletionStreamService_CompleteStreamServer interface {
	Send(*CompletionResponse) error
	grpc.ServerStream
}

type completionStreamServiceCompleteStreamServer struct {
	grpc.ServerStream
}

func (x *completionStreamServiceCompleteStreamServer) Send(m *CompletionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CompletionStreamService_ServiceDesc is the grpc.ServiceDesc for CompletionStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompletionStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "completion.CompletionStreamService",
	HandlerType: (*CompletionStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CompleteStream",
			Handler:       _CompletionStreamService_CompleteStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "completion/completion.proto",
}
