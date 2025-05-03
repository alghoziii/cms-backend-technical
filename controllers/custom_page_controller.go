package controllers

import (
	"github.com/alghoziii/cms-backend-technical/domain/dto"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/alghoziii/cms-backend-technical/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CustomPageController struct {
	DB *gorm.DB
}

func NewCustomPageController(DB *gorm.DB) CustomPageController {
	return CustomPageController{DB}
}

// @Summary Buat Page baru
// @Tags pages
// @Param body body dto.CustomPageRequest true "Data Pages"
// @Success 201 {object} dto.CustomPageResponse
// @Router /pages [post]
func (cpc *CustomPageController) CreateCustomPage(ctx *gin.Context) {
	var payload *dto.CustomPageRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newPage := models.CustomPage{
		URL:     payload.URL,
		Content: payload.Content,
		UserID:  userId,
	}

	result := cpc.DB.Create(&newPage)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": newPage.ID})
}

// @Summary Update Pages
// @Tags pages
// @Param id path int true "Pages ID"
// @Param body body dto.CustomPageRequest true "Data Pages"
// @Success 200 {object} map[string]interface{}
// @Router /pages/{id} [put]
func (cpc *CustomPageController) UpdateCustomPage(ctx *gin.Context) {
	pageId := ctx.Param("id")
	var payload *dto.CustomPageRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var page models.CustomPage
	result := cpc.DB.First(&page, "id = ? AND user_id = ?", pageId, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Page not found or not owned by user"})
		return
	}

	page.URL = payload.URL
	page.Content = payload.Content
	cpc.DB.Save(&page)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// @Summary Daftar halaman Page
// @Tags pages
// @Success 200 {array} dto.CustomPageResponse
// @Router /pages [get]
func (cpc *CustomPageController) FindCustomPages(ctx *gin.Context) {
	var pages []models.CustomPage
	results := cpc.DB.Preload("User").Find(&pages)
	if results.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
		return
	}

	var response []dto.CustomPageResponse
	for _, page := range pages {
		response = append(response, dto.CustomPageResponse{
			ID:      page.ID,
			URL:     page.URL,
			Content: page.Content,
			UserID:  page.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// @Summary Detail halaman Page
// @Tags pages
// @Param id path int true "Page ID"
// @Success 200 {array} dto.CustomPageResponse
// @Router /pages/{id} [get]
func (cpc *CustomPageController) FindCustomPageById(ctx *gin.Context) {
	pageId := ctx.Param("id")

	var page models.CustomPage
	result := cpc.DB.Preload("User").First(&page, "id = ?", pageId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.CustomPageResponse{
		ID:      page.ID,
		URL:     page.URL,
		Content: page.Content,
		UserID:  page.UserID,
	})
}

// @Summary Hapus Pages
// @Tags pages
// @Param id path int true "Pages ID"
// @Success 200 {object} map[string]interface{}
// @Router /pages/{id} [delete]
func (cpc *CustomPageController) DeleteCustomPage(ctx *gin.Context) {
	pageId := ctx.Param("id")

	userId, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	result := cpc.DB.Unscoped().Delete(&models.CustomPage{}, "id = ? AND user_id = ?", pageId, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
