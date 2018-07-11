// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	chat_service.proto

It has these top-level messages:
	SignInResponse
	SignOutResponse
	SignInRequest
	SignOutRequest
	ChatResponse
	ChatRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SignInResponse struct {
}

func (m *SignInResponse) Reset()                    { *m = SignInResponse{} }
func (m *SignInResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()               {}
func (*SignInResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SignOutResponse struct {
}

func (m *SignOutResponse) Reset()                    { *m = SignOutResponse{} }
func (m *SignOutResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignOutResponse) ProtoMessage()               {}
func (*SignOutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SignInRequest struct {
}

func (m *SignInRequest) Reset()                    { *m = SignInRequest{} }
func (m *SignInRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()               {}
func (*SignInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SignOutRequest struct {
}

func (m *SignOutRequest) Reset()                    { *m = SignOutRequest{} }
func (m *SignOutRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignOutRequest) ProtoMessage()               {}
func (*SignOutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ChatResponse struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Person string `protobuf:"bytes,2,opt,name=person" json:"person,omitempty"`
}

func (m *ChatResponse) Reset()                    { *m = ChatResponse{} }
func (m *ChatResponse) String() string            { return proto1.CompactTextString(m) }
func (*ChatResponse) ProtoMessage()               {}
func (*ChatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChatResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ChatResponse) GetPerson() string {
	if m != nil {
		return m.Person
	}
	return ""
}

type ChatRequest struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Person string `protobuf:"bytes,2,opt,name=person" json:"person,omitempty"`
}

func (m *ChatRequest) Reset()                    { *m = ChatRequest{} }
func (m *ChatRequest) String() string            { return proto1.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()               {}
func (*ChatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChatRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ChatRequest) GetPerson() string {
	if m != nil {
		return m.Person
	}
	return ""
}

func init() {
	proto1.RegisterType((*SignInResponse)(nil), "proto.SignInResponse")
	proto1.RegisterType((*SignOutResponse)(nil), "proto.SignOutResponse")
	proto1.RegisterType((*SignInRequest)(nil), "proto.SignInRequest")
	proto1.RegisterType((*SignOutRequest)(nil), "proto.SignOutRequest")
	proto1.RegisterType((*ChatResponse)(nil), "proto.ChatResponse")
	proto1.RegisterType((*ChatRequest)(nil), "proto.ChatRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatService service

type ChatServiceClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	Chattering(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatteringClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := grpc.Invoke(ctx, "/proto.ChatService/SignIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := grpc.Invoke(ctx, "/proto.ChatService/SignOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Chattering(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatteringClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatService_serviceDesc.Streams[0], c.cc, "/proto.ChatService/Chattering", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatteringClient{stream}
	return x, nil
}

type ChatService_ChatteringClient interface {
	Send(*ChatRequest) error
	Recv() (*ChatResponse, error)
	grpc.ClientStream
}

type chatServiceChatteringClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatteringClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatteringClient) Recv() (*ChatResponse, error) {
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChatService service

type ChatServiceServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	Chattering(ChatService_ChatteringServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Chattering_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Chattering(&chatServiceChatteringServer{stream})
}

type ChatService_ChatteringServer interface {
	Send(*ChatResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type chatServiceChatteringServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatteringServer) Send(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatteringServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _ChatService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _ChatService_SignOut_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chattering",
			Handler:       _ChatService_Chattering_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat_service.proto",
}

func init() { proto1.RegisterFile("chat_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2c,
	0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x4a, 0x02, 0x5c, 0x7c, 0xc1, 0x99, 0xe9, 0x79, 0x9e, 0x79, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x82, 0x5c, 0xfc, 0x20, 0x11, 0xff, 0xd2, 0x12, 0xb8, 0x10, 0x3f,
	0x17, 0x2f, 0x4c, 0x51, 0x61, 0x69, 0x6a, 0x71, 0x09, 0x4c, 0x17, 0x58, 0x0d, 0x44, 0xc4, 0x8a,
	0x8b, 0xc7, 0x39, 0x23, 0x11, 0xae, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x48, 0x2d, 0x2a, 0xce,
	0xcf, 0x93, 0x60, 0x02, 0x8b, 0x42, 0x79, 0x4a, 0x96, 0x5c, 0xdc, 0x10, 0xbd, 0x60, 0xa3, 0x48,
	0xd1, 0x6a, 0xb4, 0x93, 0x11, 0xa2, 0x37, 0x18, 0xe2, 0x37, 0x21, 0x53, 0x2e, 0x36, 0x88, 0x4b,
	0x85, 0x44, 0x20, 0xfe, 0xd4, 0x43, 0x71, 0xb8, 0x94, 0x28, 0x9a, 0x28, 0xd4, 0xb5, 0x16, 0x5c,
	0xec, 0x50, 0xff, 0x08, 0x21, 0xab, 0x40, 0xf8, 0x4f, 0x4a, 0x0c, 0x5d, 0x18, 0xaa, 0xd3, 0x92,
	0x8b, 0x0b, 0x64, 0x7f, 0x49, 0x6a, 0x51, 0x66, 0x5e, 0xba, 0x90, 0x10, 0x54, 0x15, 0x92, 0x77,
	0xa4, 0x84, 0x51, 0xc4, 0x20, 0xda, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0xe2, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xad, 0x14, 0xac, 0x9e, 0x01, 0x00, 0x00,
}
