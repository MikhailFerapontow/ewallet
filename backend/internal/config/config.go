package config

import (
	"log"

	"github.com/spf13/viper"
)

func MustLoad() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Init Config failed with error: %s", err.Error())
	}
}
