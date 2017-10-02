package runner

import (
	k "github.com/t-ashula/toubun/core"
)

func NewPublisher(c *k.ModuleConfig) Runner {
	return newRunner(publisherType, c)
}

func RegisterPublisher(name string, creator Creater) error {
	return register(publisherType, name, creator)
}

func UnregisterPublisher(name string) {
	unregister(publisherType, name)
}
