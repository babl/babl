//go:generate protoc -I ./protobuf/ ./protobuf/babl.proto --go_out=plugins=grpc:protobuf
//go:generate cd shared && go-bindata -pkg shared data/...

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/larskluge/babl/shared"
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

	stdout, stderr, exitcode, err := m.Call(shared.ReadStdin())
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
		x := strings.Split(val, "=")
		(*env)[x[0]] = x[1]
	}
}
