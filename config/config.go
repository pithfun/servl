package config

import (
	"os"
	"strings"

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
		Environment environment
		Name        string
	}

	// HTTP server configuration
	HTTPConfig struct {
		Hostname string
		Port     uint16
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
	EnvTest environment = "TEST"
	EnvDev  environment = "DEVELOPMENT"
	EnvProd environment = "PRODUCTION"
)

func SwitchEnv(env environment) {
	if err := os.Setenv("GOBLIN_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}
