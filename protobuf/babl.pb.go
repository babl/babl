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
	Empty
	Pong
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

type BinRequest struct {
	Stdin []byte            `protobuf:"bytes,1,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Env   map[string]string `protobuf:"bytes,2,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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
	Stdout   []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Exitcode int32  `protobuf:"varint,3,opt,name=exitcode" json:"exitcode,omitempty"`
}

func (m *BinReply) Reset()         { *m = BinReply{} }
func (m *BinReply) String() string { return proto.CompactTextString(m) }
func (*BinReply) ProtoMessage()    {}

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}

type Pong struct {
	Val string `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for StringUpcase service

type StringUpcaseClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
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

func (c *stringUpcaseClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.StringUpcase/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StringUpcase service

type StringUpcaseServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
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

func _StringUpcase_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(StringUpcaseServer).Ping(ctx, in)
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
		{
			MethodName: "Ping",
			Handler:    _StringUpcase_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for StringAppend service

type StringAppendClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
}

type stringAppendClient struct {
	cc *grpc.ClientConn
}

func NewStringAppendClient(cc *grpc.ClientConn) StringAppendClient {
	return &stringAppendClient{cc}
}

func (c *stringAppendClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.StringAppend/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringAppendClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.StringAppend/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StringAppend service

type StringAppendServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
}

func RegisterStringAppendServer(s *grpc.Server, srv StringAppendServer) {
	s.RegisterService(&_StringAppend_serviceDesc, srv)
}

func _StringAppend_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(StringAppendServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _StringAppend_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(StringAppendServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _StringAppend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.StringAppend",
	HandlerType: (*StringAppendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _StringAppend_IO_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _StringAppend_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Download service

type DownloadClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
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

func (c *downloadClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.Download/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Download service

type DownloadServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
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

func _Download_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DownloadServer).Ping(ctx, in)
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
		{
			MethodName: "Ping",
			Handler:    _Download_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for S3 service

type S3Client interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
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

func (c *s3Client) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.S3/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S3 service

type S3Server interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
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

func _S3_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(S3Server).Ping(ctx, in)
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
		{
			MethodName: "Ping",
			Handler:    _S3_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for ImageResize service

type ImageResizeClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
}

type imageResizeClient struct {
	cc *grpc.ClientConn
}

func NewImageResizeClient(cc *grpc.ClientConn) ImageResizeClient {
	return &imageResizeClient{cc}
}

func (c *imageResizeClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.ImageResize/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageResizeClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.ImageResize/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageResize service

type ImageResizeServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
}

func RegisterImageResizeServer(s *grpc.Server, srv ImageResizeServer) {
	s.RegisterService(&_ImageResize_serviceDesc, srv)
}

func _ImageResize_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ImageResizeServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ImageResize_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ImageResizeServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ImageResize_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.ImageResize",
	HandlerType: (*ImageResizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _ImageResize_IO_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ImageResize_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for RenderWebsite service

type RenderWebsiteClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
}

type renderWebsiteClient struct {
	cc *grpc.ClientConn
}

func NewRenderWebsiteClient(cc *grpc.ClientConn) RenderWebsiteClient {
	return &renderWebsiteClient{cc}
}

func (c *renderWebsiteClient) IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error) {
	out := new(BinReply)
	err := grpc.Invoke(ctx, "/babl.RenderWebsite/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *renderWebsiteClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.RenderWebsite/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RenderWebsite service

type RenderWebsiteServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
}

func RegisterRenderWebsiteServer(s *grpc.Server, srv RenderWebsiteServer) {
	s.RegisterService(&_RenderWebsite_serviceDesc, srv)
}

func _RenderWebsite_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RenderWebsiteServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RenderWebsite_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RenderWebsiteServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RenderWebsite_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.RenderWebsite",
	HandlerType: (*RenderWebsiteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _RenderWebsite_IO_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _RenderWebsite_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for TestFail service

type TestFailClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
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

func (c *testFailClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/babl.TestFail/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestFail service

type TestFailServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
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

func _TestFail_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TestFailServer).Ping(ctx, in)
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
		{
			MethodName: "Ping",
			Handler:    _TestFail_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
