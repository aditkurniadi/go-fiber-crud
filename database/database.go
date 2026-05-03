package database

import (
	"fmt"
	"log"

	"go-fiber-crud/config"
	"go-fiber-crud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	config.LoadEnv()

	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "user_admin")
	password := config.GetEnv("DB_PASSWORD", "password123")
	dbName := config.GetEnv("DB_NAME", "crud_db")
	sslMode := config.GetEnv("DB_SSLMODE", "disable")
	timeZone := config.GetEnv("DB_TIMEZONE", "Asia/Jakarta")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslMode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	return db
}