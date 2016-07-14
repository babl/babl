//go:generate sh -c "cd shared && go-bindata -pkg shared data/..."

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/larskluge/babl/log"
	"github.com/larskluge/babl/shared"
)

const Version = "0.3.3"

func main() {
	app := configureCli()
	app.Run(os.Args)
}

func defaultAction(module_with_tag string, envs []string, address string, async, debug bool) {
	m := shared.NewModule(module_with_tag)
	m.SetAsync(async)
	m.SetDebug(debug)

	if !debug {
		log.SetOutput(ioutil.Discard)
	}

	log.Println("Connecting to module", m.Name, m.Tag)

	applyEnv(&m.Env, envs)
	log.Printf("%+v\n", m.Env)

	m.Address = address
	log.Printf("Connecting to %s..", m.Address)

	in := shared.ReadStdin()
	log.Printf("%d bytes read from stdin", len(in))
	stdout, stderr, exitcode, err := m.Call(in)
	status := "SUCCESS"
	if err != nil || exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:, exit w/ %d", status, len(stdout), len(stderr), exitcode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	if len(stderr) > 0 {
		log.Print(string(stderr))
	}
	fmt.Printf("%s", stdout)
	os.Exit(exitcode)
}

func applyEnv(env *shared.Env, envs []string) {
	for _, val := range envs {
		x := strings.SplitN(val, "=", 2)
		(*env)[x[0]] = x[1]
	}
}
