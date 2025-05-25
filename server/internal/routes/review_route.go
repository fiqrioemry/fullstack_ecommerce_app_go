package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine, h *handlers.ReviewHandler) {
	review := r.Group("/api/reviews")
	review.GET("/:productID", h.GetProductReviews)

	auth := review.Use(middleware.AuthRequired())
	auth.POST("/order/:productID", h.CreateReview)
}
