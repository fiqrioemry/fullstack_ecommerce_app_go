package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	category := r.Group("/api/categories")
	category.GET("", h.GetAllCategories)

	admin := category.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateCategory)
	admin.PUT("/:id", h.UpdateCategory)
	admin.DELETE("/:id", h.DeleteCategory)
}
