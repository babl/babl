// Code generated by protoc-gen-go.
// source: babl.proto
// DO NOT EDIT!

/*
Package babl is a generated protocol buffer package.

It is generated from these files:
	babl.proto

It has these top-level messages:
	BinRequest
	BinReply
*/
package babl

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

type BinReply_Status int32

const (
	BinReply_SUCCESS BinReply_Status = 0
	BinReply_ERROR   BinReply_Status = 1
)

var BinReply_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
}
var BinReply_Status_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
}

func (x BinReply_Status) String() string {
	return proto.EnumName(BinReply_Status_name, int32(x))
}

type BinRequest struct {
	In  []byte            `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Env map[string]string `protobuf:"bytes,2,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BinRequest) Reset()         { *m = BinRequest{} }
func (m *BinRequest) String() string { return proto.CompactTextString(m) }
func (*BinRequest) ProtoMessage()    {}

func (m *BinRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type BinReply struct {
	Out    []byte          `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
	Status BinReply_Status `protobuf:"varint,2,opt,name=status,enum=babl.BinReply_Status" json:"status,omitempty"`
	Error  string          `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *BinReply) Reset()         { *m = BinReply{} }
func (m *BinReply) String() string { return proto.CompactTextString(m) }
func (*BinReply) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("babl.BinReply_Status", BinReply_Status_name, BinReply_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for StringUpcase service

type StringUpcaseClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
}

type stringUpcaseClient struct {
	cc *grpc.ClientConn
}

func NewStringUpcaseClient(cc *grpc.ClientConn) StringUpcaseClient {
	return &stringUpcaseClient{cc}
}

func (c *stringUpcaseClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.StringUpcase/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StringUpcase service

type StringUpcaseServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
}

func RegisterStringUpcaseServer(s *grpc.Server, srv StringUpcaseServer) {
	s.RegisterService(&_StringUpcase_serviceDesc, srv)
}

func _StringUpcase_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(StringUpcaseServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _StringUpcase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.StringUpcase",
	HandlerType: (*StringUpcaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _StringUpcase_IO_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Download service

type DownloadClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
}

type downloadClient struct {
	cc *grpc.ClientConn
}

func NewDownloadClient(cc *grpc.ClientConn) DownloadClient {
	return &downloadClient{cc}
}

func (c *downloadClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.Download/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Download service

type DownloadServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
}

func RegisterDownloadServer(s *grpc.Server, srv DownloadServer) {
	s.RegisterService(&_Download_serviceDesc, srv)
}

func _Download_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DownloadServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Download_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.Download",
	HandlerType: (*DownloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _Download_IO_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for S3 service

type S3Client interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
}

type s3Client struct {
	cc *grpc.ClientConn
}

func NewS3Client(cc *grpc.ClientConn) S3Client {
	return &s3Client{cc}
}

func (c *s3Client) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.S3/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S3 service

type S3Server interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
}

func RegisterS3Server(s *grpc.Server, srv S3Server) {
	s.RegisterService(&_S3_serviceDesc, srv)
}

func _S3_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(S3Server).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _S3_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.S3",
	HandlerType: (*S3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _S3_IO_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for TestFail service

type TestFailClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
}

type testFailClient struct {
	cc *grpc.ClientConn
}

func NewTestFailClient(cc *grpc.ClientConn) TestFailClient {
	return &testFailClient{cc}
}

func (c *testFailClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.TestFail/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestFail service

type TestFailServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
}

func RegisterTestFailServer(s *grpc.Server, srv TestFailServer) {
	s.RegisterService(&_TestFail_serviceDesc, srv)
}

func _TestFail_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TestFailServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _TestFail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.TestFail",
	HandlerType: (*TestFailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _TestFail_IO_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
