package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	SSLMode       string `mapstructure:"DB_SSL_MODE"`

	Port          string `mapstructure:"PORT"`
	LinksHost     string `mapstructure:"LINKS_HOST"`
}

var EnvConfig *Config //nolint:gochecknoglobals //global Config

func loadEnvVariables() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	_ = viper.BindEnv("DB_USER", "DB_USER")
	_ = viper.BindEnv("DB_PASSWORD", "DB_PASSWORD")
	_ = viper.BindEnv("DB_HOST", "DB_HOST")
	_ = viper.BindEnv("DB_PORT", "DB_PORT")
	_ = viper.BindEnv("DB_NAME", "DB_NAME")
	_ = viper.BindEnv("DB_SSL_MODE", "DB_SSL_MODE")
	_ = viper.BindEnv("PORT", "PORT")
	_ = viper.BindEnv("LINKS_HOST", "LINKS_HOST")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warnf("No .env file found: %s. Falling back to environment variables.", err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatalf("Error unmarshaling config: %s", err)
	}

	return &config
}

func InitConfig() {
	EnvConfig = loadEnvVariables()
}
