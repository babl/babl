package babl

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type BinaryClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
}
type BinaryServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
	Ping(context.Context, *Empty) (*Pong, error)
}

type ClientFunc func(cc *grpc.ClientConn) BinaryClient
type ServerFunc func(s *grpc.Server, srv BinaryServer)

type Component struct {
	Client ClientFunc
	Server ServerFunc
}

func NewStringUpcaseClient2(cc *grpc.ClientConn) BinaryClient {
	return &stringUpcaseClient{cc}
}
func RegisterStringUpcaseServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_StringUpcase_serviceDesc, srv)
}

func NewDownloadClient2(cc *grpc.ClientConn) BinaryClient {
	return &downloadClient{cc}
}
func RegisterDownloadServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_Download_serviceDesc, srv)
}

func NewS3Client2(cc *grpc.ClientConn) BinaryClient {
	return &s3Client{cc}
}
func RegisterS3Server2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_S3_serviceDesc, srv)
}

func NewTestFailClient2(cc *grpc.ClientConn) BinaryClient {
	return &testFailClient{cc}
}
func RegisterTestFailServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_TestFail_serviceDesc, srv)
}

var Modules = map[string]Component{
	"string-upcase": Component{Client: NewStringUpcaseClient2, Server: RegisterStringUpcaseServer2},
	"download":      Component{Client: NewDownloadClient2, Server: RegisterDownloadServer2},
	"s3":            Component{Client: NewS3Client2, Server: RegisterS3Server2},
	"test-fail":     Component{Client: NewTestFailClient2, Server: RegisterTestFailServer2},
}
