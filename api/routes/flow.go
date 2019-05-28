package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create flow data
func Create(c echo.Context) error {
	// 数据验证 && 数据创建 && 跳转到详情页
	// models.Flow

	return c.String(http.StatusOK, "created")
}
