package main

import (
	"flag"

	"github.com/sknv/passwordless-verifier/pkg/config"
)

const (
	_defaultConfigFilePath = "./configs/app.toml"
)

func ConfigFilePathFlag() *string {
	return flag.String("c", _defaultConfigFilePath, "configuration file path")
}

type LogConfig struct {
	Level string `toml:"level" env:"LOG_LEVEL" default:"info"`
}

type Config struct {
	LogConfig LogConfig `toml:"log"`
}

func ParseConfig(filePath string) (*Config, error) {
	cfg := &Config{}
	if err := config.Parse(cfg, filePath); err != nil {
		return nil, err
	}
	return cfg, nil
}
