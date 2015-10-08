package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"github.com/mattn/go-isatty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Access the Babl Network."
	app.Version = "0.0.1"
	app.Action = defaultAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "Host to connect to, e.g. babl.sh, localhost",
			Value: "localhost", // "babl.sh",
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
	return
}

func defaultAction(c *cli.Context) {
	module := c.Args().First()
	if module == "" {
		cli.ShowAppHelp(c)
		os.Exit(1)
	} else {
		log.Println("connecting to module", module)

		env := buildEnv(c.StringSlice("env"))
		log.Println("env", env)

		verbose := c.Bool("verbose")
		log.Println("verbose", verbose)

		address := fmt.Sprintf("%s:%d", c.String("host"), c.Int("port"))
		log.Printf("Connecting to %s..", address)
		run(address, module, env)
	}
}

func buildEnv(envs []string) (env map[string]string) {
	env = make(map[string]string)
	for _, val := range envs {
		x := strings.Split(val, "=")
		env[x[0]] = x[1]
	}
	return
}

func run(address string, module string, env map[string]string) {
	var in []byte
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	log.Printf("%d bytes read from stdin", len(in))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	connection := pb.Modules[module].Client(conn)
	req := pb.BinRequest{In: in, Env: env}
	res, err := connection.IO(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	log.Printf("Response: %s", res.Out)
	if res.Status != pb.BinReply_SUCCESS {
		os.Exit(int(res.Status))
	}
}
