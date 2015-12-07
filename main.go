package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/shared"
	"github.com/mattn/go-isatty"
)

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func address(c *cli.Context) string {
	return fmt.Sprintf("%s:%d", c.GlobalString("host"), c.GlobalInt("port"))
}

func defaultAction(c *cli.Context, module_with_tag string) {
	tag := ""
	parts := strings.Split(module_with_tag, ":")
	module := parts[0]
	if len(parts) > 1 {
		tag = parts[1]
	}

	log.Println("Tag", tag)

	shared.EnsureModuleExists(module)
	log.Println("connecting to module", module)

	env := make(map[string]string)
	mod, ok := Config().Defaults[module_with_tag]
	if ok {
		env = mod.Env
	}
	buildEnv(&env, c.StringSlice("env"))
	log.Println("env", env)

	verbose := c.GlobalBool("verbose")
	log.Println("verbose", verbose)

	address := address(c)
	log.Printf("Connecting to %s..", address)

	m := shared.Module{
		Name:    module,
		Address: address,
		Env:     env,
	}
	_, _, exitcode, _ := m.Call(stdin())
	os.Exit(exitcode)
}

func buildEnv(env *map[string]string, envs []string) {
	for _, val := range envs {
		x := strings.Split(val, "=")
		(*env)[x[0]] = x[1]
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
