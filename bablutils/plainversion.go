package bablutils

import (
	"fmt"
	"os"
)

func PrintPlainVersionAndExit(args []string, version string) {
	for _, v := range args {
		if v == "-plainversion" {
			fmt.Println(version)
			os.Exit(0)
		}
	}
}
