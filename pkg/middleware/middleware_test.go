package middleware

import (
	"os"
	"testing"

	"github.com/ghostyinc/ghosty/config"
	"github.com/ghostyinc/ghosty/pkg/services"
)

var (
	c *services.Container
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnv(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err := c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
