package database

import (
	"fmt"

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

type DatabaseConnection interface {
	OpenConnection()
}

type PostgresDatabaseConnection struct {
	config *DbConfig
}

func (dbc *PostgresDatabaseConnection) OpenConnection() (*gorm.DB, error) {
	// Build connection string
	dsn := fmt.Sprintf(
		"host=%[1]s user=%[2]s password=%[3]s dbname=%[4]s port=%[5]s sslmode=%[6]s TimeZone=%[7]s", 
		dbc.config.Host, dbc.config.User, dbc.config.Password, dbc.config.DbName, dbc.config.Port, dbc.config.SslMode, dbc.config.TimeZone)

	// Open connection
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Report error
	if err != nil {
		return nil, err
	}

	// Get config object
	sqlDB, err := db.DB()

	// Report error
	if err != nil {
		return nil, err
	}

	// Configure connection pooling
    sqlDB.SetMaxOpenConns(100) // Maximum number of open connections
    sqlDB.SetMaxIdleConns(10)  // Maximum number of idle connections
    sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection can be reused (0 means no limit)

	// Return db object
	return db, nil
}

func NewPostgresDatabaseConnection() *PostgresDatabaseConnection {
	// TODO: Load from .env file
	config := DbConfig{
		Host:     "localhost",
		User:     "admin",
		Password: "root",
		DbName:   "wishlist_db",
		Port:     "5432",
		SslMode:  "disable",
		TimeZone: "GMT",
	}

	return &PostgresDatabaseConnection{
		config: &config,
	}
}
