//go:generate go-bindata data/...

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
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
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})

	app := configureCli()
	app.Run(os.Args)
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Server for a Babl Module"
	app.Version = "0.3.0"
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
		if !shared.CheckModuleName(module) {
			log.WithFields(log.Fields{"module": module}).Fatal("Module name format incorrect")
		}
		command = c.String("cmd")
		address := fmt.Sprintf(":%d", c.Int("port"))

		log.Warn("Start module")

		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "address": address}).Fatal("Failed to listen at port")
		}

		certPEMBlock, _ := Asset("data/server.pem")
		keyPEMBlock, _ := Asset("data/server.key")
		cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		if err != nil {
			panic(err)
		}
		creds := credentials.NewServerTLSFromCert(&cert)
		opts := []grpc.ServerOption{grpc.Creds(creds)}

		s := grpc.NewServer(opts...)
		m := shared.NewModule(module, false)
		pb.RegisterBinaryServer(m.GrpcServiceName(), s, &server{})
		s.Serve(lis)
	}
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	start := time.Now()

	cmd := exec.Command(command)
	env := os.Environ()
	cmd.Env = []string{} //{"FOO=BAR"}

	vars := []string{}
	for k, v := range in.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		vars = append(vars, k)
	}
	cmd.Env = append(cmd.Env, env...)
	cmd.Env = append(cmd.Env, "BABL_VARS="+strings.Join(vars, ","))

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("cmd.StdinPipe")
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("cmd.StdoutPipe")
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("cmd.StderrPipe")
	}
	err = cmd.Start()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("cmd.Start")
	}

	stdin.Write(in.Stdin)
	stdin.Close()
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("ioutil.ReadAll[stdout]")
	}
	errBytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("ioutil.ReadAll[stderr]")
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
			}
		} else {
			log.WithFields(log.Fields{"error": err}).Error("cmd.Wait")
		}
	}

	status := 500
	if res.Exitcode == 0 {
		status = 200
	}

	elapsed := float64(time.Since(start).Seconds() * 1000)

	fields := log.Fields{
		"stdin":       len(in.Stdin),
		"stdout":      len(res.Stdout),
		"stderr":      len(res.Stderr),
		"exitcode":    res.Exitcode,
		"status":      status,
		"duration_ms": elapsed,
	}
	if status != 200 {
		fields["error"] = string(res.Stderr)
	}
	l := log.WithFields(fields)
	if status == 200 {
		l.Info("call")
	} else {
		l.Error("call")
	}

	return &res, nil
}

func (s *server) Ping(ctx context.Context, in *pbm.Empty) (*pbm.Pong, error) {
	log.Info("ping")
	res := pbm.Pong{Val: "pong"}
	return &res, nil
}
