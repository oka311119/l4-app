package main

import (
	"log"

	"github.com/oka311119/l4-app/backend/command/internal/config"
	"github.com/oka311119/l4-app/backend/command/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}

	app := server.NewApp(cfg)

	if err := app.Run(cfg.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
