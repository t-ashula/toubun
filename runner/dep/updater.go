package godep

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const updaterName = "go:dep"

type godepConfig struct{}

type godepUpdater struct {
	config *godepConfig
}

func init() {
	runner.RegisterUpdater(updaterName, newGodepUpdater)
}

func newGodepUpdater(c *k.ModuleConfig) runner.Runner {
	cfg := convertUpdaterConfig(c)
	u := &godepUpdater{config: cfg}
	return u
}

func (u *godepUpdater) Name() string {
	return updaterName
}

func (u *godepUpdater) ValidateConfig() error {
	if u.config == nil {
		return fmt.Errorf("no config should error")
	}
	return nil
}

func (u *godepUpdater) Run(re k.RunEnv) error {
	removeDeps(re)

	err := exec.Command("dep", "ensure").Run()
	if err != nil {
		log.Printf("failed;[%s]:dep ensure\n", re.CurrentWorkDir())
		return err
	}
	log.Printf("success;[%s]:dep ensure\n", re.CurrentWorkDir())
	return nil
}

func removeDeps(re k.RunEnv) {
	lock := filepath.Join(re.CurrentWorkDir(), "Gopkg.toml")
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

func convertUpdaterConfig(c *k.ModuleConfig) *godepConfig {
	cfg := &godepConfig{}
	return cfg
}
