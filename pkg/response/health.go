package response

import (
	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

// HealthCheck basic web api
func HealthCheck(c echo.Context) error {
	return OKResponse(c, HealthCheckResponse{
		Message: "Service is up",
	})
}
