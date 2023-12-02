package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"

	services "gobblin/pkg/services"
)

func main() {
	// Start a new container
	c := services.NewContainer()

	defer func() {
		if err := c.Shutdown(); err != nil {
			// TODO: Handle shutdown
			spew.Dump("Shutting container down…")
		}
	}()

	// Web server
	go func() {
	}()

	// GraphQL server
	go func() {
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
