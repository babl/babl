package module

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type moduleConfig struct {
	Id string `yaml:"id" json:"id"`
}

func ModuleConfig() (cfg moduleConfig) {
	contents, err := ioutil.ReadFile("babl.yml")
	if err != nil {
		check(err)
	}
	type config struct {
		Id string `yaml:"id" json:"id"`
	}
	if err := yaml.Unmarshal(contents, &cfg); err != nil {
		check(err)
	}
	return
}
