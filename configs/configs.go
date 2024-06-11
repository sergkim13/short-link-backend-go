package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
type Config struct {
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword	  string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	SSLMode       string `mapstructure:"DB_SSL_MODE"`
	
	Port          string `mapstructure:"PORT"`
}

var EnvConfig *Config

func loadEnvVariables() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error reading config file: %s", err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatalf("error unmarshaling config file: %s", err.Error())
	}

	return
}


func InitConfig() {
	EnvConfig = loadEnvVariables()
}
