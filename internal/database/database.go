package database

import (
	"golang-api-module/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(config *config.Config) *gorm.DB {
	var err error

	DB, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Database is not connected")
	}

	log.Println("Database is succesfully connected")

	sqlDB, err := DB.DB()

	if err != nil {
		log.Fatal("Failed to get generic DB object:", err)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return DB
}
