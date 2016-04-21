package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/log"
	"github.com/larskluge/babl/shared"
)

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
			Name:   "debug",
			Usage:  "Enable debug mode & verbose logging",
			EnvVar: "BABL_DEBUG",
		},
	}
	app.Action = func(c *cli.Context) {
		module := c.Args().First()
		if shared.CheckModuleName(module) {
			defaultAction(c)
		} else {
			fmt.Fprintln(app.Writer, "Incorrect Usage.")
			fmt.Fprintln(app.Writer)
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
	}

	app.Commands = []cli.Command{
		{
			Name:    "list-modules",
			Aliases: []string{"ls"},
			Usage:   "List all available modules",
			Action: func(c *cli.Context) {
				shared.PrintAvailableModules(c.Bool("defaults"))
			},
			Flags: []cli.Flag{
				cli.BoolTFlag{
					Name:  "defaults",
					Usage: "Lists all modules & their configured defaults",
				},
			},
		},
		{
			Name:  "ping",
			Usage: "ping <module>",
			Action: func(c *cli.Context) {
				module := c.Args().First()
				fmt.Print("ping.. ")
				m := shared.NewModule(module, false)
				m.Address = address(c)
				res, err := m.Ping()
				if err == nil {
					fmt.Println(res.Val)
				} else {
					log.Fatalf("Failed: %v", err)
				}
			},
		},
		{
			Name:  "logs",
			Usage: "logs <module>",
			Action: func(c *cli.Context) {
				pattern := c.Args().First()
				LogsInit()
				Logs(pattern)
			},
		},
		{
			Name:  "config",
			Usage: "Print configuration",
			Action: func(_ *cli.Context) {
				fmt.Println(shared.Config())
			},
		},
		{
			Name:  "upgrade",
			Usage: "Upgrades the client to the latest available version",
			Action: func(_ *cli.Context) {
				shared.Upgrade("babl")
			},
		},
	}
	return
}
