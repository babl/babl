package bablutils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"syscall"

	"github.com/kardianos/osext"
	"github.com/larskluge/babl/log"
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
	sh, err := exec.LookPath("sh")
	if err != nil {
		log.Fatal(err)
	}

	cmd := fmt.Sprintf("wget -q -O- '%s' | gunzip > '%s' && chmod +x '%s'", u.binUrl(), u.appPath(), u.appPath())

	err = syscall.Exec(sh, []string{"sh", "-c", cmd}, os.Environ())
	log.Fatal(err)
}

func (u *Upgrade) binUrl() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/%s_%s_%s.gz", u.App, goos, goarch)
}

func (u *Upgrade) appPath() string {
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
