package shared

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/larskluge/babl/log"
	pb "github.com/larskluge/babl/protobuf"
	"github.com/mattn/go-isatty"
)

func ModuleExists(module string) bool {
	_, exists := pb.Modules[module]
	return exists
}

func EnsureModuleExists(module string) {
	if !ModuleExists(module) {
		log.Printf("Unknown module '%s'", module)
		os.Exit(2)
	}
}

func PrintAvailableModules() {
	for _, module := range Modules() {
		fmt.Println(module)
	}
	for module, _ := range Config().Defaults {
		fmt.Println(module)
	}
}

func Modules() (modules []string) {
	for module, _ := range pb.Modules {
		modules = append(modules, module)
	}
	sort.Strings(modules)
	return
}

func Version() string {
	version, err := Asset("data/VERSION")
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(version), "\n")
}

func ReadStdin() (in []byte) {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	log.Printf("%d bytes read from stdin", len(in))
	return
}
