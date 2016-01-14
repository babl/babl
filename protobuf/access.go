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
