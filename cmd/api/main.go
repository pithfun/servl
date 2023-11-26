package main

import (
	services "gobblin/pkg/services"
)

func main() {
	// Start a new container
	c := services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Shutdown()
		}
	}()
}
