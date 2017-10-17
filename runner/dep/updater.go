package dep

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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
	removeDeps(re)

	err := exec.Command("dep", "ensure", "-update").Run()
	if err != nil {
		log.Printf("failed;[%s]:dep ensure\n", re.CurrentWorkDir())
		return err
	}
	log.Printf("success;[%s]:dep ensure\n", re.CurrentWorkDir())
	return nil
}

func removeDeps(re k.RunEnv) {
	lock := filepath.Join(re.CurrentWorkDir(), "Gopkg.lock")
	_, err := os.Stat(lock)
	if err == nil {
		os.Remove(lock)
	}
	vendor := filepath.Join(re.CurrentWorkDir(), "vendor")
	_, err = os.Stat(vendor)
	if err == nil {
		os.RemoveAll(vendor)
	}
}

func convertUpdaterConfig(c *k.ModuleConfig) *depConfig {
	cfg := &depConfig{}
	return cfg
}
