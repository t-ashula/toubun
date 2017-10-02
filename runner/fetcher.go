package runner

import (
	k "github.com/t-ashula/toubun/core"
)

func NewFetcher(c *k.ModuleConfig) Runner {
	return newRunner(fetcherType, c)
}

func RegisterFetcher(name string, creater Creater) error {
	return register(fetcherType, name, creater)
}

func UnregisterFetcher(name string) {
	unregister(fetcherType, name)
}
