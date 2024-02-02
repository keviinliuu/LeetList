package database

import (
	"fmt"
	"log"

	"github.com/keviinliuu/leetlist/graph/model"
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
		config.Host, config.Port, config.Password, config.DBName, config.User, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	fmt.Println("Successfully connected to database.")

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Question{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}