package conn

import (
	"github.com/imran103019/audit_logs/config"
	"github.com/jinzhu/gorm"
	"fmt"
	// mysql dialects need to be imported separately
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB holds the database instace
type DB struct{ *gorm.DB }

// defaultConfig is the default database instance
var defaultDB = &DB{}

// Connect sets the db client of database using configuration cfg
func (db *DB) Connect(cfg *config.Database) error {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&timeout=120s", cfg.User, cfg.Password, cfg.Host, cfg.Name)
	d, err := gorm.Open("mysql", uri)

	if err != nil {
		fmt.Print(err)
		return err
	}
	db.DB = d

	if cfg.MaxIdleConn != 0 {
		db.DB.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	}
	if cfg.MaxOpenConn != 0 {
		db.DB.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	}
	if cfg.MaxConnLifetime.Seconds() != 0 {
		db.DB.DB().SetConnMaxLifetime(cfg.MaxConnLifetime)
	}

	db.DB = d.LogMode(cfg.Debug).Debug()
	db.DB.SingularTable(true)

	return nil
}

// Default returns default db
func Default() *DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration file
func Connect() error {
	cfg := config.DB()
	return defaultDB.Connect(&cfg)
}

func Ping() error {
	return defaultDB.DB.DB().Ping()
}

