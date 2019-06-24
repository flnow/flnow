package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/flnow/server/modules/flow"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // For mssql supported
	_ "github.com/jinzhu/gorm/dialects/mysql"    // For mysql supported
	_ "github.com/jinzhu/gorm/dialects/postgres" // For postgres supported

	"github.com/flnow/server/utils"
)

func init() {
	fmt.Println("API modules starting...")
}

// Run API listen
func Run() {
	//global database connection init

	db, err := gorm.Open(strings.ToLower(utils.FlnowConf.Database.Type), utils.FlnowConf.Database.Connection)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", hello)
	r.POST("/flows/create", flow.Create(db))

	r.Run(":8081")
}

// Handler
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
