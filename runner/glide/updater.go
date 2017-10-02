package glide

import (
	"fmt"

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
	return nil
}

func convertUpdaterConfig(c *k.ModuleConfig) *guConfig {
	if c.Module == "" {
		return nil
	}
	cfg := &guConfig{}
	return cfg
}
