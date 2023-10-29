package container

import (
	"gobblin/pkg/config"

	"github.com/davecgh/go-spew/spew"
)

// Provides an easy way to do dependency injection.
type Container struct {
	Config *config.Config // The application config
}

func NewContainer() *Container {
	c := new(Container)

	return c
}

// Shutdown the container and disconnect from all services.
func (c *Container) Shutdown() error {
	// TODO: Disconnect from Redis, MySQL, etc.
	spew.Dump("Shutting down container...")
	return nil
}
