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
	"strings"
	"syscall"
)

type Upgrade struct {
	App  string
	Args []string
	url  string
}

func NewUpgrade(app string, args []string) Upgrade {
	return Upgrade{App: app, Args: args}
}

func (u *Upgrade) Upgrade(currentVersion string) {
	if u.LatestVersionRunning(currentVersion) {
		fmt.Println("Already up-to-date.")
	} else {
		u.url = u.LatestBinaryUrl()
		u.Update(currentVersion, u.LatestVersion())
	}

}

func (u *Upgrade) Update(currentVersion string, newVersion string) {
	if newVersion == "" || currentVersion == newVersion {
		return
	}

	mv, err := exec.LookPath("mv")
	check(err)

	tmpfile, err := ioutil.TempFile("", "upgrading-"+u.App)
	check(err)
	if u.url == "" {
		u.url = u.ReleaseBinaryUrl(newVersion)
	}
	resp, err := http.Get(u.url)
	check(err)
	defer resp.Body.Close()

	decompress, err := gzip.NewReader(resp.Body)
	check(err)
	_, err = io.Copy(tmpfile, decompress)
	check(err)
	decompress.Close()
	tmpfile.Close()

	info, err := os.Stat(AppPath())
	check(err)
	os.Chmod(tmpfile.Name(), info.Mode())

	if u.NewBinaryValid(tmpfile.Name(), newVersion) {
		cmd := exec.Command(mv, []string{tmpfile.Name(), AppPath()}...)
		err = cmd.Run()
		check(err)

		if len(u.Args) != 0 {
			err = syscall.Exec(u.Args[0], u.Args, os.Environ())
			check(err)
		}
		return
	} else {
		fmt.Println("Downloaded binary seems to be corrupt or wrong version, please try again later.")
		os.Exit(-1)
	}
}

func (u *Upgrade) LatestBinaryUrl() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/%s_%s_%s.gz", u.App, goos, goarch)
}

func (u *Upgrade) ReleaseBinaryUrl(version string) string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/releases/%s/%s_%s_%s.gz", version, u.App, goos, goarch)
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

func (u *Upgrade) NewBinaryValid(file string, upgradeVersion string) bool {
	out, err := exec.Command(file, "-plainversion").Output()
	check(err)
	return upgradeVersion == strings.TrimSpace(string(out))
}
