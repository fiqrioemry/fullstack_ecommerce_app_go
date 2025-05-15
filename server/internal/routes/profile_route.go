package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, handler *handlers.ProfileHandler) {
	user := r.Group("/api/user")
	user.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	user.GET("/profile", handler.GetProfile)
	user.PUT("/profile", handler.UpdateProfile)
	user.PATCH("/profile/avatar", handler.UpdateAvatar)
}
