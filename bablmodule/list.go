package bablmodule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckModuleName(module string) bool {
	r := regexp.MustCompile("^[a-z][a-z0-9-]*/[a-z][a-z0-9-:]*$")
	return r.MatchString(module)
}

func PrintAvailableModules(printDefaults bool) {
	for _, module := range Modules() {
		fmt.Println(module)
	}
	if printDefaults {
		for module, _ := range Config().Defaults {
			fmt.Println(module)
		}
	}
}

func Modules() (modules []string) {
	resp, err := http.Get("https://babl.sh/api/modules")
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	var mods []map[string]interface{}
	err = json.Unmarshal(body, &mods)
	check(err)
	for _, mod := range mods {
		if mod["type"] == "babl" {
			modules = append(modules, mod["full_name"].(string))
		}
	}
	sort.Strings(modules)
	return
}
