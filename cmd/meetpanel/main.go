package main

import (
	"fmt"
	"meetpanel/config"
	router "meetpanel/internal/router"
)

func main() {
	// Load config
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	router := router.NewRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf("%v:%v", cfg.HTTP.Hostname, cfg.HTTP.Port)))
}
