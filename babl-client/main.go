package main

import (
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

const (
	address = "localhost:4444"
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

		run(module, env)
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

func run(module string, env map[string]string) {
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

	connection := pb.NewStringUpcaseClient(conn)

	req := pb.BinRequest{In: in}
	req.Env = env

	r, err := connection.IO(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	log.Printf("Response: %s", r.Out)
	if r.Status != pb.BinReply_SUCCESS {
		os.Exit(int(r.Status))
	}
}
