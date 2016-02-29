//go:generate go-bindata data/...

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/log"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/larskluge/babl/shared"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	app.Version = "0.1.2"
	app.Action = defaultAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "module, m",
			Usage:  "Module to serve",
			EnvVar: "BABL_MODULE",
		},
		cli.StringFlag{
			Name:   "cmd",
			Usage:  "Command to be executed",
			Value:  "cat",
			EnvVar: "BABL_COMMAND",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "Port for server to be started on",
			EnvVar: "PORT",
			Value:  4444,
		},
	}
	return
}

func defaultAction(c *cli.Context) {
	module := c.String("module")
	if module == "" {
		cli.ShowAppHelp(c)
		os.Exit(1)
	} else {
		shared.EnsureModuleExists(module)
		command = c.String("cmd")
		address := fmt.Sprintf(":%d", c.Int("port"))

		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Serving module %s, listening at %s..", module, address)

		certPEMBlock, _ := Asset("data/server.pem")
		keyPEMBlock, _ := Asset("data/server.key")
		cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		if err != nil {
			log.Fatalf("Could not load key pair %v", err)
		}
		creds := credentials.NewServerTLSFromCert(&cert)
		opts := []grpc.ServerOption{grpc.Creds(creds)}

		s := grpc.NewServer(opts...)
		pb.Modules[module].Server(s, &server{})
		s.Serve(lis)
	}
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	start := time.Now()

	cmd := exec.Command(command)
	env := os.Environ()
	cmd.Env = []string{} //{"FOO=BAR"}

	for k, v := range in.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = append(cmd.Env, env...)

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
	err := cmd.Start()
	if err != nil {
		log.Printf("cmd.Start: %v", err)
	}

	stdin.Write(in.Stdin)
	stdin.Close()
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Printf("ioutil.ReadAll[stdout]: %v", err)
	}
	errBytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Printf("ioutil.ReadAll[stderr]: %v", err)
	}

	res := pbm.BinReply{
		Stdout:   outBytes,
		Stderr:   errBytes,
		Exitcode: 0,
	}

	if err := cmd.Wait(); err != nil {
		res.Exitcode = 255
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				res.Exitcode = int32(status.ExitStatus())
				log.Printf("Exit Status: %d", status.ExitStatus())
			}
		} else {
			log.Printf("cmd.Wait: %v", err)
		}
	}

	status := 500
	if res.Exitcode == 0 {
		status = 200
	}

	elapsed := float64(time.Since(start).Seconds() * 1000)

	log.Printf("stdin=%d stdout=%d stderr=%d exitcode=%d status=%d duration_ms=%.3f", len(in.Stdin), len(res.Stdout), len(res.Stderr), res.Exitcode, status, elapsed)

	return &res, nil
}

func (s *server) Ping(ctx context.Context, in *pbm.Empty) (*pbm.Pong, error) {
	log.Println("Ping Request")
	res := pbm.Pong{Val: "pong"}
	return &res, nil
}
