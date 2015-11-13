package shared

import (
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
