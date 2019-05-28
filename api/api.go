package api

import (
	"fmt"
	"net/http"

	"github.com/flnow/flnow/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	fmt.Println("API modules starting...")
}

// Run API listen
func Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/healthcheck", routes.HealthCheck)
	e.GET("/rts", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, e.Routes())
	})
	e.GET("/", hello)

	authGroup := e.Group("")
	// TODO: Auth Module
	// authGroup.Use(middleware.JWT([]byte("secret")))
	// authGroup.GET("/auth_hello", hello)
	// authGroup.GET("/test", func(ctx echo.Context) error {
	// 	user := ctx.Get("user").(*jwt.Token)
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	name := claims["name"].(string)
	// 	return ctx.String(http.StatusOK, "Welcome "+name+"!")
	// })

	//Flow biz
	flowGroup := authGroup.Group("/flows")

	flowGroup.POST("/create", nil)

	// nodeGroup := authGroup.Group("/nodes")
	// pluginGroup := authGroup.Group("/plugins")
	// systemGroup := authGroup.Group("/sys")

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
