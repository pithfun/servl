package services

import (
	"os"
	"testing"

	"github.com/tiny-blob/tinyblob/config"
)

var (
	c *Container
)

func TestMain(m *testing.M) {
	// Set test environment
	config.SwitchEnv(config.EnvTest)

	// Setup
	c = NewContainer()

	// TODO: DB creation and migration

	// Run tests
	exitCode := m.Run()

	// Teardown
	if err := c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}
