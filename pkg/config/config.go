package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/caarlos0/env.v2"
)

type Config struct {
	ServiceName      string
	Port             int    `env:"PORT" required:"true"`
	PostgresDB       string `env:"POSTGRES_DB" required:"true"`
	PostgresHost     string `env:"POSTGRES_HOST" required:"true"`
	PostgresPort     int    `env:"POSTGRES_PORT" required:"true"`
	PostgresUser     string `env:"POSTGRES_USER" required:"true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" required:"true"`
}

func New() *Config {
	// TODO(doug): Should explicitly check and fail here if certain env vars are not present
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
