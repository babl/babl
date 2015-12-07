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
	m := shared.NewModule(module_with_tag)

	log.Println("Connecting to module", m.Name, m.Tag)

	buildEnv(&m.Env, c.StringSlice("env"))
	log.Println("env", m.Env)

	// verbose := c.GlobalBool("verbose")
	// log.Println("verbose", verbose)

	m.Address = address(c)
	log.Printf("Connecting to %s..", m.Address)

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
