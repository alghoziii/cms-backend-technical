package config

import (
	"fmt"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Cfg.DBHost, Cfg.DBUser, Cfg.DBPassword, Cfg.DBName, Cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	log.Println("Database connected successfully")

	// Tambahkan log untuk debugging
	log.Println("Running database migrations...")

	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.News{},
		&models.CustomPage{},
		&models.Comment{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
	return db, nil
}
