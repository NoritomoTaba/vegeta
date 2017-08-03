// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/collection.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/collection.proto

It has these top-level messages:
	RequestFromDevice
	ResultResponse
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestFromDevice struct {
	Timestamp string `protobuf:"bytes,1,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	DeviceID  string `protobuf:"bytes,3,opt,name=DeviceID" json:"DeviceID,omitempty"`
	IPAddress string `protobuf:"bytes,4,opt,name=IPAddress" json:"IPAddress,omitempty"`
	IsSuccess bool   `protobuf:"varint,5,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
}

func (m *RequestFromDevice) Reset()                    { *m = RequestFromDevice{} }
func (m *RequestFromDevice) String() string            { return proto.CompactTextString(m) }
func (*RequestFromDevice) ProtoMessage()               {}
func (*RequestFromDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestFromDevice) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RequestFromDevice) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *RequestFromDevice) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *RequestFromDevice) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *RequestFromDevice) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type ResultResponse struct {
	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
}

func (m *ResultResponse) Reset()                    { *m = ResultResponse{} }
func (m *ResultResponse) String() string            { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()               {}
func (*ResultResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResultResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*RequestFromDevice)(nil), "protos.RequestFromDevice")
	proto.RegisterType((*ResultResponse)(nil), "protos.ResultResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Collection service

type CollectionClient interface {
	RecvData(ctx context.Context, opts ...grpc.CallOption) (Collection_RecvDataClient, error)
}

type collectionClient struct {
	cc *grpc.ClientConn
}

func NewCollectionClient(cc *grpc.ClientConn) CollectionClient {
	return &collectionClient{cc}
}

func (c *collectionClient) RecvData(ctx context.Context, opts ...grpc.CallOption) (Collection_RecvDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Collection_serviceDesc.Streams[0], c.cc, "/protos.Collection/RecvData", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectionRecvDataClient{stream}
	return x, nil
}

type Collection_RecvDataClient interface {
	Send(*RequestFromDevice) error
	CloseAndRecv() (*ResultResponse, error)
	grpc.ClientStream
}

type collectionRecvDataClient struct {
	grpc.ClientStream
}

func (x *collectionRecvDataClient) Send(m *RequestFromDevice) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collectionRecvDataClient) CloseAndRecv() (*ResultResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Collection service

type CollectionServer interface {
	RecvData(Collection_RecvDataServer) error
}

func RegisterCollectionServer(s *grpc.Server, srv CollectionServer) {
	s.RegisterService(&_Collection_serviceDesc, srv)
}

func _Collection_RecvData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollectionServer).RecvData(&collectionRecvDataServer{stream})
}

type Collection_RecvDataServer interface {
	SendAndClose(*ResultResponse) error
	Recv() (*RequestFromDevice, error)
	grpc.ServerStream
}

type collectionRecvDataServer struct {
	grpc.ServerStream
}

func (x *collectionRecvDataServer) SendAndClose(m *ResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collectionRecvDataServer) Recv() (*RequestFromDevice, error) {
	m := new(RequestFromDevice)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Collection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Collection",
	HandlerType: (*CollectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecvData",
			Handler:       _Collection_RecvData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/collection.proto",
}

func init() { proto.RegisterFile("protos/collection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x8d, 0xbf, 0xed, 0x2c, 0x04, 0xb3, 0xd0, 0x78, 0x51, 0xb8, 0x74, 0xd5, 0xd5, 0x2d,
	0xe8, 0x13, 0x14, 0x8b, 0xd0, 0x95, 0x25, 0xba, 0x97, 0x98, 0xce, 0x22, 0x90, 0x36, 0xb5, 0x93,
	0x16, 0x7c, 0x1f, 0x1f, 0x54, 0xda, 0x60, 0x2b, 0x75, 0x39, 0xdf, 0xc7, 0x9c, 0x61, 0x0e, 0xdc,
	0x74, 0xbd, 0xf3, 0x8e, 0x32, 0xed, 0xac, 0x45, 0xed, 0x8d, 0x6b, 0x0f, 0x33, 0xe1, 0xe7, 0x41,
	0x24, 0xdf, 0x0c, 0xae, 0x24, 0x7e, 0x0e, 0x48, 0xfe, 0xb9, 0x77, 0x4d, 0x81, 0xa3, 0xd1, 0xc8,
	0xef, 0x20, 0x7e, 0x33, 0x0d, 0x92, 0x57, 0x4d, 0x27, 0xd8, 0x9e, 0xa5, 0xb1, 0x5c, 0x01, 0x17,
	0x70, 0x51, 0xa9, 0x2f, 0xeb, 0x54, 0x2d, 0x8e, 0x67, 0xf7, 0x3b, 0xf2, 0x1d, 0x44, 0x21, 0xa1,
	0x2c, 0xc4, 0xc9, 0xac, 0x96, 0x79, 0xca, 0x2c, 0xab, 0xbc, 0xae, 0x7b, 0x24, 0x12, 0xa7, 0x21,
	0x73, 0x01, 0xfc, 0x1e, 0xc0, 0xd0, 0x3b, 0x0d, 0x5a, 0x4f, 0xfa, 0x6c, 0xcf, 0xd2, 0x48, 0xc6,
	0x86, 0x5e, 0x03, 0x48, 0x32, 0xb8, 0x94, 0x48, 0x83, 0xf5, 0x12, 0xa9, 0x73, 0x2d, 0xe1, 0x66,
	0x81, 0x6d, 0x16, 0x1e, 0x5e, 0x00, 0x9e, 0x96, 0x9f, 0x79, 0x0e, 0x91, 0x44, 0x3d, 0x16, 0xca,
	0x2b, 0x7e, 0x1b, 0x1a, 0xa0, 0xc3, 0xbf, 0xb7, 0x77, 0xd7, 0xab, 0xfa, 0x7b, 0x2b, 0x39, 0x4a,
	0xd9, 0x47, 0x28, 0xec, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x67, 0xf0, 0x16, 0x52, 0x01,
	0x00, 0x00,
}