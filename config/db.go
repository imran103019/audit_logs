package config

import (
	"time"
	"fmt"
	"github.com/spf13/viper"
)

// Database holds the database configuration
type Database struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

var db Database

// DB returns the default database configuration
func DB() Database {
	return db
}

// LoadDB loads database configuration
func LoadDB() {
	mu.Lock()
	defer mu.Unlock()
	db = Database{
		Name:            viper.GetString("DB_DATABASE"),
		User:            viper.GetString("DB_USERNAME"),
		Password:        viper.GetString("DB_PASSWORD"),
		Host:            fmt.Sprintf("%s:%d", viper.GetString("DB_HOST"), viper.GetInt("DB_PORT")),
		Port:            viper.GetInt("DB_PORT"),
		MaxIdleConn:     10,
		MaxOpenConn:     20,
		MaxConnLifetime: 10000,
		Debug:           viper.GetBool("DB_DEBUG"),
	}
}
