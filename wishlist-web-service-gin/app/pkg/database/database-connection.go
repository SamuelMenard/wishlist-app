package database

import (
	"errors"
	"fmt"
	enums "smenard/wishlist-web-service-gin/app/pkg/database/emums"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
	TimeZone string
}

func ConnectToDb(strategy enums.DbConnectionStrategy, config DbConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch strategy {
	case enums.Postgre:
		db, err = connectPostgreDb(config)
	case enums.None:
		db = nil
		err = errors.New("unable to connect database, no strategy was provided")
	}

	return db, err
}

// Connect to a postgre database
func connectPostgreDb(config DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
			"host=%[1]s user=%[2]s password=%[3]s dbname=%[4]s port=%[5]s sslmode=%[6]s TimeZone=%[7]s", 
			config.Host, config.User, config.Password, config.DbName, config.Port, config.SslMode, config.TimeZone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
