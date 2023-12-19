package main

import (
	"log"

	"github.com/oka311119/l4-app/backend/command/internal/config"
	"github.com/oka311119/l4-app/backend/command/internal/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
