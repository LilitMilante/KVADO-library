package app

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort int    `env:"GRPC_PORT"`
	MySQLdsn string `env:"MYSQL_DSN"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}

	cfg := &Config{}

	err = env.ParseWithOptions(cfg, env.Options{RequiredIfNoDef: true})
	if err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}
