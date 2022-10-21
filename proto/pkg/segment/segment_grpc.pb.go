// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: segment.proto

package segment

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

// SegmentServerClient is the client API for SegmentServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SegmentServerClient interface {
	Start(ctx context.Context, in *StartSegmentServerRequest, opts ...grpc.CallOption) (*StartSegmentServerResponse, error)
	Stop(ctx context.Context, in *StopSegmentServerRequest, opts ...grpc.CallOption) (*StopSegmentServerResponse, error)
	CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveBlock(ctx context.Context, in *RemoveBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...grpc.CallOption) (*GetBlockInfoResponse, error)
	ActivateSegment(ctx context.Context, in *ActivateSegmentRequest, opts ...grpc.CallOption) (*ActivateSegmentResponse, error)
	InactivateSegment(ctx context.Context, in *InactivateSegmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppendToBlock(ctx context.Context, in *AppendToBlockRequest, opts ...grpc.CallOption) (*AppendToBlockResponse, error)
	ReadFromBlock(ctx context.Context, in *ReadFromBlockRequest, opts ...grpc.CallOption) (*ReadFromBlockResponse, error)
	LookupOffsetInBlock(ctx context.Context, in *LookupOffsetInBlockRequest, opts ...grpc.CallOption) (*LookupOffsetInBlockResponse, error)
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type segmentServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSegmentServerClient(cc grpc.ClientConnInterface) SegmentServerClient {
	return &segmentServerClient{cc}
}

func (c *segmentServerClient) Start(ctx context.Context, in *StartSegmentServerRequest, opts ...grpc.CallOption) (*StartSegmentServerResponse, error) {
	out := new(StartSegmentServerResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) Stop(ctx context.Context, in *StopSegmentServerRequest, opts ...grpc.CallOption) (*StopSegmentServerResponse, error) {
	out := new(StopSegmentServerResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/CreateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) RemoveBlock(ctx context.Context, in *RemoveBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/RemoveBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...grpc.CallOption) (*GetBlockInfoResponse, error) {
	out := new(GetBlockInfoResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/GetBlockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) ActivateSegment(ctx context.Context, in *ActivateSegmentRequest, opts ...grpc.CallOption) (*ActivateSegmentResponse, error) {
	out := new(ActivateSegmentResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/ActivateSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) InactivateSegment(ctx context.Context, in *InactivateSegmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/InactivateSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) AppendToBlock(ctx context.Context, in *AppendToBlockRequest, opts ...grpc.CallOption) (*AppendToBlockResponse, error) {
	out := new(AppendToBlockResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/AppendToBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) ReadFromBlock(ctx context.Context, in *ReadFromBlockRequest, opts ...grpc.CallOption) (*ReadFromBlockResponse, error) {
	out := new(ReadFromBlockResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/ReadFromBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) LookupOffsetInBlock(ctx context.Context, in *LookupOffsetInBlockRequest, opts ...grpc.CallOption) (*LookupOffsetInBlockResponse, error) {
	out := new(LookupOffsetInBlockResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/LookupOffsetInBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServerClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/linkall.vanus.segment.SegmentServer/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SegmentServerServer is the server API for SegmentServer service.
// All implementations should embed UnimplementedSegmentServerServer
// for forward compatibility
type SegmentServerServer interface {
	Start(context.Context, *StartSegmentServerRequest) (*StartSegmentServerResponse, error)
	Stop(context.Context, *StopSegmentServerRequest) (*StopSegmentServerResponse, error)
	CreateBlock(context.Context, *CreateBlockRequest) (*emptypb.Empty, error)
	RemoveBlock(context.Context, *RemoveBlockRequest) (*emptypb.Empty, error)
	GetBlockInfo(context.Context, *GetBlockInfoRequest) (*GetBlockInfoResponse, error)
	ActivateSegment(context.Context, *ActivateSegmentRequest) (*ActivateSegmentResponse, error)
	InactivateSegment(context.Context, *InactivateSegmentRequest) (*emptypb.Empty, error)
	AppendToBlock(context.Context, *AppendToBlockRequest) (*AppendToBlockResponse, error)
	ReadFromBlock(context.Context, *ReadFromBlockRequest) (*ReadFromBlockResponse, error)
	LookupOffsetInBlock(context.Context, *LookupOffsetInBlockRequest) (*LookupOffsetInBlockResponse, error)
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
}

// UnimplementedSegmentServerServer should be embedded to have forward compatible implementations.
type UnimplementedSegmentServerServer struct {
}

func (UnimplementedSegmentServerServer) Start(context.Context, *StartSegmentServerRequest) (*StartSegmentServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSegmentServerServer) Stop(context.Context, *StopSegmentServerRequest) (*StopSegmentServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSegmentServerServer) CreateBlock(context.Context, *CreateBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlock not implemented")
}
func (UnimplementedSegmentServerServer) RemoveBlock(context.Context, *RemoveBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlock not implemented")
}
func (UnimplementedSegmentServerServer) GetBlockInfo(context.Context, *GetBlockInfoRequest) (*GetBlockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockInfo not implemented")
}
func (UnimplementedSegmentServerServer) ActivateSegment(context.Context, *ActivateSegmentRequest) (*ActivateSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateSegment not implemented")
}
func (UnimplementedSegmentServerServer) InactivateSegment(context.Context, *InactivateSegmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InactivateSegment not implemented")
}
func (UnimplementedSegmentServerServer) AppendToBlock(context.Context, *AppendToBlockRequest) (*AppendToBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendToBlock not implemented")
}
func (UnimplementedSegmentServerServer) ReadFromBlock(context.Context, *ReadFromBlockRequest) (*ReadFromBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFromBlock not implemented")
}
func (UnimplementedSegmentServerServer) LookupOffsetInBlock(context.Context, *LookupOffsetInBlockRequest) (*LookupOffsetInBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupOffsetInBlock not implemented")
}
func (UnimplementedSegmentServerServer) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

// UnsafeSegmentServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SegmentServerServer will
// result in compilation errors.
type UnsafeSegmentServerServer interface {
	mustEmbedUnimplementedSegmentServerServer()
}

func RegisterSegmentServerServer(s grpc.ServiceRegistrar, srv SegmentServerServer) {
	s.RegisterService(&SegmentServer_ServiceDesc, srv)
}

func _SegmentServer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSegmentServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).Start(ctx, req.(*StartSegmentServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopSegmentServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).Stop(ctx, req.(*StopSegmentServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_CreateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).CreateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/CreateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).CreateBlock(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_RemoveBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).RemoveBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/RemoveBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).RemoveBlock(ctx, req.(*RemoveBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_GetBlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).GetBlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/GetBlockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).GetBlockInfo(ctx, req.(*GetBlockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_ActivateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).ActivateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/ActivateSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).ActivateSegment(ctx, req.(*ActivateSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_InactivateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InactivateSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).InactivateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/InactivateSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).InactivateSegment(ctx, req.(*InactivateSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_AppendToBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendToBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).AppendToBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/AppendToBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).AppendToBlock(ctx, req.(*AppendToBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_ReadFromBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFromBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).ReadFromBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/ReadFromBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).ReadFromBlock(ctx, req.(*ReadFromBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_LookupOffsetInBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupOffsetInBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).LookupOffsetInBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/LookupOffsetInBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).LookupOffsetInBlock(ctx, req.(*LookupOffsetInBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentServer_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkall.vanus.segment.SegmentServer/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServerServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SegmentServer_ServiceDesc is the grpc.ServiceDesc for SegmentServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SegmentServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkall.vanus.segment.SegmentServer",
	HandlerType: (*SegmentServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _SegmentServer_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _SegmentServer_Stop_Handler,
		},
		{
			MethodName: "CreateBlock",
			Handler:    _SegmentServer_CreateBlock_Handler,
		},
		{
			MethodName: "RemoveBlock",
			Handler:    _SegmentServer_RemoveBlock_Handler,
		},
		{
			MethodName: "GetBlockInfo",
			Handler:    _SegmentServer_GetBlockInfo_Handler,
		},
		{
			MethodName: "ActivateSegment",
			Handler:    _SegmentServer_ActivateSegment_Handler,
		},
		{
			MethodName: "InactivateSegment",
			Handler:    _SegmentServer_InactivateSegment_Handler,
		},
		{
			MethodName: "AppendToBlock",
			Handler:    _SegmentServer_AppendToBlock_Handler,
		},
		{
			MethodName: "ReadFromBlock",
			Handler:    _SegmentServer_ReadFromBlock_Handler,
		},
		{
			MethodName: "LookupOffsetInBlock",
			Handler:    _SegmentServer_LookupOffsetInBlock_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _SegmentServer_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "segment.proto",
}