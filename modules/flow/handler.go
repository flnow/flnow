package flow

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create Flow
func Create(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Flow create...")
		flow := new(Flow)
		if err := c.Bind(flow); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		flow.initialize()
		db.Create(flow)

		c.JSON(http.StatusOK, flow)
	}
}
