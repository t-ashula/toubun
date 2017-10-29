package core

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// AppConfig holds whole application config
type AppConfig struct {
	Fetch   ModuleConfig `yaml:"fetch"`
	Update  ModuleConfig `yaml:"update"`
	Publish ModuleConfig `yaml:"publish"`
}

type ModuleConfig struct {
	Module string                      `yaml:"module"`
	Config map[interface{}]interface{} `yaml:"config"`
}

// LoadConfig load AppConfig
func LoadConfig(filePath string) (*AppConfig, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := &AppConfig{}

	err = yaml.UnmarshalStrict(content, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetConfigStringValue(c *ModuleConfig, key string, defaultValue string) (string, bool) {
	value := defaultValue
	v, ok := c.Config[key]
	if ok {
		value, ok = v.(string)
	}
	return value, ok
}

func GetConfigIntValue(c *ModuleConfig, key string, defaultValue int) (int, bool) {
	value := defaultValue
	v, ok := c.Config[key]
	if ok {
		value, ok = v.(int)
	}
	return value, ok
}

func NewModuleConfig(yml string) *ModuleConfig {
	var mc ModuleConfig
	err := yaml.UnmarshalStrict([]byte(yml), &mc)
	if err != nil {
		fmt.Printf("unmershal %s failed; %s", yml, err)
		return nil
	}
	return &mc
}

func GetConfigBoolValue(c *ModuleConfig, key string, defaultValue bool) (bool, bool) {
	value := defaultValue
	v, ok := c.Config[key]
	if ok {
		value, ok = v.(bool)
	}
	return value, ok
}
