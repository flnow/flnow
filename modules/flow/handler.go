package flow

import (
	"encoding/json"
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
		if !db.HasTable(&Flow{}) {
			db.CreateTable(&Flow{})
		}
		// var pl Pipeline
		// json.Unmarshal([]byte(flow.Pipeline), &pl)

		// var pipeline Pipeline
		// pipeline.Unmarshal(flow.Pipeline)

		pipe := Pipeline{}
		if flow.Pipeline != "{}" && len(flow.Pipeline) > 0 {
			if err := json.Unmarshal([]byte(flow.Pipeline), &pipe); err != nil {
				fmt.Println(err)
			}
		}
		transaction := db.Begin()

		if !pipe.IsZero() {
			// not zero value pipeline
			fmt.Println("un-zero pipeline configuration...")
			// Add more steps to transaction
			nodes, nodeConfigs := pipe.ToRelational(flow.ID, "", "", 1)

			bNodes, _ := json.Marshal(nodes)
			bConfigs, _ := json.Marshal(nodeConfigs)
			fmt.Println(string(bNodes))
			fmt.Println(string(bConfigs))
		}
		transaction.LogMode(true)
		transaction.Create(&flow)
		transaction.LogMode(false)
		transaction.Commit()
		c.JSON(http.StatusOK, flow)
	}
}

// Detail of Flow
func Detail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Flow Detail...")
		flow := new(Flow)
		db.Where("id = ?", c.Param("flowID")).First(flow)
		if flow.ID == "" {
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
