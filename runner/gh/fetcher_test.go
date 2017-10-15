package gh

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	k "github.com/t-ashula/toubun/core"
)

func TestFetcherName(t *testing.T) {
	c := &k.ModuleConfig{Module: "github"}
	f := newGithubFetcher(c)
	if f.Name() != "github" {
		t.Fatalf("name should 'github'")
	}
}

func TestFetcherValidteConfig(t *testing.T) {

	patterns := []struct {
		should  bool
		message string
		yml     string
	}{
		{false, "no config should error", ""},
		{false, "no url should error", `
module: github
config:
`},
		{false, "empty url should error", `
module: github
config:
  url: ''
`},
		{false, "other than git/https scheme should error", `
module: github
config:
  url: 'file:///home/foo/path/to'
`},
		{false, "not parsable url should fail", `
module: github
config:
  url: :foo
`},
		{false, "git(ssh) url should invalid", `
module: github
config:
  url: 'git@github.com:t-ashula/toubun/'
`},
		{true, "https url should valid", `
module: github
config:
  url: 'https://github.com/t-ashula/toubun/'
`},
		{false, "negative depth should invalid", `
module: github
config:
  url: 'https://github.com/t-ashula/toubun/'
  depth: -2
`},
	}

	for _, p := range patterns {
		c := k.NewModuleConfig(p.yml)
		f := newGithubFetcher(c)
		err := f.ValidateConfig()
		// fmt.Printf("%+v;%s\n", c, err)
		if p.should && err != nil || !p.should && err == nil {
			t.Errorf("validate with config:%s should %v; but not", p.yml, p.should)
		}
	}
}

func TestFetcherRun(t *testing.T) {

}

func TestFetcherAuthedRespositoryURL(t *testing.T) {
	yml := `
module: github
config:
  url: 'https://github.com/t-ashula/toubun/'
`
	c := k.NewModuleConfig(yml)
	f := newGithubFetcher(c).(*githubFetcher) //
	if f == nil {
		t.Fatal("new ghf faile")
	}

	err := f.ValidateConfig()
	if err != nil {
		t.Fatal("config validation failed")
	}

	re := k.NewRunEnv()
	uri, err := f.authedRepositoryURL(re)
	if err != nil {
		t.Fatal("base url should no error for valid config")
	}
	expected := "https://github.com/t-ashula/toubun/"
	if uri != expected {
		t.Fatalf("something wrong; %s", uri)
	}

	// with token

	// TODO: use random string generator?
	token := strconv.FormatInt(time.Now().UnixNano(), 36)
	os.Setenv(githubTokenEnvName, token)
	f2 := newGithubFetcher(c).(*githubFetcher)
	uri, err = f2.authedRepositoryURL(re)
	if err != nil {
		t.Fatal("base url should no error for valid config")
	}
	expected = fmt.Sprintf("https://%s:%s@github.com/t-ashula/toubun/", token, "x-oauth-basic")
	if uri != expected {
		t.Fatalf("something wrong;\nexpected:%s\nactual:%s\n", expected, uri)
	}
}
