package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handlers.ProductHandler) {
	product := r.Group("/api/product")
	product.GET("", h.SearchProducts)
	product.GET("/:slug", h.GetProductBySlug)

	admin := product.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateProduct)
	admin.PUT("/:id", h.UpdateProduct)
	admin.DELETE("/:id", h.DeleteProduct)
}
