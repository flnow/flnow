package flow

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create Flow
func Create(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Flow create...")
		var flow Flow
		if err := c.Bind(&flow); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		flow.initialize()
		db.Create(&flow)

		c.JSON(http.StatusOK, flow)
	}
}

// Detail of Flow
func Detail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Flow Detail...")
		flow := new(Flow)
		db.Where("id = ?", c.Param("flowID")).First(flow)
		if flow.ID == 0 {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		c.JSON(http.StatusOK, flow)
	}
}

// Update a flow
func Update(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Flow Update...")
		flow := new(Flow)
		if err := c.Bind(flow); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		db.Model(&flow).Where("id = ?", flow.ID).Updates(flow)

		c.JSON(http.StatusOK, flow)
	}
}

// List of flows
func List(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Flow List...")
		flows := make([]Flow, 0)
		db.Find(&flows)
		c.JSON(http.StatusOK, flows)
	}
}
