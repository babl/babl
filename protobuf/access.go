package babl

import (
	pb "github.com/larskluge/babl/protobuf/messages"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type BinaryClient interface {
	IO(ctx context.Context, in *pb.BinRequest, opts ...grpc.CallOption) (*pb.BinReply, error)
	Ping(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.Pong, error)
}
type BinaryServer interface {
	IO(context.Context, *pb.BinRequest) (*pb.BinReply, error)
	Ping(context.Context, *pb.Empty) (*pb.Pong, error)
}

type ClientFunc func(cc *grpc.ClientConn) BinaryClient
type ServerFunc func(s *grpc.Server, srv BinaryServer)

type Component struct {
	Client ClientFunc
	Server ServerFunc
}
