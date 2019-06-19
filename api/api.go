package api

import (
	"fmt"
	"net/http"

	"github.com/flnow/flnow/api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("API modules starting...")
}

// Run API listen
func Run() {
	r := gin.Default()
	r.GET("/", hello)
	r.POST("/flows/create", routes.FlowCreate)
	r.Run(":8081")
}

// Handler
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
