package gh

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os/exec"
	"path"
	"strings"

	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const githubTokenEnvName = "TOUBUN_GITHUB_TOKEN"
const fetcherName = "github"

type githubFetcher struct {
	config        *ghfConfig
	orginalConfig *k.ModuleConfig

	validConfig bool
}

type ghfConfig struct {
	url    string
	branch string
	depth  int
}

func init() {
	err := runner.RegisterFetcher(fetcherName, newGithubFetcher)
	if err != nil {
		log.Printf("gh/fetcher registration failed; %s", err)
	}
}

func newGithubFetcher(c *k.ModuleConfig) runner.Runner {
	ghfc := convertFetcherConfig(c)
	ghf := &githubFetcher{
		config:        ghfc,
		orginalConfig: c,
		validConfig:   false,
	}
	return ghf
}

func (f *githubFetcher) Name() string {
	return fetcherName
}

func (f *githubFetcher) ValidateConfig() error {
	if f.config == nil {
		return fmt.Errorf("no config found")
	}

	if f.config.url == "" {
		return fmt.Errorf("url required")
	}

	u, err := url.Parse(f.config.url)
	if err != nil {
		return fmt.Errorf("unparsable url;%s;%s", f.config.url, err)
	}

	if u.Scheme != "https" {
		return fmt.Errorf("unsupported scheme; git or https required")
	}

	if f.config.depth < 0 {
		return fmt.Errorf("clone depth option should not be negative")
	}

	f.validConfig = true
	return nil
}

func (f *githubFetcher) Run(re k.RunEnv) error {
	if !f.validConfig {
		if err := f.ValidateConfig(); err != nil {
			return fmt.Errorf("invalid config;%s", err)
		}
	}

	repositoryURL, err := f.authedRepositoryURL(re)
	if err != nil {
		log.Printf("authedRepositoryURL failed\n")
		return err
	}

	cloneDir, _ := f.cloneDir(re)

	// TODO: use go-git package?
	git := "git"
	args := []string{"clone", repositoryURL}
	if f.config.depth > 0 {
		args = append(args, "--depth", fmt.Sprintf("%d", f.config.depth))
	}
	if f.config.branch != "" {
		args = append(args, "--branch", f.config.branch)
	}
	args = append(args, cloneDir)

	cmd := exec.Command(git, args...)
	err = cmd.Run()
	if err != nil {
		log.Printf("failed;[%s]:git %+v\n", re.CurrentWorkDir(), args)
		return err
	}

	re.ChangeWorkDir(cloneDir, true)
	log.Printf("success;[%s]:git %+v\n", re.CurrentWorkDir(), args)
	return nil
}

func (f *githubFetcher) cloneDir(re k.RunEnv) (string, error) {
	cd := re.CurrentWorkDir()
	base, err := ioutil.TempDir(cd, "toubun")
	if err != nil {
		return cd, err
	}

	u, err := url.Parse(f.config.url)
	if err != nil {
		return base, err
	}

	ps := strings.Split(u.Path, "/")
	ps[len(ps)-1] = strings.TrimSuffix(ps[len(ps)-1], ".git")

	ds := []string{base, "src", u.Hostname()}
	ds = append(ds, ps...)
	full := path.Join(ds...)

	return full, nil
}

func (f *githubFetcher) authedRepositoryURL(re k.RunEnv) (string, error) {
	u, err := url.Parse(f.config.url)
	if err != nil {
		return "", err
	}

	if u.Scheme != "https" {
		return "", fmt.Errorf("unsupport url:%s", f.config.url)
	}

	envs := re.EnvVars()
	token := envs[githubTokenEnvName]

	if token == "" {
		// TODO: log invalid config, or return error ?
		return f.config.url, nil
	}

	u.User = url.UserPassword(token, "x-oauth-basic")

	return u.String(), nil
}

func convertFetcherConfig(c *k.ModuleConfig) *ghfConfig {
	if c.Module == "" {
		return nil
	}

	u, _ := k.GetConfigStringValue(c, "url", "")

	branch, _ := k.GetConfigStringValue(c, "branch", "")
	if branch == "" {
		// TODO: log 'use default branch'
	}

	depth, _ := k.GetConfigIntValue(c, "depth", 0)

	ghfc := &ghfConfig{
		url:    u,
		branch: branch,
		depth:  depth,
	}

	return ghfc
}
