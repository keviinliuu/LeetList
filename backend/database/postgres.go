package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host		string
	Port		string
	Password	string
	User 		string
	DBName		string
	SSLMode		string
}

func NewConnection(config *Config)(*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s dbname=%s user=%s sslmode=%s", 
		config.Host, config.Port, config.Password, config.User, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	fmt.Println("Successfully connected to database.")

	return db, nil
}