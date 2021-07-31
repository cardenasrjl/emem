package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	Environment     string `envconfig:"ENVIRONMENT" required:"true"`
	DatabaseHost    string `envconfig:"DATABASE_HOST" required:"true"`
	DatabasePort    string `envconfig:"DATABASE_PORT" required:"true"`
	DatabaseUser    string `envconfig:"DATABASE_USER" required:"true"`
	DatabasePass    string `envconfig:"DATABASE_PASS" required:"true"`
	DatabaseSchema  string `envconfig:"DATABASE_SCHEMA" required:"true"`
	DatabaseTimeout string `envconfig:"DATABASE_TIMEOUT" required:"true"`

	HttpPort string `envconfig:"HTTP_PORT" required:"true"`
	GrpcPort string `envconfig:"GRPC_PORT" required:"true"`

	LogLevel int `envconfig:"LOG_LEVEL" required:"true"`
}

func New() (*Config, error) {
	gotenv.Load(".env")
	gotenv.Load("../.env")
	gotenv.Load("../../.env")
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
