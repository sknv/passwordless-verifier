package main

import (
	"flag"

	"github.com/sknv/passwordless-verifier/pkg/config"
)

const _defaultConfigFilePath = "./configs/app.toml"

func ConfigFilePathFlag() *string {
	return flag.String("c", _defaultConfigFilePath, "configuration file path")
}

type AppConfig struct {
	Name string `toml:"name" env:"APP_NAME"`
}

type LogConfig struct {
	Level string `toml:"level" env:"LOG_LEVEL" default:"info"`
}

type HTTPConfig struct {
	Address string `toml:"address" env:"HTTP_ADDRESS" default:":4000"`
}

type JaegerConfig struct {
	Host  string  `toml:"host" env:"JAEGER_HOST"`
	Port  string  `toml:"port" env:"JAEGER_PORT"`
	Ratio float64 `toml:"ratio" env:"JAEGER_RATIO"`
}

type Config struct {
	App       AppConfig    `toml:"app"`
	LogConfig LogConfig    `toml:"log"`
	HTTP      HTTPConfig   `toml:"http"`
	Jaeger    JaegerConfig `toml:"jaeger"`
}

func ParseConfig(filePath string) (*Config, error) {
	cfg := &Config{}
	if err := config.Parse(cfg, filePath); err != nil {
		return nil, err
	}

	return cfg, nil
}
