package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"

	"github.com/DavidHuie/quartz/go/quartz"
	"github.com/larskluge/babl/shared"
)

var (
	printVersion = flag.Bool("version", false, "print version & exit")
)

type Babl struct{}

type ModuleRequest struct {
	Name  string
	Env   map[string]string
	Stdin string
}

type ModuleResponse struct {
	Stdout   string
	Stderr   string
	Exitcode int
}

func (_ *Babl) Module(req ModuleRequest, response *ModuleResponse) error {
	if shared.CheckModuleName(req.Name) {
		m := shared.NewModule(req.Name, false)
		m.Env = req.Env

		stdin, err := base64.StdEncoding.DecodeString(req.Stdin)
		if err != nil {
			return err
		}

		stdout, stderr, exitcode, err := m.Call(stdin)

		response.Stdout = base64.StdEncoding.EncodeToString(stdout)
		response.Stderr = base64.StdEncoding.EncodeToString(stderr)
		response.Exitcode = exitcode

		return err
	} else {
		return errors.New("babl-rpc: module name format incorrect")
	}
}

func main() {
	flag.Parse()
	if *printVersion {
		fmt.Println(shared.Version())
	} else {
		Babl := &Babl{}
		quartz.RegisterName("babl", Babl)
		quartz.Start()
	}
}
