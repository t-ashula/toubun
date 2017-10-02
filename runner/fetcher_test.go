package runner

import (
	"testing"

	k "github.com/t-ashula/toubun/core"
)

type testFetcher struct {
	nameValue      string
	validateResult error
	runResult      error
}

func (f *testFetcher) Name() string          { return f.nameValue }
func (f *testFetcher) ValidateConfig() error { return f.validateResult }
func (f *testFetcher) Run(k.RunEnv) error    { return f.runResult }

func TestNewFetcher(t *testing.T) {
	clearFetchers()

	f := NewFetcher(&k.ModuleConfig{Module: ""})
	if f != nil {
		t.Errorf("empty name module should return nil")
	}
	f = NewFetcher(&k.ModuleConfig{Module: "false"})
	if f != nil {
		t.Errorf("unregisterd module should return nil")
	}
	err := RegisterFetcher("true", func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err != nil {
		t.Errorf("true fetcher registreation should not fail; %s\n", err)
	}
	f = NewFetcher(&k.ModuleConfig{Module: "true"})
	if f == nil {
		t.Errorf("registerd module should not return nil\n")
	}
	err = RegisterFetcher("nil", func(*k.ModuleConfig) Runner { return nil })
	if err != nil {
		t.Errorf("nil fetcher registreation should not fail; %s\n", err)
	}
	f = NewFetcher(&k.ModuleConfig{Module: "nil"})
	if f != nil {
		t.Errorf("new fetcher may return nil;%+v\n", f)
	}
}

func TestRegisterFetcher(t *testing.T) {
	clearFetchers()

	fetchers := runners[fetcherType]
	if len(fetchers) != 0 {
		t.Fatalf("no fetcher registered; %v\n", len(fetchers))
	}

	err := RegisterFetcher("test", func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err != nil {
		t.Fatalf("should register return no error; %v\n", err)
	}

	err = RegisterFetcher("test", func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err == nil {
		t.Fatalf("same name registration should return error\n")
	}

	err = RegisterFetcher("", func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err == nil {
		t.Fatalf("empty name should return error\n")
	}
}

func TestUnregisterFetcher(t *testing.T) {
	clearFetchers()
	fetchers := runners[fetcherType]
	if len(fetchers) != 0 {
		t.Fatalf("no fetcher registered; %v", len(fetchers))
	}

	name := "test"
	UnregisterFetcher(name) // should nothing happen

	err := RegisterFetcher(name, func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err != nil {
		t.Fatalf("register failed\n")
	}

	UnregisterFetcher(name) // should nothing happen

	// register -> unregister -> register should success
	err = RegisterFetcher(name, func(*k.ModuleConfig) Runner { return &testFetcher{} })
	if err != nil {
		t.Fatalf("register failed; %v\n", err)
	}
}
