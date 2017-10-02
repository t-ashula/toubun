package runner

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

type testPublisher struct {
	nameValue      string
	validateResult error
	runResult      error
}

func (f *testPublisher) Name() string          { return f.nameValue }
func (f *testPublisher) ValidateConfig() error { return f.validateResult }
func (f *testPublisher) Run(k.RunEnv) error    { return f.runResult }

func TestNewPublisher(t *testing.T) {
	clearPublishers()

	f := NewPublisher(&k.ModuleConfig{Module: ""})
	if f != nil {
		t.Errorf("empty name module should return nil")
	}
	f = NewPublisher(&k.ModuleConfig{Module: "false"})
	if f != nil {
		t.Errorf("unregisterd module should return nil")
	}
	err := RegisterPublisher("true", func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err != nil {
		t.Errorf("true publisher registreation should not fail; %s\n", err)
	}
	f = NewPublisher(&k.ModuleConfig{Module: "true"})
	if f == nil {
		t.Errorf("registerd module should not return nil\n")
	}
	err = RegisterPublisher("nil", func(*k.ModuleConfig) Runner { return nil })
	if err != nil {
		t.Errorf("nil publisher registreation should not fail; %s\n", err)
	}
	f = NewPublisher(&k.ModuleConfig{Module: "nil"})
	if f != nil {
		t.Errorf("new publisher may return nil;%+v\n", f)
	}
}

func TestRegisterPublisher(t *testing.T) {
	clearPublishers()

	publishers := runners[publisherType]
	if len(publishers) != 0 {
		t.Fatalf("no publisher registered; %v\n", len(publishers))
	}

	err := RegisterPublisher("test", func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err != nil {
		t.Fatalf("should register return no error; %v\n", err)
	}

	err = RegisterPublisher("test", func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err == nil {
		t.Fatalf("same name registration should return error\n")
	}

	err = RegisterPublisher("", func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err == nil {
		t.Fatalf("empty name should return error\n")
	}
}

func TestUnregisterPublisher(t *testing.T) {
	clearPublishers()
	publishers := runners[publisherType]
	if len(publishers) != 0 {
		t.Fatalf("no publisher registered; %v", len(publishers))
	}

	name := "test"
	UnregisterPublisher(name) // should nothing happen

	err := RegisterPublisher(name, func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err != nil {
		t.Fatalf("register failed\n")
	}

	UnregisterPublisher(name) // should nothing happen

	// register -> unregister -> register should success
	err = RegisterPublisher(name, func(*k.ModuleConfig) Runner { return &testPublisher{} })
	if err != nil {
		t.Fatalf("register failed; %v\n", err)
	}
}
