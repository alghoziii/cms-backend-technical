package controllers

import (
	"net/http"

	"github.com/alghoziii/cms-backend-technical/domain/dto"
	"github.com/alghoziii/cms-backend-technical/domain/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(DB *gorm.DB) CategoryController {
	return CategoryController{DB}
}

// @Summary Buat kategori baru
// @Tags categories
// @Param body body dto.CategoryRequest true "Data kategori"
// @Success 201 {object} dto.CategoryResponse
// @Router /categories [post]
func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var payload *dto.CategoryRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory := models.Category{
		Name: payload.Name,
	}

	result := cc.DB.Create(&newCategory)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": newCategory.ID})
}

// @Summary Update kategori
// @Tags categories
// @Param id path int true "Category ID"
// @Param body body dto.CategoryRequest true "Data kategori"
// @Success 200 {object} map[string]interface{}
// @Router /categories/{id} [put]
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	var payload *dto.CategoryRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	result := cc.DB.First(&category, "id = ?", categoryId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	category.Name = payload.Name
	cc.DB.Save(&category)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// @Summary List kategori
// @Tags categories
// @Success 200 {array} dto.CategoryResponse
// @Router /categories [get]
func (cc *CategoryController) AllCategory(ctx *gin.Context) {
	var categories []models.Category
	results := cc.DB.Find(&categories)
	if results.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
		return
	}

	var response []dto.CategoryResponse
	for _, category := range categories {
		response = append(response, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// @Summary Detail kategori
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories/{id} [get]
func (cc *CategoryController) FindCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	var category models.Category
	result := cc.DB.First(&category, "id = ?", categoryId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

// @Summary Hapus kategori
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Router /categories/{id} [delete]
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	result := cc.DB.Unscoped().Delete(&models.Category{}, "id = ?", categoryId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
