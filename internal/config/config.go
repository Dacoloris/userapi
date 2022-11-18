package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort  int    `env:"APP_PORT" envDefault:"3333"`
	FileName string `env:"FILE_NAME" envDefault:"users.json"`
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
