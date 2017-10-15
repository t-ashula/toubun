package gh

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

func TestUpdaterName(t *testing.T) {
	c := &k.ModuleConfig{Module: "github:pr"}
	p := newGithubPublisher(c)
	if p.Name() != "github:pr" {
		t.Fatalf("name should 'github:pr'")
	}
}

func TestPublisherValidateConfig(t *testing.T) {
	patterns := []struct {
		should  bool
		message string
		yml     string
	}{
		{false, "no config should error", ""},
		{false, "empty config should error", `
module: github
config:
`},
		{false, "committer name required", `
module: github
config:
  committer_name: ''
`},
		{false, "committer mail required", `
module: github
config:
  committer_name: 'dev'
  committer_mail: ''
`},
		{false, "pr title required", `
module: github
config:
  committer_name: 'dev'
  committer_mail: 't.ashula+dev@gmail.com'
`},
		{false, "pr body required", `
module: github
config:
  committer_name: 'dev'
  committer_mail: 't.ashula+dev@gmail.com'
  pr_title: 'update dependncies by toubun at %f'
`},
		{true, "pr base branch does not have to present", `
module: github
config:
  committer_name: 'dev'
  committer_mail: 't.ashula+dev@gmail.com'
  pr_title: 'update dependncies by toubun at %f'
  pr_body: 'toubun update '
`},
		{true, "pr base branch present ok", `
module: github
config:
  committer_name: 'dev'
  committer_mail: 't.ashula+dev@gmail.com'
  pr_title: 'update dependncies by toubun at %f'
  pr_body: 'toubun update '
  pr_base_branch: develop
`},
	}

	for _, p := range patterns {
		c := k.NewModuleConfig(p.yml)
		r := newGithubPublisher(c)
		err := r.ValidateConfig()
		// fmt.Printf("%+v;%s\n", c, err)
		if p.should && err != nil || !p.should && err == nil {
			t.Errorf("%s:validate with config:%s should %v; but not;%q", p.message, p.yml, p.should, err)
		}
	}
}

func TestPublisherRun(t *testing.T) {
}
