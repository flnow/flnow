package routes

import (
	"net/http"
	"time"

	"github.com/flnow/flnow/models"
	"github.com/labstack/echo/v4"
)

// Create flow data
func Create(c echo.Context) (err error) {
	// 数据验证 && 数据创建 && 跳转到详情页
	flow := new(models.Flow)
	if err = c.Bind(flow); err != nil {
		//TODO: sth wrong
		return c.JSON(http.StatusInternalServerError, err)
	}

	flow.Owner = 1
	flow.CreatedAt = time.Now()
	flow.UpdatedAt = time.Now()
	flow.Create()

	return c.JSON(http.StatusOK, flow)
}
