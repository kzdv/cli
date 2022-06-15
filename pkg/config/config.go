package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

const ConfigFile = "config.yaml"

type Config struct {
	Token string `json:"token" yaml:"token"`
}

func GetConfigDir() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(config, "zdv"), nil
}

func MakeConfigDir() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(configDir, 0700)
}

func ReadConfig() (*Config, error) {
	cfg := DefaultConfig()

	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(dir, ConfigFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func DefaultConfig() *Config {
	return &Config{
		Token: "",
	}
}
