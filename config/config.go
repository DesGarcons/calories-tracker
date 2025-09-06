package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		Telegram `yaml:"telegram"`
		Postgres `yaml:"postgres"`
	}

	Telegram struct {
		Token string `env-required:"true" yaml:"token" env:"TG_TOKEN"`
	}
	Postgres struct {
		User     string `env-required:"true" yaml:"user" env:"POSTGRES_USER"`
		Password string `env-required:"true" yaml:"password" env:"POSTGRES_PASSWORD"`
		Url      string `env-required:"true" yaml:"url" env:"POSTGRES_URL"`
	}
)

func NewConfig(pathToEnv string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(pathToEnv, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
