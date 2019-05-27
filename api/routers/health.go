package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck to know is http server alive
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "flnow")
}
