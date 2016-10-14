package bablutils

import (
	"github.com/kardianos/osext"
	"github.com/mattn/go-isatty"
	"io/ioutil"
	"os"
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

func AppPath() string {
	app, err := osext.Executable()
	check(err)
	return app
}
