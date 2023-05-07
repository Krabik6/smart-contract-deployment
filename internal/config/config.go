package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Image string `yaml:"image"`
}

type EnvConfig struct {
	PrivateKey string `env:"PRIVATE_KEY"`
	Url        string `env:"URL"`
}

type Config struct {
	AppConfig
	EnvConfig
}

func Load() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var appConfig AppConfig
	if err := viper.UnmarshalKey("app", &appConfig); err != nil {
		return nil, err
	}

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	envConfig := EnvConfig{
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		Url:        os.Getenv("URL"),
	}

	config := &Config{
		AppConfig: appConfig,
		EnvConfig: envConfig,
	}
	return config, nil
}
