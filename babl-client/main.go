package main

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:4444"
	// address     = "localdocker:8080"
	defaultName = "world"
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

		envs := c.StringSlice("env")
		log.Println("env", len(envs), envs)

		verbose := c.Bool("verbose")
		log.Println("verbose", verbose)

		run(c)
	}
}

func run(cli *cli.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// module := cli.Args().First()
	connection := pb.NewStringUpcaseClient(conn)

	in := []byte("hello there")
	req := pb.BinRequest{In: in}
	req.Env = make(map[string]string)
	envs := cli.StringSlice("env")
	for _, val := range envs {
		x := strings.Split(val, "=")
		req.Env[x[0]] = x[1]
	}

	r, err := connection.IO(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	log.Printf("Response: %s", r.Out)
}
