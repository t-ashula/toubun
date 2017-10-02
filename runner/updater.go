package runner

import (
	k "github.com/t-ashula/toubun/core"
)

func NewUpdater(c *k.ModuleConfig) Runner {
	return newRunner(updaterType, c)
}

func RegisterUpdater(name string, creator Creater) error {
	return register(updaterType, name, creator)
}

func UnregisterUpdater(name string) {
	unregister(updaterType, name)
}
