package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"goblin/config"
)

// Container for all the application services
type Container struct {
	// Cache stores the cache client
	Cache *CacheClient
	// Config stores the application configuration
	Config *config.Config
	// Web stores the web framework
	Web *echo.Echo
	// Validator stores the validator
	Validator *Validator
}

// NewContainer creates and initialize a new container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()
	c.initCache()

	return c
}

// Shutdown the container and disconnect from all connections
func (c *Container) Shutdown() error {
	if err := c.Cache.Close(); err != nil {
		return err
	}

	return nil
}

// Application configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web server
func (c *Container) initWeb() {
	c.Web = echo.New()

	switch c.Config.App.Environment {
	case config.EnvProd:
		c.Web.Logger.SetLevel(log.WARN)
	default:
		c.Web.Logger.SetLevel(log.DEBUG)
	}

	c.Web.Validator = c.Validator
}

// initCache initializes the cache client
func (c *Container) initCache() {
	var err error
	if c.Cache, err = NewCacheClient(c.Config); err != nil {
		panic(err)
	}
}
