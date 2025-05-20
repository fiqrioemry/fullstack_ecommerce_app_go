package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/send-otp", handler.SendOTP)
		auth.POST("/register", handler.Register)
		auth.POST("/verify-otp", handler.VerifyOTP)
		auth.POST("/refresh-token", handler.RefreshToken)

		auth.GET("/google", handler.GoogleOAuthRedirect)
		auth.GET("/google/callback", handler.GoogleOAuthCallback)

		protected := auth.Group("")
		protected.Use(middleware.AuthRequired(), middleware.RoleOnly("admin", "customer"))
		protected.GET("/me", handler.AuthMe)
	}
}
