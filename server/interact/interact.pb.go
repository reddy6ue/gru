// Code generated by protoc-gen-go.
// source: interact.proto
// DO NOT EDIT!

/*
Package interact is a generated protocol buffer package.

It is generated from these files:
	interact.proto

It has these top-level messages:
	ServerStatus
	ClientStatus
	Token
	Session
	Req
	Question
	Answer
	Response
	Status
*/
package interact

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
const _ = proto.ProtoPackageIsVersion1

type ServerStatus struct {
	TimeLeft string `protobuf:"bytes,1,opt,name=timeLeft" json:"timeLeft,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientStatus struct {
	CurrQuestion string `protobuf:"bytes,1,opt,name=currQuestion" json:"currQuestion,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *ClientStatus) Reset()                    { *m = ClientStatus{} }
func (m *ClientStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()               {}
func (*ClientStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Token struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Session struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Req struct {
	Repeat   bool   `protobuf:"varint,1,opt,name=repeat" json:"repeat,omitempty"`
	Sid      string `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	TestType string `protobuf:"bytes,4,opt,name=testType" json:"testType,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Question struct {
	Id         string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str        string    `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
	Options    []*Answer `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	IsMultiple bool      `protobuf:"varint,4,opt,name=isMultiple" json:"isMultiple,omitempty"`
	Positive   float32   `protobuf:"fixed32,6,opt,name=positive" json:"positive,omitempty"`
	Negative   float32   `protobuf:"fixed32,7,opt,name=negative" json:"negative,omitempty"`
	Totscore   float32   `protobuf:"fixed32,8,opt,name=totscore" json:"totscore,omitempty"`
}

func (m *Question) Reset()                    { *m = Question{} }
func (m *Question) String() string            { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()               {}
func (*Question) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Question) GetOptions() []*Answer {
	if m != nil {
		return m.Options
	}
	return nil
}

type Answer struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str string `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
}

func (m *Answer) Reset()                    { *m = Answer{} }
func (m *Answer) String() string            { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()               {}
func (*Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Response struct {
	Qid      string   `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Aid      []string `protobuf:"bytes,2,rep,name=aid" json:"aid,omitempty"`
	Sid      string   `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Token    string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	TestType string   `protobuf:"bytes,5,opt,name=testType" json:"testType,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Status struct {
	Status int64 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*ServerStatus)(nil), "interact.ServerStatus")
	proto.RegisterType((*ClientStatus)(nil), "interact.ClientStatus")
	proto.RegisterType((*Token)(nil), "interact.Token")
	proto.RegisterType((*Session)(nil), "interact.Session")
	proto.RegisterType((*Req)(nil), "interact.Req")
	proto.RegisterType((*Question)(nil), "interact.Question")
	proto.RegisterType((*Answer)(nil), "interact.Answer")
	proto.RegisterType((*Response)(nil), "interact.Response")
	proto.RegisterType((*Status)(nil), "interact.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GruQuiz service

type GruQuizClient interface {
	Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error)
	GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error)
	SendAnswer(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Status, error)
	StreamChan(ctx context.Context, opts ...grpc.CallOption) (GruQuiz_StreamChanClient, error)
}

type gruQuizClient struct {
	cc *grpc.ClientConn
}

func NewGruQuizClient(cc *grpc.ClientConn) GruQuizClient {
	return &gruQuizClient{cc}
}

func (c *gruQuizClient) Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/GetQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) SendAnswer(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/SendAnswer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) StreamChan(ctx context.Context, opts ...grpc.CallOption) (GruQuiz_StreamChanClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GruQuiz_serviceDesc.Streams[0], c.cc, "/interact.GruQuiz/StreamChan", opts...)
	if err != nil {
		return nil, err
	}
	x := &gruQuizStreamChanClient{stream}
	return x, nil
}

type GruQuiz_StreamChanClient interface {
	Send(*ClientStatus) error
	Recv() (*ServerStatus, error)
	grpc.ClientStream
}

type gruQuizStreamChanClient struct {
	grpc.ClientStream
}

func (x *gruQuizStreamChanClient) Send(m *ClientStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gruQuizStreamChanClient) Recv() (*ServerStatus, error) {
	m := new(ServerStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GruQuiz service

type GruQuizServer interface {
	Authenticate(context.Context, *Token) (*Session, error)
	GetQuestion(context.Context, *Req) (*Question, error)
	SendAnswer(context.Context, *Response) (*Status, error)
	StreamChan(GruQuiz_StreamChanServer) error
}

func RegisterGruQuizServer(s *grpc.Server, srv GruQuizServer) {
	s.RegisterService(&_GruQuiz_serviceDesc, srv)
}

func _GruQuiz_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Authenticate(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).GetQuestion(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_SendAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).SendAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/SendAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).SendAnswer(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_StreamChan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GruQuizServer).StreamChan(&gruQuizStreamChanServer{stream})
}

type GruQuiz_StreamChanServer interface {
	Send(*ServerStatus) error
	Recv() (*ClientStatus, error)
	grpc.ServerStream
}

type gruQuizStreamChanServer struct {
	grpc.ServerStream
}

func (x *gruQuizStreamChanServer) Send(m *ServerStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gruQuizStreamChanServer) Recv() (*ClientStatus, error) {
	m := new(ClientStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GruQuiz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interact.GruQuiz",
	HandlerType: (*GruQuizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _GruQuiz_Authenticate_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _GruQuiz_GetQuestion_Handler,
		},
		{
			MethodName: "SendAnswer",
			Handler:    _GruQuiz_SendAnswer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChan",
			Handler:       _GruQuiz_StreamChan_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x72, 0xd3, 0x30,
	0x10, 0xae, 0xe3, 0xc6, 0x31, 0xdb, 0x50, 0x8a, 0x86, 0x09, 0xc6, 0x07, 0xa6, 0xa3, 0x53, 0xa6,
	0x87, 0x0e, 0x13, 0xfa, 0x02, 0x6d, 0x0f, 0xe5, 0x00, 0x87, 0x3a, 0x7d, 0x01, 0x91, 0x2e, 0x54,
	0x43, 0x6a, 0xbb, 0xd2, 0xba, 0x0c, 0x3c, 0x23, 0xaf, 0xc2, 0x3b, 0xa0, 0x3f, 0x5b, 0x26, 0x03,
	0xd3, 0x9b, 0xbe, 0xfd, 0xd1, 0xf7, 0xed, 0xa7, 0x15, 0x1c, 0xca, 0x9a, 0x50, 0x89, 0x0d, 0x9d,
	0xb6, 0xaa, 0xa1, 0x86, 0xe5, 0x3d, 0xe6, 0x17, 0x30, 0x5f, 0xa3, 0x7a, 0x44, 0xb5, 0x26, 0x41,
	0x9d, 0x66, 0x25, 0xe4, 0x24, 0xef, 0xf1, 0x23, 0x7e, 0xa1, 0x22, 0x39, 0x4e, 0x96, 0xcf, 0xaa,
	0x01, 0xb3, 0x05, 0x64, 0xda, 0x55, 0x15, 0x13, 0x97, 0x09, 0x88, 0x7f, 0x80, 0xf9, 0xe5, 0x56,
	0x62, 0x4d, 0xe1, 0x0e, 0x0e, 0xf3, 0x4d, 0xa7, 0xd4, 0x75, 0x87, 0x9a, 0x64, 0x53, 0x87, 0x7b,
	0xfe, 0x8a, 0xb1, 0x57, 0x30, 0xa5, 0xe6, 0x1b, 0xd6, 0xe1, 0x2a, 0x0f, 0xf8, 0x6b, 0x98, 0xde,
	0xd8, 0x03, 0x3b, 0x84, 0x89, 0xbc, 0x0d, 0x8d, 0xe6, 0xc4, 0xdf, 0xc0, 0x6c, 0x8d, 0x5a, 0xdb,
	0xce, 0xdd, 0x94, 0x80, 0xb4, 0xc2, 0x07, 0x2b, 0x4e, 0x61, 0x8b, 0xc2, 0xcb, 0xce, 0xab, 0x80,
	0xd8, 0x11, 0xa4, 0xda, 0xd4, 0x7b, 0x1a, 0x7b, 0x8c, 0xd4, 0xe9, 0x88, 0xda, 0x0d, 0x6e, 0xb4,
	0xdd, 0xfc, 0x68, 0xb1, 0xd8, 0x0f, 0x83, 0x07, 0xcc, 0x7f, 0x25, 0x90, 0x0f, 0xca, 0x77, 0xf8,
	0x1d, 0x01, 0xa9, 0x81, 0x80, 0x14, 0x3b, 0x81, 0x59, 0xd3, 0xda, 0x5a, 0x6d, 0x28, 0xd2, 0xe5,
	0xc1, 0xea, 0xe8, 0x74, 0xf0, 0xff, 0xbc, 0xd6, 0xdf, 0x51, 0x55, 0x7d, 0x01, 0x7b, 0x0b, 0x20,
	0xf5, 0xa7, 0x6e, 0x4b, 0xb2, 0xdd, 0x7a, 0xe2, 0xbc, 0x1a, 0x45, 0xac, 0xac, 0xb6, 0xd1, 0x92,
	0xe4, 0x23, 0x16, 0x99, 0xc9, 0x4e, 0xaa, 0x01, 0xdb, 0x5c, 0x8d, 0x5f, 0x85, 0xcb, 0xcd, 0x7c,
	0xae, 0xc7, 0x6e, 0x9c, 0x86, 0xf4, 0xa6, 0x51, 0x58, 0xe4, 0x3e, 0xd7, 0x63, 0x7e, 0x02, 0x99,
	0x97, 0xf1, 0xf4, 0x2c, 0x5c, 0x41, 0x5e, 0xa1, 0x6e, 0x8d, 0x54, 0xb4, 0xd9, 0x87, 0xa1, 0xdc,
	0x1e, 0x6d, 0x44, 0x38, 0x73, 0x53, 0x1b, 0x11, 0xe1, 0x06, 0x13, 0x49, 0xff, 0x61, 0xf7, 0xfe,
	0xff, 0xec, 0x9e, 0xee, 0xd8, 0x7d, 0x0c, 0x59, 0xd8, 0xa4, 0xb8, 0x71, 0x96, 0x34, 0xed, 0x37,
	0x6e, 0xf5, 0x3b, 0x81, 0xd9, 0x95, 0xea, 0xae, 0x3b, 0xf9, 0x93, 0x9d, 0xc1, 0xfc, 0xbc, 0xa3,
	0x3b, 0xb3, 0x7e, 0x72, 0x23, 0x08, 0xd9, 0x8b, 0x68, 0xb6, 0xdb, 0xa5, 0xf2, 0x65, 0x0c, 0x84,
	0x1d, 0xe2, 0x7b, 0x6c, 0x05, 0x07, 0x57, 0x48, 0xc3, 0xa3, 0x3e, 0x8f, 0x35, 0x66, 0x99, 0x4a,
	0x16, 0x61, 0x5f, 0x62, 0x7a, 0xce, 0x00, 0xd6, 0x58, 0xdf, 0x06, 0xef, 0xd8, 0xb8, 0xc5, 0x3b,
	0x54, 0x8e, 0x1e, 0xda, 0x4f, 0x60, 0xba, 0x2e, 0x4c, 0x17, 0x29, 0x14, 0xf7, 0x97, 0x77, 0xa2,
	0x66, 0x8b, 0x58, 0x31, 0xfe, 0x33, 0xe5, 0x62, 0x2c, 0x32, 0xfe, 0x47, 0xbe, 0xb7, 0x4c, 0xde,
	0x25, 0x9f, 0x33, 0xf7, 0x6d, 0xdf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x65, 0xc1, 0x09, 0x78,
	0xc8, 0x03, 0x00, 0x00,
}
