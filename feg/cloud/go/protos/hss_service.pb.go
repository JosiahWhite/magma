// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/hss_service.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/lte/cloud/go/protos"
	protos1 "magma/orc8r/lib/go/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("feg/protos/hss_service.proto", fileDescriptor_6adda26d69f7818f) }

var fileDescriptor_6adda26d69f7818f = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x0b, 0x82, 0x60, 0x40, 0xda, 0x06, 0x41, 0x36, 0xd5, 0x4b, 0x7f, 0x40, 0x16, 0xf4,
	0xe2, 0x4d, 0xad, 0x01, 0x2d, 0x78, 0x5b, 0xf4, 0xe0, 0x45, 0xb2, 0xc9, 0x6c, 0x0c, 0xec, 0x76,
	0xca, 0x24, 0xeb, 0xdf, 0xf2, 0x2f, 0x8a, 0xd9, 0x56, 0x5b, 0xdc, 0xbd, 0xf4, 0x14, 0x78, 0xef,
	0xf1, 0xcd, 0x9b, 0x24, 0xec, 0xa2, 0x02, 0x97, 0xaf, 0x09, 0x23, 0x86, 0xfc, 0x23, 0x84, 0xf7,
	0x00, 0xf4, 0xe9, 0x0d, 0xc8, 0x24, 0xf1, 0x93, 0x46, 0xbb, 0x46, 0xcb, 0x0a, 0x9c, 0xc8, 0x90,
	0xcc, 0x0d, 0x6d, 0xa3, 0x06, 0x9b, 0x06, 0x57, 0x5d, 0x4a, 0x5c, 0xd6, 0x11, 0xb6, 0x46, 0x68,
	0xcb, 0x60, 0xc8, 0x97, 0x40, 0xb6, 0xec, 0xec, 0xab, 0xaf, 0x23, 0x36, 0x7e, 0x2a, 0x8a, 0x07,
	0x5c, 0x55, 0xde, 0xb5, 0xa4, 0x23, 0x12, 0xbf, 0x65, 0xa7, 0xf7, 0xd6, 0x16, 0xbf, 0x61, 0x9e,
	0xc9, 0x6e, 0x54, 0x1d, 0x41, 0xfe, 0xc9, 0x4a, 0x47, 0x2d, 0xa6, 0x1b, 0x2b, 0x15, 0x90, 0xaf,
	0xe8, 0xed, 0x7c, 0xc4, 0xef, 0xd8, 0x44, 0x41, 0x0d, 0x11, 0x76, 0x18, 0xe7, 0xbd, 0x8c, 0xa5,
	0xea, 0x27, 0x28, 0x36, 0x79, 0x59, 0x5b, 0xbd, 0x47, 0x98, 0xf5, 0x12, 0xba, 0x58, 0x3f, 0x65,
	0xc9, 0xa6, 0x8f, 0x10, 0xf7, 0x1b, 0x0f, 0x17, 0x19, 0xde, 0x72, 0x3e, 0xe2, 0x0b, 0x36, 0x7e,
	0xf6, 0x61, 0x87, 0x15, 0xf8, 0xff, 0x91, 0x42, 0x0c, 0xb0, 0x0b, 0x88, 0x69, 0xa9, 0x33, 0x05,
	0x04, 0xce, 0x87, 0x08, 0x74, 0xe8, 0xd5, 0x2c, 0x66, 0x6f, 0x59, 0x52, 0xf3, 0x9f, 0xcf, 0x61,
	0x6a, 0x6c, 0x6d, 0xee, 0x70, 0xf3, 0xc2, 0xe5, 0x71, 0x3a, 0xaf, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0xb0, 0x29, 0xc9, 0x3a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HSSConfiguratorClient is the client API for HSSConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HSSConfiguratorClient interface {
	// Adds a new subscriber to the store.
	// Throws ALREADY_EXISTS if the subscriber already exists.
	//
	AddSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error)
	// Deletes an existing subscriber.
	// If the subscriber is not already present, this request is ignored.
	//
	DeleteSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error)
	// Updates an existing subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	UpdateSubscriber(ctx context.Context, in *protos.SubscriberUpdate, opts ...grpc.CallOption) (*protos1.Void, error)
	// Returns the SubscriberData for a subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	GetSubscriberData(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos.SubscriberData, error)
	// List the subscribers in the store.
	//
	ListSubscribers(ctx context.Context, in *protos1.Void, opts ...grpc.CallOption) (*protos.SubscriberIDSet, error)
	// De-register an authenticated subscriber
	DeregisterSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error)
}

type hSSConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewHSSConfiguratorClient(cc grpc.ClientConnInterface) HSSConfiguratorClient {
	return &hSSConfiguratorClient{cc}
}

func (c *hSSConfiguratorClient) AddSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/AddSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) DeleteSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/DeleteSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) UpdateSubscriber(ctx context.Context, in *protos.SubscriberUpdate, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/UpdateSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) GetSubscriberData(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos.SubscriberData, error) {
	out := new(protos.SubscriberData)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/GetSubscriberData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) ListSubscribers(ctx context.Context, in *protos1.Void, opts ...grpc.CallOption) (*protos.SubscriberIDSet, error) {
	out := new(protos.SubscriberIDSet)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/ListSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) DeregisterSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/DeregisterSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HSSConfiguratorServer is the server API for HSSConfigurator service.
type HSSConfiguratorServer interface {
	// Adds a new subscriber to the store.
	// Throws ALREADY_EXISTS if the subscriber already exists.
	//
	AddSubscriber(context.Context, *protos.SubscriberData) (*protos1.Void, error)
	// Deletes an existing subscriber.
	// If the subscriber is not already present, this request is ignored.
	//
	DeleteSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error)
	// Updates an existing subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	UpdateSubscriber(context.Context, *protos.SubscriberUpdate) (*protos1.Void, error)
	// Returns the SubscriberData for a subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	GetSubscriberData(context.Context, *protos.SubscriberID) (*protos.SubscriberData, error)
	// List the subscribers in the store.
	//
	ListSubscribers(context.Context, *protos1.Void) (*protos.SubscriberIDSet, error)
	// De-register an authenticated subscriber
	DeregisterSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error)
}

// UnimplementedHSSConfiguratorServer can be embedded to have forward compatible implementations.
type UnimplementedHSSConfiguratorServer struct {
}

func (*UnimplementedHSSConfiguratorServer) AddSubscriber(ctx context.Context, req *protos.SubscriberData) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) DeleteSubscriber(ctx context.Context, req *protos.SubscriberID) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) UpdateSubscriber(ctx context.Context, req *protos.SubscriberUpdate) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) GetSubscriberData(ctx context.Context, req *protos.SubscriberID) (*protos.SubscriberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberData not implemented")
}
func (*UnimplementedHSSConfiguratorServer) ListSubscribers(ctx context.Context, req *protos1.Void) (*protos.SubscriberIDSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscribers not implemented")
}
func (*UnimplementedHSSConfiguratorServer) DeregisterSubscriber(ctx context.Context, req *protos.SubscriberID) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterSubscriber not implemented")
}

func RegisterHSSConfiguratorServer(s *grpc.Server, srv HSSConfiguratorServer) {
	s.RegisterService(&_HSSConfigurator_serviceDesc, srv)
}

func _HSSConfigurator_AddSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).AddSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/AddSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).AddSubscriber(ctx, req.(*protos.SubscriberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_DeleteSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).DeleteSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/DeleteSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).DeleteSubscriber(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_UpdateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).UpdateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/UpdateSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).UpdateSubscriber(ctx, req.(*protos.SubscriberUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_GetSubscriberData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).GetSubscriberData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/GetSubscriberData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).GetSubscriberData(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_ListSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos1.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).ListSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/ListSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).ListSubscribers(ctx, req.(*protos1.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_DeregisterSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).DeregisterSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/DeregisterSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).DeregisterSubscriber(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

var _HSSConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.HSSConfigurator",
	HandlerType: (*HSSConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSubscriber",
			Handler:    _HSSConfigurator_AddSubscriber_Handler,
		},
		{
			MethodName: "DeleteSubscriber",
			Handler:    _HSSConfigurator_DeleteSubscriber_Handler,
		},
		{
			MethodName: "UpdateSubscriber",
			Handler:    _HSSConfigurator_UpdateSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriberData",
			Handler:    _HSSConfigurator_GetSubscriberData_Handler,
		},
		{
			MethodName: "ListSubscribers",
			Handler:    _HSSConfigurator_ListSubscribers_Handler,
		},
		{
			MethodName: "DeregisterSubscriber",
			Handler:    _HSSConfigurator_DeregisterSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/hss_service.proto",
}
