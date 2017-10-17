package dep

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const updaterName = "dep"

type depConfig struct{}

type depUpdater struct {
	config *depConfig
}

func init() {
	runner.RegisterUpdater(updaterName, newDepUpdater)
}

func newDepUpdater(c *k.ModuleConfig) runner.Runner {
	cfg := convertUpdaterConfig(c)
	u := &depUpdater{config: cfg}
	return u
}

func (u *depUpdater) Name() string {
	return updaterName
}

func (u *depUpdater) ValidateConfig() error {
	if u.config == nil {
		return fmt.Errorf("no config should error")
	}
	return nil
}

func (u *depUpdater) Run(re k.RunEnv) error {
	vendor := filepath.Join(re.CurrentWorkDir(), "vendor")
	if _, err := os.Stat(vendor); err == nil {
		os.RemoveAll(vendor)
	}

	// try override GOPATH
	// fetched at $TEMP/toubunXXX/src/github.com/owner/repo
	cd := re.CurrentWorkDir()
	ps := strings.Split(cd, "/")
	srcIndex := len(ps) - 1
	for i := range ps {
		if ps[i] == "src" {
			srcIndex = i
			break
		}
	}
	newPath := "/" + path.Join(ps[:srcIndex]...)
	oldPath := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", oldPath)
	os.Setenv("GOPATH", newPath)

	cmd := exec.Command("dep", "ensure", "-update")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("updater;[%s]:dep ensure failed;%s;%v\n", re.CurrentWorkDir(), out, err)
		return err
	}

	log.Printf("updater;[%s]:dep ensure success;%s\n", re.CurrentWorkDir(), out)
	return nil
}

func convertUpdaterConfig(c *k.ModuleConfig) *depConfig {
	cfg := &depConfig{}
	return cfg
}
