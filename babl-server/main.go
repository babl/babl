package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

var command string

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Server for a Babl Module"
	app.Version = "0.0.1"
	app.Action = defaultAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "cmd",
			Usage: "Command to be executed",
			Value: "cat",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "Port for server to be started on",
			EnvVar: "PORT",
			Value:  4444,
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose logging",
		},
	}
	return
}

func defaultAction(c *cli.Context) {
	// verbose := c.Bool("verbose")
	command = c.String("cmd")
	port := fmt.Sprintf(":%d", c.Int("port"))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening at %s..", port)
	s := grpc.NewServer()
	pb.RegisterStringUpcaseServer(s, &server{})
	s.Serve(lis)
}

func (s *server) IO(ctx context.Context, in *pb.BinRequest) (*pb.BinReply, error) {
	log.Printf("Received: %s", in.In)

	log.Printf("Executing %s", command)
	cmd := exec.Command(command)
	cmd.Env = make([]string, len(in.Env)) //{"FOO=BAR"}

	for k, v := range in.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

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

	return &res, nil
}
