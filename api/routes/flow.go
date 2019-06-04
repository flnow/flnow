package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flnow/flnow/models"
	"github.com/labstack/echo/v4"
)

// CreateFlow to create flow data
func CreateFlow(c echo.Context) (err error) {
	fmt.Println("11111")
	// 数据验证 && 数据创建 && 跳转到详情页
	flow := new(models.Flow)
	if err = c.Bind(flow); err != nil {
		//TODO: sth wrong
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("222222")

	flow.Owner = 1
	flow.State = "CREATED"
	flow.NodeCount = 10
	flow.RunAt = "ALL"
	flow.Pointer = "-1"
	flow.LastExecutedAt = time.Now()
	flow.CreatedAt = time.Now()
	flow.UpdatedAt = time.Now()
	fmt.Println(flow)
	fmt.Println("33333")
	flow.Create()
	fmt.Println("44444")
	return c.JSON(http.StatusOK, flow)
}
