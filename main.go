package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"github.com/larskluge/babl/shared"
	"github.com/mattn/go-isatty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Client to access the Babl Network."
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "Host to connect to, e.g. babl.sh, localhost",
			Value: "babl.sh",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "Port to connect to",
			EnvVar: "PORT",
			Value:  4444,
		},
		cli.StringSliceFlag{
			Name:  "env, e",
			Usage: "Send environment variables, e.g. -e FOO=42",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose logging",
		},
	}
	app.HideHelp = true

	pingCommands := []cli.Command{}
	for _, module := range shared.Modules() {
		pingCommands = append(pingCommands, cli.Command{
			Name: module,
			Action: func(c *cli.Context) {
				module := c.Command.Name
				fmt.Print("pinging.. ")
				conn := conn(address(c))
				defer conn.Close()
				connection := pb.Modules[module].Client(conn)
				req := pb.Empty{}
				res, err := connection.Ping(context.Background(), &req)
				if err == nil {
					fmt.Println(res.Val)
				} else {
					log.Fatalf("Failed: %v", err)
				}
			},
		})
	}

	app.Commands = []cli.Command{
		{
			Name:    "list-modules",
			Aliases: []string{"ls"},
			Usage:   "List all available modules",
			Action: func(_ *cli.Context) {
				shared.PrintAvailableModules()
			},
			SkipFlagParsing: true,
		},
		{
			Name:            "ping",
			Usage:           "ping <module>",
			SkipFlagParsing: true,
			Subcommands:     pingCommands,
		},
	}

	for _, module := range shared.Modules() {
		app.Commands = append(app.Commands, cli.Command{
			Name:  module,
			Usage: "MODULE",
			Action: func(c *cli.Context) {
				module := c.Command.Name
				defaultAction(c, module)
			},
		})
	}
	return
}

func address(c *cli.Context) string {
	return fmt.Sprintf("%s:%d", c.GlobalString("host"), c.GlobalInt("port"))
}

func defaultAction(c *cli.Context, module string) {
	shared.EnsureModuleExists(module)
	log.Println("connecting to module", module)

	env := buildEnv(c.GlobalStringSlice("env"))
	log.Println("env", env)

	verbose := c.GlobalBool("verbose")
	log.Println("verbose", verbose)

	address := address(c)
	log.Printf("Connecting to %s..", address)
	run(address, module, env)
}

func buildEnv(envs []string) (env map[string]string) {
	env = make(map[string]string)
	for _, val := range envs {
		x := strings.Split(val, "=")
		env[x[0]] = x[1]
	}
	return
}

func stdin() (in []byte) {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	log.Printf("%d bytes read from stdin", len(in))
	return
}

func conn(address string) *grpc.ClientConn {
	data, err := Asset("data/ca.pem")
	if err != nil {
		log.Fatal("asset not found")
	}
	sn := "babl.test.youtube.com"
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(data) {
		log.Fatal("credentials: failed to append certificates")
	}

	creds := credentials.NewClientTLSFromCert(cp, sn)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func run(address string, module string, env map[string]string) {
	conn := conn(address)
	defer conn.Close()

	connection := pb.Modules[module].Client(conn)
	req := pb.BinRequest{Stdin: stdin(), Env: env}
	res, err := connection.IO(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	status := "SUCCESS"
	if res.Exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:", status, len(res.Stdout), len(res.Stderr))
	log.Print(string(res.Stderr))
	fmt.Printf("%s", res.Stdout)
	os.Exit(int(res.Exitcode))
}
