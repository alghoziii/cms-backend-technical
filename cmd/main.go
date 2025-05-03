package main

import (
	"log"

	"github.com/alghoziii/cms-backend-technical/config"
	_ "github.com/alghoziii/cms-backend-technical/docs"
	"github.com/alghoziii/cms-backend-technical/routes"
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


	router := routes.SetupRouter(db)


	log.Printf("Server running on port %s", cfg.AppPort)
	log.Fatal(router.Run(":" + cfg.AppPort))
}
