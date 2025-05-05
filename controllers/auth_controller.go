package controllers

import (
	"net/http"

	domain "github.com/alghoziii/cms-backend-technical/domain/dto"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/alghoziii/cms-backend-technical/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// @Summary Login user
// @Tags login
// @Accept json
// @Param body body dto.LoginRequest true "Login data"
// @Success 200 {object} dto.LoginResponse
// @Router /login [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	var payload *domain.LoginRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "username = ?", payload.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, domain.LoginResponse{Token: token})
}
