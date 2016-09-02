package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"

	"github.com/DavidHuie/quartz/go/quartz"
	"github.com/larskluge/babl/bablmodule"
)

const Version = "0.1.1"

var (
	printVersion = flag.Bool("version", false, "print version & exit")
	endpoint     = flag.String("endpoint", "babl.sh:4444", "Endpoint to connect to")
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
	if bablmodule.CheckModuleName(req.Name) {
		m := bablmodule.New(req.Name)
		m.Env = req.Env
		m.Address = *endpoint

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
		fmt.Println(Version)
	} else {
		Babl := &Babl{}
		quartz.RegisterName("babl", Babl)
		quartz.Start()
	}
}
