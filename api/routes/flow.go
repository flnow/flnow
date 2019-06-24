package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flnow/server/models"
	"github.com/gin-gonic/gin"
)

// FlowCreate to create flow data
func FlowCreate(c *gin.Context) {
	fmt.Println("Flow create")
	// 数据验证 && 数据创建 && 跳转到详情页
	flow := new(models.Flow)
	if err := c.Bind(flow); err != nil {
		//TODO: sth wrong
		c.JSON(http.StatusInternalServerError, err)
	}

	flow.Owner = 1
	flow.State = "CREATED"
	flow.NodeCount = 10
	flow.RunAt = "ANY"
	flow.Pointer = "-1"
	flow.LastExecutedAt = time.Now()
	flow.CreatedAt = time.Now()
	flow.UpdatedAt = time.Now()
	fmt.Println(flow)

	flow.Create()

	c.JSON(http.StatusOK, flow)
}

// Flows is the data list
func Flows(c *gin.Context) {
	fmt.Println("Flow list")

}
