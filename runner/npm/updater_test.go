package npm

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

func TestUpdaterName(t *testing.T) {
	c := &k.ModuleConfig{Module: "npm"}
	p := newNpmUpdater(c)
	if p.Name() != "npm" {
		t.Fatalf("name should 'npm'")
	}
}

func TestUpdaterValidateConfig(t *testing.T) {
	patterns := []struct {
		should  bool
		message string
		yml     string
	}{
		{true, "no config should ok", ""},
		{true, "empty config should be ok", `
module: npm
config:
`},
		{true, "accept dev flag setting", `
module: npm
config:
  dev: true
`},
	}

	for _, p := range patterns {
		c := k.NewModuleConfig(p.yml)
		r := newNpmUpdater(c)
		err := r.ValidateConfig()
		// fmt.Printf("%+v;%s\n", c, err)
		if p.should && err != nil || !p.should && err == nil {
			t.Errorf("%s:validate with config:%s should %v; but not;%q", p.message, p.yml, p.should, err)
		}
	}
}

// func TestUpdaterRun(t *testing.T) {}
