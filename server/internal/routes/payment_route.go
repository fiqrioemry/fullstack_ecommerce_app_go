package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, h *handlers.PaymentHandler) {
	order := r.Group("/api/payments")
	order.POST("/notifications", h.HandlePaymentNotifications)
	order.GET("", middleware.AuthRequired(), middleware.RoleOnly("customer", "admin"), h.GetAllUserPayments)

}
