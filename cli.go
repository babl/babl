package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/log"
	"github.com/larskluge/babl/shared"
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
				fmt.Print("ping.. ")
				m := shared.NewModule(c.Command.Name)
				m.Address = address(c)
				res, err := m.Ping()
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
	app.Version = shared.Version()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "Host to connect to, e.g. babl.sh, localhost",
			Value: "babl.sh",
		},
		cli.IntFlag{
			Name:  "port",
			Usage: "Port to connect to",
			Value: 4444,
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug mode & verbose logging",
		},
	}
	app.Action = func(c *cli.Context) {
		fmt.Fprintln(app.Writer, "Incorrect Usage.")
		fmt.Fprintln(app.Writer)
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

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
				fmt.Println(shared.Config())
			},
		},
	}

	for _, module := range shared.Modules() {
		appendModuleCommand(&app.Commands, module)
	}
	for module, _ := range shared.Config().Defaults {
		appendModuleCommand(&app.Commands, module)
	}
	return
}
