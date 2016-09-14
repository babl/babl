package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/DavidHuie/quartz/go/quartz"
	"github.com/larskluge/babl/bablmodule"
)

const (
	Version             = "0.2.1"
	DefaultBablEndpoint = "babl.sh:4444" // lock in Babl v4 for the time being
)

var (
	printVersion = flag.Bool("version", false, "print version & exit")
)

type Babl struct{}

type ModuleRequest struct {
	Name            string
	Env             map[string]string
	Stdin           string
	PayloadUrl      string
	BablEndpoint    string
	StorageEndpoint string
}

type ModuleResponse struct {
	Stdout     string
	Stderr     string
	Exitcode   int
	PayloadUrl string
}

func (req *ModuleRequest) bablEndpoint() (be string) {
	be = req.BablEndpoint
	if be != "" {
		return
	}
	be = os.Getenv("BABL_ENDPOINT")
	if be != "" {
		return
	}
	return DefaultBablEndpoint
}

func (req *ModuleRequest) storageEndpoint() (se string) {
	se = req.StorageEndpoint
	if se != "" {
		return
	}
	se = os.Getenv("BABL_STORAGE")
	return
}

func (_ *Babl) Module(req ModuleRequest, response *ModuleResponse) error {
	if bablmodule.CheckModuleName(req.Name) {
		m := bablmodule.New(req.Name)
		m.Env = req.Env
		m.SetEndpoint(req.bablEndpoint())
		m.SetStorageEndpoint(req.storageEndpoint())
		m.FetchPayload = false
		m.PayloadUrl = req.PayloadUrl

		stdin, err := base64.StdEncoding.DecodeString(req.Stdin)
		if err != nil {
			return err
		}

		stdout, stderr, exitcode, payloadUrl, err := m.Call(stdin)
		if err != nil {
			return err
		}

		response.Stdout = base64.StdEncoding.EncodeToString(stdout)
		response.Stderr = base64.StdEncoding.EncodeToString(stderr)
		response.Exitcode = exitcode
		response.PayloadUrl = payloadUrl
		return nil
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
