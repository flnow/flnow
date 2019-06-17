package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck to know is http server alive
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "flnow")
}
