//go:generate sh -c "cd bablmodule && go-bindata -pkg bablmodule -prefix data data/..."
//go:generate sh -c "cd protobuf && protoc -I ./messages/ ./messages/main.proto --go_out=plugins=grpc:messages"

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl/bablmodule"
	"github.com/larskluge/babl/bablutils"
)

const Version = "0.4.4"

func main() {
	bablutils.PrintPlainVersionAndExit(os.Args, Version)
	app := configureCli()
	app.Run(os.Args)
}

func defaultAction(module_with_tag string, envs []string, address string, async, debug bool) {
	m := bablmodule.New(module_with_tag)
	m.SetAsync(async)
	m.SetDebug(debug)

	log.SetOutput(os.Stderr)
	log.SetLevel(log.ErrorLevel)

	if !debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Infof("Connecting to module", m.Name, m.Tag)

	applyEnv(&m.Env, envs)
	log.Debugf("%+v\n", m.Env)

	m.Address = address

	in := bablutils.ReadStdin()
	log.Infof("%d bytes read from stdin", len(in))
	stdout, stderr, exitcode, err := m.Call(in)
	status := "SUCCESS"
	if err != nil || exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:, exit w/ %d", status, len(stdout), len(stderr), exitcode)
	if err != nil {
		log.Errorf("%+v\n", err)
	}
	if len(stderr) > 0 {
		log.Error(string(stderr))
	}
	fmt.Print(string(stdout))
	os.Exit(exitcode)
}

func applyEnv(env *bablmodule.Env, envs []string) {
	for _, val := range envs {
		x := strings.SplitN(val, "=", 2)
		(*env)[x[0]] = x[1]
	}
}
