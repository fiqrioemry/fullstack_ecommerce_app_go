package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BannerRoutes(r *gin.Engine, h *handlers.BannerHandler) {
	banner := r.Group("/api/banners")
	banner.GET("", h.GetAllBanners)
	banner.GET("/:position", h.GetBanner)

	admin := banner.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.UploadBanner)
	admin.PUT("/:id", h.UpdateBanner)
	admin.DELETE("/:id", h.DeleteBanner)
}
