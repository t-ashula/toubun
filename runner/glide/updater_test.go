package glide

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

func TestUpdaterName(t *testing.T) {
	u := newGlideUpdater(&k.ModuleConfig{Module: "glide"})
	if u.Name() != "glide" {
		t.Fatalf("name should 'glide'")
	}
}

func TestUpdaterValidateConfig(t *testing.T) {
	patterns := []struct {
		should  bool
		message string
		yml     string
	}{
		{false, "no config should error", ""},
		{true, "empty config mey valid", `
module: glide
config:
`},
	}

	for _, p := range patterns {
		c := k.NewModuleConfig(p.yml)
		r := newGlideUpdater(c)
		err := r.ValidateConfig()
		if p.should && err != nil || !p.should && err == nil {
			t.Errorf("%s:validate with config:%s should %v; but not;%q", p.message, p.yml, p.should, err)
		}
	}

}
