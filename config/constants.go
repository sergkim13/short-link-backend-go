package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type config struct {
	DBHost       string         `env:"DB_HOST" envDefault:"localhost"`
	DBPort       int            `env:"DB_PORT" envDefault:"5432"`
	DBUser       string         `env:"DB_USER" envDefault:"postgres"`
	DBPassword   string         `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName		 string         `env:"DB_NAME" envDefault:"short_link_db"`
}

func GetConfig() (config) {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal(err)
	}

	return cfg
}
