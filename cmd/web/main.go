package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ghostyinc/ghosty/pkg/routes"
	"github.com/ghostyinc/ghosty/pkg/services"
)

func main() {
	// Initiate a new container
	c := services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// Build the router
	routes.BuildRouter(c)

	// Start web server
	go func() {
		srv := http.Server{
			Addr:    fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler: c.Web,
		}

		// TODO: Handle TLS certificates
		if err := c.Web.StartServer(&srv); err != http.ErrServerClosed {
			c.Web.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.Web.Shutdown(ctx); err != nil {
		c.Web.Logger.Fatal(err)
	}
}
