package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"syscall"

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

	// cmd := exec.Command("cat")
	cmd := exec.Command("exit", "1")
	// cmd := exec.Command("bash", "-c", "echo error >&2")

	grepIn, errIn := cmd.StdinPipe()
	if errIn != nil {
		// 	panic(errIn)
	}
	grepOut, errOut := cmd.StdoutPipe()
	if errOut != nil {
		// 	panic(errOut)
	}
	cmd.Start()

	grepIn.Write(in.In)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	// if err != nil {
	// 	panic(err)
	// }

	res := pb.BinReply{Out: grepBytes}
	res.Status = pb.BinReply_SUCCESS

	if err := cmd.Wait(); err != nil {
		res.Status = pb.BinReply_ERROR
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
			}
		} else {
			log.Printf("cmd.Wait: %v", err)
		}
	}

	// msg := fmt.Sprintf("Hello %s", in.In)
	return &res, nil
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
