package gh

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

const publisherName = "github:pr"

type githubPublisher struct {
	config         *ghpConfig
	originalConfig *k.ModuleConfig
}

type ghpConfig struct {
	committerName string
	committerMail string
	commitMessage string
	prTitle       string
	prBody        string
	prBaseBranch  string
	apiEndPoint   string
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
	log.Printf("publisher:[%s]:start\n", re.CurrentWorkDir())

	// prepare
	token, ok := re.EnvVars()[githubTokenEnvName]
	if !ok || token == "" {
		return fmt.Errorf("set github oauth2 token to env-var:%s", githubTokenEnvName)
	}

	if err := exec.Command("git", "config", "user.name", p.config.committerName).Run(); err != nil {
		return err
	}

	if err := exec.Command("git", "config", "user.email", p.config.committerMail).Run(); err != nil {
		return err
	}

	// changed into new branch
	log.Printf("publisher:[%s]:git checkout\n", re.CurrentWorkDir())
	branchName := fmt.Sprintf("toubun/upgrade-%d", time.Now().UnixNano())
	args := []string{"checkout", "-b", branchName}
	cmd := exec.Command("git", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("publisher:[%s]:%s:failed:%s:%v\n", re.CurrentWorkDir(), cmd.Args, out, err)
		return err
	}

	// make commit
	log.Printf("publisher:[%s]:git commit\n", re.CurrentWorkDir())
	args = []string{"commit", "--allow-empty", "--add", "--message", fmt.Sprintf("'%s'", p.config.commitMessage)}
	cmd = exec.Command("git", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("publisher:[%s]:%s:failed:%s:%v\n", re.CurrentWorkDir(), cmd.Args, out, err)
		return err
	}

	// push
	owner, repo := p.guessRepository()
	ghHost := p.guessHost()
	repositoryURL := fmt.Sprintf("https://%s:x-oauth-basic@%s/%s/%s", token, ghHost, owner, repo)
	log.Printf("publisher:[%s]:git push %s %s\n", re.CurrentWorkDir(), repositoryURL, branchName)
	args = []string{"push", repositoryURL, branchName}
	cmd = exec.Command("git", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("publisher:[%s]:%s:failed:%s:%v\n", re.CurrentWorkDir(), cmd.Args, out, err)
		return err
	}

	// send pull req
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	ghc := github.NewClient(tc)
	if p.config.apiEndPoint != "" {
		u, err := url.Parse(p.config.apiEndPoint)
		if err == nil {
			ghc.BaseURL = u
		}
	}
	pull := &github.NewPullRequest{
		Title: github.String(p.config.prTitle),
		Head:  github.String(branchName),
		Base:  github.String(p.config.prBaseBranch),
		Body:  github.String(p.config.prBody),
	}

	log.Printf("publisher:[%s]:create pr %s %s\n", re.CurrentWorkDir(), owner, repo)
	_, _, err := ghc.PullRequests.Create(ctx, owner, repo, pull)
	if err != nil {
		return err
	}

	log.Printf("publisher:[%s]:done\n", re.CurrentWorkDir())
	return nil
}

func (p *githubPublisher) guessRepository() (owner string, repo string) {
	res, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil || len(res) == 0 {
		return
	}

	remoteURL := string(res)
	remoteURL = strings.TrimSpace(remoteURL)

	// git@github.com:t-ashula/toubun.git
	if strings.HasPrefix(remoteURL, "git") {
		t1 := strings.SplitN(remoteURL, ":", 2)
		t2 := strings.SplitN(t1[1], "/", 2)
		owner = t2[0]
		repo = t2[1]
		repo = strings.TrimSuffix(repo, ".git")
		log.Printf("guess owner/repo:<%s>:<%s><%s>\n", remoteURL, owner, repo)
		return
	}

	// https://github.com/t-ashula/toubun.git
	if strings.HasPrefix(remoteURL, "http") {
		parts := strings.Split(remoteURL, "/")
		l := len(parts)
		if l < 2 {
			return
		}
		owner = parts[l-1-1]
		repo = parts[l-1-0]
		repo = strings.TrimSuffix(repo, ".git")
		log.Printf("guess owner/repo:<%s>:<%s><%s>\n", remoteURL, owner, repo)
		return
	}

	// can't detect
	owner, repo = "", ""
	return
}

func (p *githubPublisher) guessHost() string {
	endPoint := p.config.apiEndPoint
	if endPoint == "" {
		return "github.com"
	}

	u, _ := url.Parse(endPoint)
	if u.Host == "api.github.com" {
		return "github.com"
	}

	return u.Host
}

func convertPublisherConfig(c *k.ModuleConfig) *ghpConfig {
	if c.Config == nil {
		return nil
	}

	name, _ := k.GetConfigStringValue(c, "committer_name", "")
	mail, _ := k.GetConfigStringValue(c, "committer_mail", "")
	msg, _ := k.GetConfigStringValue(c, "commit_message", "")
	title, _ := k.GetConfigStringValue(c, "pr_title", "")
	body, _ := k.GetConfigStringValue(c, "pr_body", "")
	branch, _ := k.GetConfigStringValue(c, "pr_base_branch", "")
	api, _ := k.GetConfigStringValue(c, "api_endpoint", "")
	ghpc := &ghpConfig{
		committerName: name,
		committerMail: mail,
		commitMessage: msg,
		prTitle:       title,
		prBody:        body,
		prBaseBranch:  branch,
		apiEndPoint:   api,
	}
	return ghpc
}
