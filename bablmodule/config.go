package bablmodule

import (
	"io/ioutil"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	ConfigFile = ".babl.yml"
)

type Cfg struct {
	Defaults map[string]module
}

type module struct {
	Module string
	Env    Env
}

func Config() (cfg Cfg) {
	filename := path.Join(os.Getenv("HOME"), ConfigFile)

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
	}
	return
}
