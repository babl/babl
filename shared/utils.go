package shared

import (
	"fmt"
	"log"
	"os"
	"sort"

	pb "github.com/larskluge/babl/protobuf"
)

func EnsureModuleExists(module string) {
	if _, exists := pb.Modules[module]; exists == false {
		log.Printf("Unknown module '%s'", module)
		os.Exit(2)
	}
}

func PrintAvailableModules() {
	var modules []string
	for module, _ := range pb.Modules {
		modules = append(modules, module)
	}
	sort.Strings(modules)
	for _, module := range modules {
		fmt.Println(module)
	}
}
