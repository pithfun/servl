package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_GetConfig(t *testing.T) {
	_, err := GetConfig()
	require.NoError(t, err)
}

func TestConfig_SwitchEnvironment(t *testing.T) {
	var env environment = EnvTest
	SwitchEnv(env)

	cfg, err := GetConfig()
	require.NoError(t, err)

	assert.Equal(t, env, cfg.App.Environment)
}

func TestConfig_OverrideWithEnvVariables(t *testing.T) {
	os.Setenv("TINYBLOB_HTTP_HOSTNAME", "127.0.0.1")
	os.Setenv("TINYBLOB_HTTP_PORT", "8080")

	cfg, err := GetConfig()
	require.NoError(t, err)

	assert.Equal(t, "127.0.0.1", cfg.HTTP.Hostname)
	assert.Equal(t, uint16(8080), cfg.HTTP.Port)
}
