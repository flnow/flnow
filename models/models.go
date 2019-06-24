package models

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // For mssql supported
	_ "github.com/jinzhu/gorm/dialects/mysql"    // For mysql supported
	_ "github.com/jinzhu/gorm/dialects/postgres" // For postgres supported
	//_ "github.com/jinzhu/gorm/dialects/sqlite"   // For sqlite supported

	"github.com/flnow/server/utils"
)

var (
	// DatabaseType to descript database type
	DatabaseType string
	// DatabaseConnection to descript database connection string
	DatabaseConnection string
	// DatabaseEngine from ORM framework
	DatabaseEngine *gorm.DB
)

func init() {
	// DatabaseEngine.DB().SetMaxOpenConns(10)

	// Sync tables
	if err := engineInit(); err != nil {
		log.Fatal(err)
	}
}

func engineInit() (err error) {
	DatabaseEngine, err = gorm.Open(strings.ToLower(utils.FlnowConf.Database.Type), utils.FlnowConf.Database.Connection)
	DatabaseEngine.LogMode(true)

	//defer DatabaseEngine.Close()
	if err != nil {
		return err
	}

	if !DatabaseEngine.HasTable(&Flow{}) {
		DatabaseEngine.CreateTable(&Flow{})
	}

	return err
}
