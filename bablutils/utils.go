package bablutils

import (
	"io/ioutil"
	"os"

	"github.com/mattn/go-isatty"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadStdin() (in []byte) {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	return
}
