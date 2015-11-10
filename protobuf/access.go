package babl

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type BinaryClient interface {
	IO(ctx context.Context, in *BinRequest, opts ...grpc.CallOption) (*BinReply, error)
}
type BinaryServer interface {
	IO(context.Context, *BinRequest) (*BinReply, error)
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

var Modules = map[string]Component{
	"string-upcase": Component{Client: NewStringUpcaseClient2, Server: RegisterStringUpcaseServer2},
	"download":      Component{Client: NewDownloadClient2, Server: RegisterDownloadServer2},
}
