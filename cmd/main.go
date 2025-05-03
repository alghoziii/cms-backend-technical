package main

import (
	"github.com/alghoziii/cms-backend-technical/config"
	_ "github.com/alghoziii/cms-backend-technical/docs"
	"github.com/alghoziii/cms-backend-technical/routes"
	"log"
)

// @title CMS API
// @version 1.0
// @description API untuk Content Management System
// @host localhost:8080
// @BasePath /
func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize DB
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := config.SeedDB(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Setup routes
	router := routes.SetupRouter(db)

	// Start server
	log.Printf("Server running on port %s", cfg.AppPort)
	log.Fatal(router.Run(":" + cfg.AppPort))
}
