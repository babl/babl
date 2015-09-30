package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":4444"
)

// server is used to implement hellowrld.GreeterServer.
type server struct{}

func (s *server) IO(ctx context.Context, in *pb.BinRequest) (*pb.BinReply, error) {
	log.Printf("Received: %s", in.In)

	grepCmd := exec.Command("cat")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write(in.In)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// msg := fmt.Sprintf("Hello %s", in.In)
	return &pb.BinReply{Out: grepBytes}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening at %s..", port)
	s := grpc.NewServer()
	pb.RegisterStringUpcaseServer(s, &server{})
	s.Serve(lis)
}
