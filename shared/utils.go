package shared

import (
	"fmt"
	"log"
	"os"

	pb "github.com/larskluge/babl/protobuf"
)

func EnsureModuleExists(module string) {
	if _, exists := pb.Modules[module]; exists == false {
		log.Printf("Unknown module '%s'", module)
		os.Exit(2)
	}
}

func PrintAvailableModules() {
	for k, _ := range pb.Modules {
		fmt.Println(k)
	}
}
