package response

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	STATUS_SUCCESS = "success"
	STATUS_FAILED  = "failed"
	STATUS_ERROR   = "error"
)

var mapHTTPStatus = map[int]string{
	http.StatusBadRequest:          STATUS_FAILED,
	http.StatusUnauthorized:        STATUS_FAILED,
	http.StatusServiceUnavailable:  STATUS_ERROR,
	http.StatusInternalServerError: STATUS_ERROR,
	http.StatusOK:                  STATUS_SUCCESS,
}

type BaseResponse struct {
	Data             interface{} `json:"data"`
	Status           string      `json:"status"`
	TimeElapsedMs    int64       `json:"time_elapsed_ms"`
	TimeElapsedHuman string      `json:"time_elapsed_human"`
}

// getBaseResponse wrap message with base response
func getBaseResponse(c echo.Context, httpCode int, body interface{}) error {
	var timestamp time.Time

	if t, ok := c.Get("timestamp").(time.Time); ok {
		timestamp = t
	}

	return c.JSON(httpCode, BaseResponse{
		Data:             body,
		Status:           mapHTTPStatus[httpCode],
		TimeElapsedMs:    time.Since(timestamp).Milliseconds(),
		TimeElapsedHuman: time.Since(timestamp).String(),
	})
}

// BadRequestResponse for bad request response
func BadRequestResponse(c echo.Context, body interface{}) error {
	return getBaseResponse(c, http.StatusBadRequest, body)
}

// InternalServerErrorResponse for internal server error response
func InternalServerErrorResponse(c echo.Context, body interface{}) error {
	return getBaseResponse(c, http.StatusInternalServerError, body)
}

// OKResponse for OK response
func OKResponse(c echo.Context, body interface{}) error {
	return getBaseResponse(c, http.StatusOK, body)
}

// ServiceUnavailableResponse for service unavailable response
func ServiceUnavailableResponse(c echo.Context, body interface{}) error {
	return getBaseResponse(c, http.StatusServiceUnavailable, body)
}

// UnauthorizedResponse for unauthorized response
func UnauthorizedResponse(c echo.Context, body interface{}) error {
	return getBaseResponse(c, http.StatusUnauthorized, body)
}
