package glide

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const updaterName = "glide"

type guConfig struct {
}

type glideUpdater struct {
	config *guConfig
}

func init() {
	runner.RegisterUpdater(updaterName, newGlideUpdater)
}

func newGlideUpdater(c *k.ModuleConfig) runner.Runner {
	cfg := convertUpdaterConfig(c)
	u := &glideUpdater{config: cfg}
	return u
}

func (u *glideUpdater) Name() string {
	return updaterName
}

func (u *glideUpdater) ValidateConfig() error {
	if u.config == nil {
		return fmt.Errorf("no config should error")
	}
	return nil
}

func (u *glideUpdater) Run(re k.RunEnv) error {
	lock := filepath.Join(re.CurrentWorkDir(), "glide.lock")
	_, err := os.Stat(lock)
	if err == nil {
		os.Remove(lock)
	}
	err = exec.Command("glide", "up").Run()
	if err != nil {
		log.Printf("failed;[%s]:glide up\n", re.CurrentWorkDir())
		return err
	}
	log.Printf("success;[%s]:glide up\n", re.CurrentWorkDir())
	return err
}

func convertUpdaterConfig(c *k.ModuleConfig) *guConfig {
	if c.Module == "" {
		return nil
	}
	cfg := &guConfig{}
	return cfg
}
