package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	t.Run("should load config", func(t *testing.T) {
		_, err := GetConfig()
		require.NoError(t, err)
	})

	t.Run("override with env variables", func(t *testing.T) {
		os.Setenv("MTPL_HTTP_HOSTNAME", "127.0.0.1")
		os.Setenv("MTPL_HTTP_PORT", "8080")

		cfg, err := GetConfig()
		require.NoError(t, err)

		require.Equal(t, "127.0.0.1", cfg.HTTP.Hostname)
		require.Equal(t, uint16(8080), cfg.HTTP.Port)
	})
}
