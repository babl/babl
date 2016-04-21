package babl

import (
	"fmt"

	pb "github.com/larskluge/babl/protobuf/messages"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//
// CLIENT
//

type BinaryClient interface {
	IO(serviceName string, ctx context.Context, in *pb.BinRequest, opts ...grpc.CallOption) (*pb.BinReply, error)
	Ping(serviceName string, ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.Pong, error)
}

type binaryClient struct {
	cc *grpc.ClientConn
}

func NewBinaryClient(cc *grpc.ClientConn) BinaryClient {
	return &binaryClient{cc}
}

func servicePath(serviceName, method string) string {
	return fmt.Sprintf("/%s/%s", serviceName, method)
}

func (c *binaryClient) IO(serviceName string, ctx context.Context, in *pb.BinRequest, opts ...grpc.CallOption) (*pb.BinReply, error) {
	out := new(pb.BinReply)
	err := grpc.Invoke(ctx, servicePath(serviceName, "IO"), in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryClient) Ping(serviceName string, ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.Pong, error) {
	out := new(pb.Pong)
	err := grpc.Invoke(ctx, servicePath(serviceName, "Ping"), in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
// SERVER
//

type BinaryServer interface {
	IO(context.Context, *pb.BinRequest) (*pb.BinReply, error)
	Ping(context.Context, *pb.Empty) (*pb.Pong, error)
}

func RegisterBinaryServer(serviceName string, s *grpc.Server, srv BinaryServer) {
	desc := grpc.ServiceDesc{
		ServiceName: serviceName,
		HandlerType: (*BinaryServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "IO",
				Handler:    _Binary_IO_Handler,
			},
			{
				MethodName: "Ping",
				Handler:    _Binary_Ping_Handler,
			},
		},
		Streams: []grpc.StreamDesc{},
	}
	s.RegisterService(&desc, srv)
}

func _Binary_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BinaryServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Binary_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BinaryServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}
