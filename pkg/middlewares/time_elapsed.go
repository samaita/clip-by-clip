package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func TimeElapsed(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("timestamp", time.Now())
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
