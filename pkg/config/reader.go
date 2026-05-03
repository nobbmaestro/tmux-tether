package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadUserConfig(path string) *UserConfig {
	cfg := GetDefaultUserConfig()

	f, err := os.Open(path)
	if err != nil {
		return cfg
	}

	defer func() {
		if cerr := f.Close(); err == nil {
			err = cerr
		}
	}()

	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return cfg
	}

	return cfg
}
