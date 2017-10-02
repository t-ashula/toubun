package runner

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

type testUpdater struct {
	nameValue      string
	validateResult error
	runResult      error
}

func (f *testUpdater) Name() string          { return f.nameValue }
func (f *testUpdater) ValidateConfig() error { return f.validateResult }
func (f *testUpdater) Run(k.RunEnv) error    { return f.runResult }

func TestNewUpdater(t *testing.T) {
	clearUpdaters()

	f := NewUpdater(&k.ModuleConfig{Module: ""})
	if f != nil {
		t.Errorf("empty name module should return nil")
	}
	f = NewUpdater(&k.ModuleConfig{Module: "false"})
	if f != nil {
		t.Errorf("unregisterd module should return nil")
	}
	err := RegisterUpdater("true", func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err != nil {
		t.Errorf("true updater registreation should not fail; %s\n", err)
	}
	f = NewUpdater(&k.ModuleConfig{Module: "true"})
	if f == nil {
		t.Errorf("registerd module should not return nil\n")
	}
	err = RegisterUpdater("nil", func(*k.ModuleConfig) Runner { return nil })
	if err != nil {
		t.Errorf("nil updater registreation should not fail; %s\n", err)
	}
	f = NewUpdater(&k.ModuleConfig{Module: "nil"})
	if f != nil {
		t.Errorf("new updater may return nil;%+v\n", f)
	}
}

func TestRegisterUpdater(t *testing.T) {
	clearUpdaters()

	updaters := runners[updaterType]
	if len(updaters) != 0 {
		t.Fatalf("no updater registered; %v\n", len(updaters))
	}

	err := RegisterUpdater("test", func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err != nil {
		t.Fatalf("should register return no error; %v\n", err)
	}

	err = RegisterUpdater("test", func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err == nil {
		t.Fatalf("same name registration should return error\n")
	}

	err = RegisterUpdater("", func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err == nil {
		t.Fatalf("empty name should return error\n")
	}
}

func TestUnregisterUpdater(t *testing.T) {
	clearUpdaters()
	updaters := runners[updaterType]
	if len(updaters) != 0 {
		t.Fatalf("no updater registered; %v", len(updaters))
	}

	name := "test"
	UnregisterUpdater(name) // should nothing happen

	err := RegisterUpdater(name, func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err != nil {
		t.Fatalf("register failed\n")
	}

	UnregisterUpdater(name) // should nothing happen

	// register -> unregister -> register should success
	err = RegisterUpdater(name, func(*k.ModuleConfig) Runner { return &testUpdater{} })
	if err != nil {
		t.Fatalf("register failed; %v\n", err)
	}
}
