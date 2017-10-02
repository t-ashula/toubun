package github

import (
	"fmt"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const publisherName = "github"

type githubPublisher struct {
	config         *ghpConfig
	originalConfig *k.ModuleConfig
}

type ghpConfig struct {
	committerName string
	committerMail string
	prTitle       string
	prBody        string
	prBaseBranch  string
}

func init() {
	runner.RegisterPublisher(publisherName, newGithubPublisher)
}

func newGithubPublisher(c *k.ModuleConfig) runner.Runner {
	ghpc := convertPublisherConfig(c)
	ghp := &githubPublisher{
		config:         ghpc,
		originalConfig: c,
	}
	return ghp
}

func (p *githubPublisher) Name() string {
	return publisherName
}

func (p *githubPublisher) ValidateConfig() error {
	if p.config == nil {
		return fmt.Errorf("no config found")
	}
	if p.config.committerName == "" {
		return fmt.Errorf("no committer name found")
	}
	if p.config.committerMail == "" {
		return fmt.Errorf("no committer mail found")
	}
	if p.config.prTitle == "" {
		return fmt.Errorf("no pr title found")
	}
	if p.config.prBody == "" {
		return fmt.Errorf("no pr title found")
	}

	return nil
}

func (p *githubPublisher) Run(re k.RunEnv) error {
	return nil
}

func convertPublisherConfig(c *k.ModuleConfig) *ghpConfig {
	if c.Config == nil {
		return nil
	}

	name, _ := k.GetConfigStringValue(c, "committer_name", "")
	mail, _ := k.GetConfigStringValue(c, "committer_mail", "")
	title, _ := k.GetConfigStringValue(c, "pr_title", "")
	body, _ := k.GetConfigStringValue(c, "pr_body", "")
	branch, _ := k.GetConfigStringValue(c, "pr_base_branch", "")
	ghpc := &ghpConfig{
		committerName: name,
		committerMail: mail,
		prTitle:       title,
		prBody:        body,
		prBaseBranch:  branch,
	}
	return ghpc
}
