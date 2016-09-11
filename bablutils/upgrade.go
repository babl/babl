package bablutils

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
)

type Upgrade struct {
	App string
}

func NewUpgrade(app string) Upgrade {
	return Upgrade{App: app}
}

func (u *Upgrade) Upgrade(currentVersion string) {
	if u.LatestVersionRunning(currentVersion) {
		fmt.Println("Already up-to-date.")
		return
	}

	mv, err := exec.LookPath("mv")
	check(err)

	tmpfile, err := ioutil.TempFile("", "upgrading-"+u.App)
	check(err)

	resp, err := http.Get(u.BinaryUrl())
	check(err)
	defer resp.Body.Close()

	decompress, err := gzip.NewReader(resp.Body)
	check(err)
	_, err = io.Copy(tmpfile, decompress)
	check(err)
	decompress.Close()
	tmpfile.Close()

	info, err := os.Stat(u.AppPath())
	check(err)
	os.Chmod(tmpfile.Name(), info.Mode())

	err = syscall.Exec(mv, []string{"mv", tmpfile.Name(), u.AppPath()}, os.Environ())
	check(err)
}

func (u *Upgrade) BinaryUrl() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/%s_%s_%s.gz", u.App, goos, goarch)
}

func (u *Upgrade) AppPath() string {
	app, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func (u *Upgrade) latestVersionUrl() string {
	return fmt.Sprintf("http://s3.amazonaws.com/babl/%s-latest-version.txt", u.App)
}

func (u *Upgrade) LatestVersion() string {
	res, err := http.Get(u.latestVersionUrl())
	check(err)
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	check(err)
	version := string(content)
	re := regexp.MustCompile("([0-9\\.]+)\n?$")
	matches := re.FindStringSubmatch(version)
	if len(matches) < 2 {
		msg := fmt.Sprintf("Version number not detected in '%s'", version)
		panic(msg)
	}
	return matches[1]
}

func (u *Upgrade) LatestVersionRunning(version string) bool {
	return u.LatestVersion() == version
}
