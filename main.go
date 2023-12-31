package main

import (
	"github.com/labstack/echo/v4"
	"github.com/samaita/clip-by-clip/configs"
	"github.com/samaita/clip-by-clip/pkg/log"
	"github.com/samaita/clip-by-clip/pkg/middlewares"
	"github.com/samaita/clip-by-clip/pkg/response"
)

func init() {
	configs.LoadConfig()
}

func main() {
	// Create a new Echo instance
	e := echo.New()
	e.Use(middlewares.TimeElapsed)

	// Define routes
	e.GET("/health", response.HealthCheck)

	// Start the server
	err := e.Start(configs.Conf.App.Server.Port)
	if err != nil {
		log.Sugar.Fatalf("Failed to start, %v", err)
	}
}
