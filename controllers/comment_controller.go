package controllers

import (
	"net/http"
	"strconv"

	"github.com/alghoziii/cms-backend-technical/domain/dto"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(DB *gorm.DB) CommentController {
	return CommentController{DB}
}

// @Summary Tambah komentar
// @Tags comments
// @Param id path int true "News ID"
// @Param body body dto.CommentRequest true "Komentar"
// @Success 201
// @Router /news/{id}/comments [post]
func (cc *CommentController) CreateComment(ctx *gin.Context) {
	newsIdStr := ctx.Param("id")

	// Convert string ke uint
	newsIdUint, err := strconv.ParseUint(newsIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	var payload *dto.CommentRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment := models.Comment{
		Name:    payload.Name,
		Comment: payload.Comment,
		NewsID:  uint(newsIdUint),
	}

	result := cc.DB.Create(&newComment)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}


