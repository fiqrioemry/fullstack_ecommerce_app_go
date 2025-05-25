package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(ProductService services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService}
}

func extractUploadedImages(c *gin.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return nil, err
	}
	files := form.File["images"]
	if len(files) == 0 {
		return nil, nil
	}
	return utils.UploadMultipleImagesWithValidation(files)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")
	req.IsFeatured, _ = utils.ParseBoolFormField(c, "isFeatured")

	uploadedURLs, err := extractUploadedImages(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid image upload", "error": err.Error()})
		return
	}
	req.ImageURLs = uploadedURLs

	if err := h.ProductService.CreateProduct(req); err != nil {
		for _, img := range uploadedURLs {
			utils.CleanupImageOnError(img)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var req dto.UpdateProductRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")
	req.IsFeatured, _ = utils.ParseBoolFormField(c, "isFeatured")

	uploadedURLs, err := extractUploadedImages(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid image upload", "error": err.Error()})
		return
	}
	req.ImageURLs = uploadedURLs

	if err := h.ProductService.UpdateProduct(productID, req); err != nil {
		for _, img := range uploadedURLs {
			utils.CleanupImageOnError(img)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	if err := h.ProductService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	var params dto.GetAllProductsRequest
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	result, pagination, err := h.ProductService.SearchProducts(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to search products", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": pagination,
	})
}

func (h *ProductHandler) GetProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	product, err := h.ProductService.GetProductBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
