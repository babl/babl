package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"github.com/larskluge/babl/shared"
	"golang.org/x/net/context"
)

func appendModuleCommand(cmds *[]cli.Command, module string) {
	*cmds = append(*cmds, cli.Command{
		Name:  module,
		Usage: "MODULE",
		Action: func(c *cli.Context) {
			module := c.Command.Name
			defaultAction(c, module)
		},
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "env, e",
				Usage: "Send environment variables, e.g. -e FOO=42",
			},
		},
	})
}

func pingSubCommands() (cmds []cli.Command) {
	for _, module := range shared.Modules() {
		cmds = append(cmds, cli.Command{
			Name: module,
			Action: func(c *cli.Context) {
				module := c.Command.Name
				fmt.Print("ping.. ")
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
	return
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
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose logging",
		},
	}
	app.HideHelp = true

	app.Commands = []cli.Command{
		{
			Name:    "list-modules",
			Aliases: []string{"ls"},
			Usage:   "List all available modules",
			Action: func(_ *cli.Context) {
				shared.PrintAvailableModules()
			},
		},
		{
			Name:        "ping",
			Usage:       "ping <module>",
			Subcommands: pingSubCommands(),
			Action: func(_ *cli.Context) {
				fmt.Println("Unknown module")
				os.Exit(3)
			},
		},
		{
			Name:  "config",
			Usage: "Print configuration",
			Action: func(_ *cli.Context) {
				fmt.Println(Config())
			},
		},
	}

	for _, module := range shared.Modules() {
		appendModuleCommand(&app.Commands, module)
	}
	for module, _ := range Config().Defaults {
		appendModuleCommand(&app.Commands, module)
	}
	return
}
