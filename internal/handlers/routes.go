package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/samaita/clip-by-clip/pkg/middlewares"
	"github.com/samaita/clip-by-clip/pkg/response"
)

func InitRoutes(e *echo.Echo) {

	// setup top level api
	e.Use(middlewares.TimeElapsed)
	e.GET("/health", response.HealthCheck)

	// Define routes
	apiV1Video := e.Group("/api/v1/video")
	apiV1Video.POST("/reverse", ReverseVideo)
}
