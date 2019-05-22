package main

import (
	"github.com/gin-gonic/gin"
	// "gopkg.in/robfig/cron.v3"
	"fmt"
)

func main() {
	fmt.Println("FL(N)OW")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
