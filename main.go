package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define routes
	e.GET("/", helloHandler)

	// Start the server
	e.Start(":8080")
}

// Handler function for the root route
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
