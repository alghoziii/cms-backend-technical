package config

import (
	"os"

	"github.com/alghoziii/cms-backend-technical/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var Cfg *Config 

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	config := &Config{
		AppPort:    appPort,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "cms_db"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}

	Cfg = config 
	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func SeedDB(db *gorm.DB) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	users := []models.User{
		{Username: "admin", Password: string(hashedPassword)},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error; err != nil {
			return err
		}
	}
	return nil
}
