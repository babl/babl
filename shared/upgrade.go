package shared

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/kardianos/osext"
	"github.com/larskluge/babl/log"
)

func binUrl(app string) string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	return fmt.Sprintf("http://s3.amazonaws.com/babl/%s_%s_%s.gz", app, goos, goarch)
}

func appPath() string {
	app, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func Upgrade(app string) {
	sh, err := exec.LookPath("sh")
	if err != nil {
		log.Fatal(err)
	}

	cmd := fmt.Sprintf("wget -O- '%s' | gunzip > '%s' && chmod +x '%s'", binUrl(app), appPath(), appPath())

	err = syscall.Exec(sh, []string{"sh", "-c", cmd}, os.Environ())
	log.Fatal(err)
}
