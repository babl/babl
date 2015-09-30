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
	In []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (m *BinRequest) Reset()         { *m = BinRequest{} }
func (m *BinRequest) String() string { return proto.CompactTextString(m) }
func (*BinRequest) ProtoMessage()    {}

type BinReply struct {
	Out    []byte          `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
	Status BinReply_Status `protobuf:"varint,2,opt,name=status,enum=babl.BinReply_Status" json:"status,omitempty"`
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

func _StringUpcase_IO_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(BinRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
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
