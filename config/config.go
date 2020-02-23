package config

import (
	"sync"
	"github.com/spf13/viper"
)

var mu sync.Mutex

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	LoadDB()
}