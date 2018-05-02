package core

import (
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configPath := getFixturePath("no-such-file.yml")
	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("should fail; file:%s, err:%v", configPath, err)
	}

	configPath = getFixturePath("other-type.yml")
	_, err = LoadConfig(configPath)
	if err == nil {
		t.Fatalf("should fail; file:%s, err:%v", configPath, err)
	}

	configPath = getFixturePath("valid.yml")
	_, err = LoadConfig(configPath)
	if err != nil {
		t.Fatalf("should success; file:%s, err:%v", configPath, err)
	}
	// fmt.Printf("conf:%v, err:%v", c, err)
}

func getFixturePath(name string) string {
	return filepath.Join(".test-fixtures", name)
}

func TestGetConfigStringValue(t *testing.T) {
	mc := NewModuleConfig(`
module: test
config:
  key: 'foobar'
`)
	v, ok := GetConfigStringValue(mc, "key", "default")
	if v != "foobar" {
		t.Errorf("value should not 'default' but '%s'", v)
	}
	if !ok {
		t.Errorf("should ok")
	}
	v, ok = GetConfigStringValue(mc, "nokey", "default")
	if v != "default" {
		t.Errorf("should 'default' but '%s'", v)
	}
	if ok {
		t.Errorf("should not ok")
	}
}

func TestGetConfigIntValue(t *testing.T) {
	mc := NewModuleConfig(`
module: test
config:
  key: 100
`)
	v, ok := GetConfigIntValue(mc, "key", 42)
	if v != 100 {
		t.Errorf("value hould not 'default' but '%d'", v)
	}
	if !ok {
		t.Errorf("should ok")
	}
	v, ok = GetConfigIntValue(mc, "nokey", 42)
	if v != 42 {
		t.Errorf("should 'default' but '%d'", v)
	}
	if ok {
		t.Errorf("should not ok")
	}
}

func TestNewModuleConfi(t *testing.T) {
	mc := NewModuleConfig(`
module: test
config:
  key: 42
`)
	if mc == nil {
		t.Errorf("should not nil")
	}
	mc = NewModuleConfig(`
foo: bar
`)
	if mc != nil {
		t.Errorf("should nil, %+v", mc)
	}
}
