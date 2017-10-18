package bundler

import (
	"fmt"
	"log"
	"os/exec"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const updaterName = "bundler"

type bundlerConfig struct {
	parallel int
}

type bundlerUpdater struct {
	config *bundlerConfig
}

func init() {
	runner.RegisterUpdater(updaterName, newBundlerUpdater)
}

func newBundlerUpdater(c *k.ModuleConfig) runner.Runner {
	cfg := convertUpdaterConfig(c)
	u := &bundlerUpdater{config: cfg}
	return u
}

func (u *bundlerUpdater) Name() string {
	return updaterName
}

func (u *bundlerUpdater) ValidateConfig() error {
	if u.config == nil {
		return fmt.Errorf("no config should error")
	}
	return nil
}

func (u *bundlerUpdater) Run(re k.RunEnv) error {
	args := []string{"update"}
	if u.config.parallel > 0 {
		args = append(args, "--jobs", fmt.Sprintf("%d", u.config.parallel))
	}

	cmd := exec.Command("bundle", args...)
	res := "success"
	out, err := cmd.CombinedOutput()
	if err != nil {
		res = "failed"
	}
	log.Printf("updater;[%s]:%s:%v:%s;%v\n", re.CurrentWorkDir(), res, cmd.Args, out, err)
	return err
}

func convertUpdaterConfig(c *k.ModuleConfig) *bundlerConfig {
	cfg := &bundlerConfig{parallel: 2}
	return cfg
}
