package models

import (
	"strings"

	"github.com/jinzhu/gorm"
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
	DatabaseEngine, _ = gorm.Open(strings.ToLower(DatabaseType), DatabaseConnection)
	DatabaseEngine = DatabaseEngine
}
