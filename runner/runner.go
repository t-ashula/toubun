package runner

import (
	"fmt"

	k "github.com/t-ashula/toubun/core"
)

// Runner interface describe what fetcher/updater/publisher should impl.
type Runner interface {
	Name() string
	ValidateConfig() error
	Run(k.RunEnv) error
}

// Creater is function that creat Runner
type Creater func(c *k.ModuleConfig) Runner

const (
	fetcherType   = "fetcher"
	updaterType   = "updater"
	publisherType = "publisher"
)

var runners map[string]map[string]Creater
var runnerTypes = []string{fetcherType, updaterType, publisherType}

func newRunner(sub string, c *k.ModuleConfig) Runner {
	name := c.Module
	if name == "" || !isKnownRunner(sub, name) {
		return nil
	}
	return runners[sub][name](c)
}

func register(sub, name string, creater Creater) error {
	if sub == "" || name == "" {
		return fmt.Errorf("runner %s/%s required", sub, name)
	}

	if isKnownRunner(sub, name) {
		return fmt.Errorf("runner %s/%s already registered", sub, name)
	}

	runners[sub][name] = creater
	return nil
}

func unregister(sub, name string) {
	if !isKnownRunner(sub, name) {
		return
	}
	sr := runners[sub]
	delete(sr, name)
}

func isKnownSub(sub string) bool {
	_, knownSub := runners[sub]
	return knownSub
}

func isKnownRunner(sub, name string) bool {
	if !isKnownSub(sub) {
		return false
	}
	_, known := runners[sub][name]
	return known
}

func init() {
	clearRunners()
}

func clearRunners() {
	runners = make(map[string]map[string]Creater)
	clearFetchers()
	clearUpdaters()
	clearPublishers()
}

func clearFetchers() {
	runners[fetcherType] = make(map[string]Creater)
}
func clearUpdaters() {
	runners[updaterType] = make(map[string]Creater)
}
func clearPublishers() {
	runners[publisherType] = make(map[string]Creater)
}
