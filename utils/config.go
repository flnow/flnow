package config

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
	FlnowConf Configuration
)

func init() {
	// environment variables first, configuration files second. no more way to do settings.

}

// Read config from file
func readFromFile() {
	viper.SetConfigName("flnow")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/flnow/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
}

// Read config from environment variables
func readFromEnv() {

}
