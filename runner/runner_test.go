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
