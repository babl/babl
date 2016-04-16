//go:generate bin/module-proto
//go:generate bin/protoc
//go:generate sh -c "bin/module-mapping > protobuf/access_mapping.go"
//go:generate sh -c "cd shared && go-bindata -pkg shared data/..."

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/log"
	"github.com/larskluge/babl/shared"
)

var debug = false

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func address(c *cli.Context) string {
	return fmt.Sprintf("%s:%d", c.GlobalString("host"), c.GlobalInt("port"))
}

func defaultAction(c *cli.Context, module_with_tag string) {
	debug = c.GlobalBool("debug")
	m := shared.NewModule(module_with_tag, debug)

	if !debug {
		log.SetOutput(ioutil.Discard)
	}

	log.Println("Connecting to module", m.Name, m.Tag)

	applyEnv(&m.Env, c.StringSlice("env"))
	log.Println("env", m.Env)

	m.Address = address(c)
	log.Printf("Connecting to %s..", m.Address)

	in := shared.ReadStdin()
	log.Printf("%d bytes read from stdin", len(in))
	stdout, stderr, exitcode, err := m.Call(in)
	status := "SUCCESS"
	if err != nil || exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:, exit w/ %d", status, len(stdout), len(stderr), exitcode)
	if len(stderr) > 0 {
		log.Print(string(stderr))
	}
	fmt.Printf("%s", stdout)
	os.Exit(exitcode)
}

func applyEnv(env *map[string]string, envs []string) {
	for _, val := range envs {
		x := strings.SplitN(val, "=", 2)
		(*env)[x[0]] = x[1]
	}
}
