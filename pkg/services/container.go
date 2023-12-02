package services

import (
	"fmt"
	"gobblin/config"

	"github.com/davecgh/go-spew/spew"
)

type Container struct {
	// Application configuration
	Config *config.Config
	// Validator stores a validator
	Validator *Validator
}

// Create and initialize a new container.
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	return c
}

// Shutdown and disconnect from all services.
func (c *Container) Shutdown() error {
	// TODO: Disconnect from Redis, MySQL, etc.
	spew.Dump("SHUTING DOWN")
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

// Initialize the validator.
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}
