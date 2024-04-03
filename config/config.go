package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database       DatabaseConfig
		Server         ServerConfig
		Authentication AuthenticationConfig
	}

	DatabaseConfig struct {
		DSN string
	}

	ServerConfig struct {
		Name         string
		ApiPort      string
		WebPort      string
		Debug        bool
		Environment  string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	AuthenticationConfig struct {
		Key       string
		SecretKey string
		SaltKey   string
	}
)

func LoadConfig(env string) (Config, error) {
	viper.SetConfigFile(fmt.Sprintf("config/config-%s.yaml", env))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		return Config{}, err
	}
	return config, nil
}
