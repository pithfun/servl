package main

import (
	container "gobblin/pkg/services"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// Start a new container
	c := container.NewContainer()
	defer c.Shutdown()

	// Spew the container
	spew.Dump(c)
}
