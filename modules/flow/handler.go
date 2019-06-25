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
		fmt.Println(flow)
		flow.initialize()
		fmt.Println(flow)
		db.Create(&flow)

		c.JSON(http.StatusOK, flow)
	}
}

// Detail of Flow
func Detail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Flow Detail...")
		// fmt.Println(c.Param("test_empty"))
		// if c.Param("flowID") == "" {
		// 	c.String(http.StatusInternalServerError, "No ID")
		// }
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
		fmt.Println(flow)
		if err := c.Bind(flow); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println(flow)
		db.Model(&flow).Where("id = ?", flow.ID).Updates(flow)

		c.JSON(http.StatusOK, flow)
	}
}
