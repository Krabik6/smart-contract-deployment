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
	MumbaiPrivateKey     string `env:"MUMBAI_PRIVATE_KEY"`
	MumbaiProvider       string `env:"MUMBAI_PROVIDER"`
	SepoliaPrivateKey    string `env:"SEPOLIA_PRIVATE_KEY"`
	SepoliaProvider      string `env:"SEPOLIA_PROVIDER"`
	BscTestnetPrivateKey string `env:"BSC_TESTNET_PRIVATE_KEY"`
	BscTestnetProvider   string `env:"BSC_TESTNET_PROVIDER"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Config struct {
	AppConfig
	EnvConfig
	Server
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

	var srvConfig Server
	if err := viper.UnmarshalKey("server", &srvConfig); err != nil {
		return nil, err
	}

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	envConfig := EnvConfig{
		MumbaiPrivateKey:     os.Getenv("MUMBAI_PRIVATE_KEY"),
		MumbaiProvider:       os.Getenv("MUMBAI_PROVIDER"),
		SepoliaPrivateKey:    os.Getenv("SEPOLIA_PRIVATE_KEY"),
		SepoliaProvider:      os.Getenv("SEPOLIA_PROVIDER"),
		BscTestnetPrivateKey: os.Getenv("BSC_TESTNET_PRIVATE_KEY"),
		BscTestnetProvider:   os.Getenv("BSC_TESTNET_PROVIDER"),
	}

	config := &Config{
		AppConfig: appConfig,
		EnvConfig: envConfig,
		Server:    srvConfig,
	}
	return config, nil
}
