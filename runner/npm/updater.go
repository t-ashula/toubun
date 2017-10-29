package npm

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const updaterName = "npm"

type npmConfig struct {
	dev bool
}
type npmUpdater struct {
	config         *npmConfig
	originalConfig *k.ModuleConfig
}

func init() {
	runner.RegisterUpdater(updaterName, newNpmUpdater)
}

func newNpmUpdater(c *k.ModuleConfig) runner.Runner {
	cfg := convertUpdaterConfig(c)
	u := &npmUpdater{
		config:         cfg,
		originalConfig: c,
	}
	return u
}

func (u *npmUpdater) Name() string {
	return updaterName
}

func (u *npmUpdater) ValidateConfig() error {
	if u.config == nil {
		return fmt.Errorf("no config should error")
	}

	dev, ok := u.originalConfig.Config["dev"].(string)
	if ok && (dev != "" && dev != "true" && dev != "false") {
		return fmt.Errorf("no config should error")
	}

	return nil
}

func (u *npmUpdater) Run(re k.RunEnv) error {
	module := filepath.Join(re.CurrentWorkDir(), "node_modules")
	if _, err := os.Stat(module); err == nil {
		os.RemoveAll(module)
	}

	shrink := filepath.Join(re.CurrentWorkDir(), "npm-shrinkwrap.json")
	if _, err := os.Stat(shrink); err == nil {
		os.Remove(shrink)
	}

	args := []string{"update"}
	if u.config.dev {
		args = append(args, "--dev")
	}
	cmd := exec.Command("npm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("updater;[%s]:dep ensure failed;%s;%v\n", re.CurrentWorkDir(), out, err)
		return err
	}

	log.Printf("updater;[%s]:dep ensure success;%s\n", re.CurrentWorkDir(), out)
	return nil
}

func convertUpdaterConfig(c *k.ModuleConfig) *npmConfig {
	dev, _ := k.GetConfigBoolValue(c, "dev", false)
	cfg := &npmConfig{dev: dev}
	return cfg
}
