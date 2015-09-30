package main

import (
	// "fmt"
	"log"
	"net"

	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":4444"
)

// server is used to implement hellowrld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %s", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) IO(ctx context.Context, in *pb.BinRequest) (*pb.BinReply, error) {
	log.Printf("Received: %s", in.In)
	// msg := fmt.Sprintf("Hello %s", in.In)
	return &pb.BinReply{Out: in.In}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening at %s..", port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterStringUpcaseServer(s, &server{})
	s.Serve(lis)
}
