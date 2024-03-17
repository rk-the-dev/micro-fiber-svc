package config

import (
	"fmt"
	"os"
	"strings"
)

// ConfigHelper holds the configuration loaded from environment variables.
type ConfigHelper struct {
	configMap map[string]string
}

// New creates a new instance of ConfigHelper.
func New() *ConfigHelper {
	return &ConfigHelper{
		configMap: make(map[string]string),
	}
}

// Load reads all the environment variables and stores them in a map.
func (ch *ConfigHelper) Load() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			ch.configMap[pair[0]] = pair[1]
		}
	}
	fmt.Println("Config loaded successfully")
}

// Get retrieves the value associated with the given key from the map.
func (ch *ConfigHelper) Get(key string) (string, bool) {
	value, exists := ch.configMap[key]
	return value, exists
}
