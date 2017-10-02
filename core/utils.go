package core

import (
	"os"
	"strings"
)

// EnvMap return environment variables hash map
func EnvMap(prefix string) map[string]string {
	vars := os.Environ()
	data := make(map[string]string)
	for _, val := range vars {
		kv := strings.SplitN(val, "=", 2)
		k := kv[0]
		v := kv[1]
		data[k] = v
	}

	if prefix == "" {
		return data
	}

	for k := range data {
		if !strings.HasPrefix(k, prefix) {
			delete(data, k)
		}
	}

	return data
}
