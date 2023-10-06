package main

import (
	"github.com/labstack/echo/v4"
	"github.com/samaita/clip-by-clip/configs"
	"github.com/samaita/clip-by-clip/internal/handlers"
	"github.com/samaita/clip-by-clip/pkg/log"
)

func init() {
	configs.LoadConfig()
	log.InitLogger()
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Init Routes
	handlers.InitRoutes(e)

	// Start the server
	err := e.Start(configs.Conf.App.Server.Port)
	if err != nil {
		log.Sugar.Fatalf("Failed to start, %v", err)
	}
}
