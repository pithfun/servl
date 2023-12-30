package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	// Complete application configuration
	Config struct {
		App  AppConfig
		HTTP HTTPConfig
	}

	// Application configuration
	AppConfig struct {
		EncryptionKey string
		Environment   environment
		Name          string
		Timeout       time.Duration
	}

	// HTTP server configuration
	HTTPConfig struct {
		Hostname string
		Port     uint16
		TLS      struct {
			Enabled bool
		}
	}
)

func GetConfig() (Config, error) {
	var c Config

	// Set default values
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// Read config
	viper.SetEnvPrefix("goblin")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}

type environment string

const (
	EnvTest environment = "test"
	EnvDev  environment = "development"
	EnvProd environment = "production"
)

func SwitchEnv(env environment) {
	if err := os.Setenv("GOBLIN_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}
