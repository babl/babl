package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/bablmodule"
	"github.com/larskluge/babl/bablutils"
)

type envFlags []string

func (e *envFlags) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlags) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func parseEnvFlags(flags []string) (envs envFlags) {
	set := flag.NewFlagSet("env", flag.ExitOnError)
	set.Var(&envs, "env", "Send environment variables, e.g. -e FOO=42")
	set.Var(&envs, "e", "Send environment variables, e.g. -e FOO=42")
	set.Parse(flags)
	return
}

func address(c *cli.Context) string {
	return c.GlobalString("connect-to")
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Client to access the Babl Network."
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "connect-to, c",
			Usage:  "Host & port to connect to, e.g. babl.sh:4444, localhost:4445",
			Value:  "babl.sh:4444",
			EnvVar: "BABL_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "storage",
			Usage:  "Endpoint for Babl storage",
			Value:  "babl.sh:4443",
			EnvVar: "BABL_STORAGE",
		},
		cli.BoolFlag{
			Name:   "async",
			Usage:  "Flag request to be processed asynchronously and do not wait for a response",
			EnvVar: "BABL_ASYNC",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Enable debug mode & verbose logging",
			EnvVar: "BABL_DEBUG",
		},
	}
	app.Action = func(c *cli.Context) {
		mod := c.Args().First()
		if bablmodule.CheckModuleName(mod) {
			envs := parseEnvFlags(c.Args().Tail())
			async := c.GlobalBool("async")
			debug := c.GlobalBool("debug")
			storageEndpoint := c.GlobalString("storage")
			defaultAction(mod, envs, address(c), storageEndpoint, async, debug)
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
				bablmodule.PrintAvailableModules(c.Bool("defaults"))
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
				mod := c.Args().First()
				fmt.Print("ping.. ")
				m := bablmodule.New(mod)
				m.SetEndpoint(address(c))
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
			Name:  "home",
			Usage: "Open the module page in your browser",
			Action: func(c *cli.Context) {
				cfg := bablmodule.ModuleConfig()
				url := fmt.Sprintf("https://babl.sh/%s", cfg.Id)
				_, err := exec.Command("open", url).Output()
				if err == nil {
					os.Exit(0)
				} else {
					os.Exit(1)
				}
			},
		},
		{
			Name:  "info",
			Usage: "Prints runtime information for the current module",
			Action: func(c *cli.Context) {
				cfg := bablmodule.ModuleConfig()
				id := cfg.Id
				m := bablmodule.New("babl/runtime-info")
				m.Env = bablmodule.Env{"MODULE": id}
				stdout, stderr, exitcode, _, err := m.Call([]byte{})
				if err != nil {
					panic(err)
				}
				if exitcode == 0 {
					fmt.Print(string(stdout))
				} else {
					fmt.Print(string(stderr))
				}
				os.Exit(exitcode)
			},
		},
		{
			Name:  "config",
			Usage: "Print configuration",
			Action: func(_ *cli.Context) {
				fmt.Println(bablmodule.Config())
			},
		},
		{
			Name:  "upgrade",
			Usage: "Upgrades the client to the latest available version",
			Action: func(_ *cli.Context) {
				m := bablutils.NewUpgrade("babl")
				m.Upgrade(Version)
			},
		},
	}
	return
}
