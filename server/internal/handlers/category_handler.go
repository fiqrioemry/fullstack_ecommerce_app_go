package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	var param dto.CategoryQueryParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query param"})
		return
	}
	if param.Page == 0 {
		param.Page = 1
	}
	if param.Limit == 0 {
		param.Limit = 10
	}

	result, pagination, err := h.categoryService.GetAllCategories(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch categories", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": pagination,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	uploadedURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
		return
	}
	req.ImageURL = uploadedURL

	if err := h.categoryService.CreateCategory(req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var req dto.UpdateCategoryRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.categoryService.UpdateCategory(categoryID, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update category", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := h.categoryService.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
