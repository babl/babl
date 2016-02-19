package main

import (
	"fmt"

	"github.com/larskluge/babl/shared"
)

func main() {
	m := shared.NewModule("string-append", false)
	m.Env = map[string]string{
		"APPENDIX": "bar",
	}
	stdout, _, _, _ := m.Call(shared.ReadStdin())

	m2 := shared.NewModule("string-upcase", false)
	stdout2, _, _, _ := m2.Call(stdout)

	fmt.Println(string(stdout2))
}
