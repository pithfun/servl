package container

import (
	"fmt"
	"gobblin/pkg/config"
)

type Container struct {
	// Application configuration
	Config *config.Config
}

// Create and initialize a new container.
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	return c
}

// Shutdown and disconnect from all services.
func (c *Container) Shutdown() error {
	// TODO: Disconnect from Redis, MySQL, etc.
	return nil
}

// Initialize the application configuration.
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}
