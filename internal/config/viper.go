package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	if os.Getenv("APP_ENV") == "local" {
		config.SetConfigName("config")
	} else {
		config.SetConfigName("config-local")

	}

	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config

}
