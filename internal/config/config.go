package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Postgres Postgres `yaml:"postgres"`
		Http     Http     `yaml:"http"`
	}

	Postgres struct {
		Url             string        `env-required:"true" yaml:"url"       env:"PG_URL"`
		MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns"`
		MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `env-required:"true" yaml:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime"`
	}

	Http struct {
		Port uint16 `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	const configPath = "./config/config.yaml"

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func MustConfig() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
