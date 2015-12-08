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

	applyEnv(&m.Env, c.StringSlice("env"))
	log.Println("env", m.Env)

	// verbose := c.GlobalBool("verbose")
	// log.Println("verbose", verbose)

	m.Address = address(c)
	log.Printf("Connecting to %s..", m.Address)

	stdout, stderr, exitcode, err := m.Call(stdin())
	status := "SUCCESS"
	if err != nil || exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:", status, len(stdout), len(stderr))
	if len(stderr) > 0 {
		log.Print(string(stderr))
	}
	fmt.Printf("%s", stdout)
	os.Exit(exitcode)
}

func applyEnv(env *map[string]string, envs []string) {
	for _, val := range envs {
		x := strings.Split(val, "=")
		(*env)[x[0]] = x[1]
	}
}

func stdin() (in []byte) {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	log.Printf("%d bytes read from stdin", len(in))
	return
}
