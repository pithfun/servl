package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	// Config store
	Config struct {
		HTTP HTTPConfig
	}

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

	// Read config
	viper.SetEnvPrefix("mtpl")
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
