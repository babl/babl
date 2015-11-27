package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	"gopkg.in/yaml.v2"
)

const (
	ConfigFile = ".babl.yml"
)

type Cfg struct {
	Defaults map[string]Module
}

type Module struct {
	Module string
	Env    map[string]string
}

func Config() (cfg *Cfg) {
	usr, _ := user.Current()
	filename := path.Join(usr.HomeDir, ConfigFile)

	if _, err := os.Stat(filename); err == nil {
		// config found

		if contents, err := ioutil.ReadFile(filename); err == nil {
			err := yaml.Unmarshal(contents, &cfg)
			if err != nil {
				log.Printf("Please check your configuration file (%s): %v", ConfigFile, err)
			}
		} else {
			log.Printf("Could not read configuration file (%s): %v", ConfigFile, err)
		}
		return
	} else {
		return nil
	}
}
