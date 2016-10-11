package bablutils

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

type Update struct {
	App     string
	Version string
}

func NewUpdate(app string, version string) Update {
	return Update{App: app, Version: version}
}

func GetDesiredVersion(app string, version string, args []string) {
	bv := os.Getenv("BABL_DESIRED_SERVER_VERSION")
	if bv == "" || bv == version {
		return
	} else {
		u := NewUpdate(app, bv)
		u.Update(args)
	}
	return
}

func (u *Update) BinaryUrl() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/releases/%s/%s_%s_%s.gz", u.Version, u.App, goos, goarch)
}

func (u *Update) downloadVersionBinary(file string) bool {
	out, err := exec.Command(file, "-plainversion").Output()
	check(err)
	return u.Version == strings.TrimSpace(string(out))
}

func (u *Update) Update(args []string) {

	mv, err := exec.LookPath("mv")
	check(err)

	tmpfile, err := ioutil.TempFile("", "updating-"+u.App)
	check(err)

	resp, err := http.Get(u.BinaryUrl())
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if http.DetectContentType(body) != "application/x-gzip" {
		panic(fmt.Sprintf("File does not exist %s /n Are you sure %s is the right version?", u.BinaryUrl(), u.Version))
	}

	decompress, err := gzip.NewReader(resp.Body)
	check(err)
	_, err = io.Copy(tmpfile, decompress)
	check(err)
	decompress.Close()
	tmpfile.Close()

	info, err := os.Stat(AppPath())
	check(err)
	os.Chmod(tmpfile.Name(), info.Mode())
	if downloadVersionBinary(tmpfile.Name()) {
		cmd := exec.Command(mv, []string{tmpfile.Name(), AppPath()}...)
		err = cmd.Run()
		check(err)

		err = syscall.Exec(args[0], args, os.Environ())
		check(err)
	} else {
		panic("")
	}

}
