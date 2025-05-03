package utils

import (
	"errors"
	"github.com/alghoziii/cms-backend-technical/config"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/gin-gonic/gin"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(config.Cfg.JWTSecret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Cfg.JWTSecret), nil
	})
}

func GetUserIDFromToken(ctx *gin.Context) (uint, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return 0, errors.New("authorization header missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := ValidateJWT(tokenString)
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token")
	}

	return uint(userIDFloat), nil
}
