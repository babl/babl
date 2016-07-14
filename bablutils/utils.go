package bablutils

import (
	"io/ioutil"
	"os"

	"github.com/mattn/go-isatty"
)

func ReadStdin() (in []byte) {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	return
}
