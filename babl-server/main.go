package main

import (
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
	// cmd := exec.Command("exit", "1")
	cmd := exec.Command("bash", "-c", "echo error >&2")

	stdin, errIn := cmd.StdinPipe()
	if errIn != nil {
		log.Printf("cmd.StdinPipe: %v", errIn)
	}
	stdout, errOut := cmd.StdoutPipe()
	if errOut != nil {
		log.Printf("cmd.StdoutPipe: %v", errOut)
	}
	stderr, errErr := cmd.StderrPipe()
	if errErr != nil {
		log.Printf("cmd.StderrPipe: %v", errErr)
	}
	cmd.Start()

	stdin.Write(in.In)
	stdin.Close()
	grepBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
	}
	errBytes, _ := ioutil.ReadAll(stderr)
	if len(errBytes) > 0 {
		log.Printf("stderr: %s", errBytes)
	}

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
