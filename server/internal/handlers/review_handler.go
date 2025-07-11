package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService services.ReviewService
}

func NewReviewHandler(reviewService services.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	itemID := c.Param("itemID")
	userID := utils.MustGetUserID(c)

	var req dto.CreateReviewRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}
	var uploadedURL string

	if req.Image != nil && req.Image.Filename != "" {
		var err error
		uploadedURL, err = utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Image upload failed",
				"error":   err.Error(),
			})
			return
		}
	}

	req.ImageURL = uploadedURL

	if err := h.reviewService.CreateReview(userID, itemID, req); err != nil {
		utils.CleanupImageOnError(uploadedURL)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

func (h *ReviewHandler) GetProductReviews(c *gin.Context) {
	productID := c.Param("productID")

	var param dto.ReviewQueryParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query param"})
		return
	}

	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 10
	}

	result, pagination, err := h.reviewService.GetReviewsByProductID(productID, param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to get reviews", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": pagination,
	})
}
