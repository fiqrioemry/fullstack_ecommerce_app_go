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
	orderID := c.Param("orderID")
	productID := c.Param("productID")
	userID := utils.MustGetUserID(c)

	var req dto.CreateReviewRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	err := h.reviewService.CreateReview(orderID, userID, productID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

func (h *ReviewHandler) GetProductReviews(c *gin.Context) {
	productID := c.Param("productID")

	result, err := h.reviewService.GetReviewsByProductID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to get reviews"})
		return
	}

	c.JSON(http.StatusOK, result)
}
