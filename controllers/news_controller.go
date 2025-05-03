package controllers

import (
	"github.com/alghoziii/cms-backend-technical/domain/dto"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/alghoziii/cms-backend-technical/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type NewsController struct {
	DB *gorm.DB
}

func NewNewsController(DB *gorm.DB) NewsController {
	return NewsController{DB}
}

// @Summary Buat berita baru
// @Tags news
// @Param body body dto.NewsRequest true "Data berita"
// @Success 201 {object} dto.NewsResponse
// @Router /news [post]
func (nc *NewsController) CreateNews(ctx *gin.Context) {
	var payload *dto.NewsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newNews := models.News{
		Title:      payload.Title,
		Content:    payload.Content,
		CategoryID: payload.CategoryID,
		UserID:     userId,
	}

	result := nc.DB.Create(&newNews)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": newNews.ID})
}

// @Summary Buat berita baru
// @Tags news
// @Param body body dto.NewsRequest true "Data berita"
// @Success 200 {object} dto.NewsResponse
// @Router /news/{id} [put]
func (nc *NewsController) UpdateNews(ctx *gin.Context) {
	newsId := ctx.Param("id")
	var payload *dto.NewsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var news models.News
	result := nc.DB.First(&news, "id = ? AND user_id = ?", newsId, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "News not found or not owned by user"})
		return
	}

	news.Title = payload.Title
	news.Content = payload.Content
	news.CategoryID = payload.CategoryID
	nc.DB.Save(&news)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// @Summary List semua berita
// @Tags news
// @Success 200 {array} dto.NewsResponse
// @Router /news [get]
func (nc *NewsController) FindNews(ctx *gin.Context) {
	var news []models.News
	results := nc.DB.Preload("Category").Preload("User").Find(&news)
	if results.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
		return
	}

	var response []dto.NewsResponse
	for _, item := range news {
		response = append(response, dto.NewsResponse{
			ID:         item.ID,
			Title:      item.Title,
			Content:    item.Content,
			CategoryID: item.CategoryID,
			UserID:     item.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// @Summary Detail Berita
// @Tags news
// @Param id path int true "Data Berita"
// @Success 200 {object} dto.NewsResponse
// @Router /news/{id} [get]
func (nc *NewsController) FindNewsById(ctx *gin.Context) {
	newsId := ctx.Param("id")

	var news models.News
	result := nc.DB.Preload("Category").Preload("User").First(&news, "id = ?", newsId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.NewsResponse{
		ID:         news.ID,
		Title:      news.Title,
		Content:    news.Content,
		CategoryID: news.CategoryID,
		UserID:     news.UserID,
	})
}

// @Summary Hapus Berita
// @Tags news
// @Param id path int true "Data Berita"
// @Success 200 {object} map[string]interface{}
// @Router /news/{id} [delete]
func (nc *NewsController) DeleteNews(ctx *gin.Context) {
	newsId := ctx.Param("id")

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	result := nc.DB.Delete(&models.News{}, "id = ? AND user_id = ?", newsId, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
