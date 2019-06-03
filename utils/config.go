package utils

import (
	"github.com/spf13/viper"
)

// Configuration struct
type Configuration struct {
	Database struct {
		Type       string `json:"Type"`
		Connection string `json:"Connection"`
	} `json:"Database"`
}

var (
	// FlnowConf to descript the configurations of Flnow
	FlnowConf *Configuration
)

func init() {
	// environment variables first, configuration files second. no more way to do settings.
	readFromFile()
	readFromEnv()
}

// Read config from file
func readFromFile() (err error) {
	vFromFile := viper.New()
	vFromFile.SetConfigName("flnow")
	vFromFile.SetConfigType("yml")
	vFromFile.AddConfigPath("/etc/flnow/")
	vFromFile.AddConfigPath(".")
	vFromFile.AddConfigPath("./configs/")

	err = vFromFile.ReadInConfig()
	if err != nil {
		return err
	}
	FlnowConf = new(Configuration)
	FlnowConf.Database.Type = vFromFile.GetString("Database.Type")
	FlnowConf.Database.Connection = vFromFile.GetString("Database.Connection")
	return err
}

// Read config from environment variables
func readFromEnv() {
	vFromEnv := viper.New()
	vFromEnv.SetEnvPrefix("FLNOW")
	vFromEnv.AutomaticEnv()

	dbType := vFromEnv.GetString("DATABASE_TYPE")
	if dbType != "" {
		FlnowConf.Database.Type = dbType
	}
	dbConn := vFromEnv.GetString("DATABASE_CONNECTION")
	if dbConn != "" {
		FlnowConf.Database.Connection = dbConn
	}
}
