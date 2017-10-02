package git

import (
	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

type gitFetcher struct {
	config    *gfConfig
	validated bool
}

type gfConfig struct {
	url   string
	depth int
}

const fetcherName = "git"

func init() {
	runner.RegisterFetcher(fetcherName, newGitFetcher)
}

func newGitFetcher(c *k.ModuleConfig) runner.Runner {
	gf := &gitFetcher{
		config: convertConfig(c),
	}
	return gf
}

func (f *gitFetcher) Name() string {
	return fetcherName
}

func (f *gitFetcher) ValidateConfig() error {
	return nil
}

func (f *gitFetcher) Run(c k.RunEnv) error {
	return nil
}

func convertConfig(c *k.ModuleConfig) *gfConfig {
	return nil
}
