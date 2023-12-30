package services

import (
	"fmt"
	"goblin/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Container struct {
	Config    *config.Config
	Web       *echo.Echo
	Validator *Validator
}

// Create and initialize a new container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()

	return c
}

// Shutdown the container
func (c *Container) Shutdown() error {
	// TODO: Disconnect from Redis, MySQL, etc.
	spew.Dump("SHUTING DOWN")
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

// Validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// Web server
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
