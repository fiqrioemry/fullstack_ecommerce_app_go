package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, h *handlers.OrderHandler) {

	order := r.Group("/api/orders", middleware.AuthRequired())
	order.POST("/check-shipping", h.CheckShippingCost)

	order.POST("", middleware.RoleOnly("customer"), h.Checkout)
	order.GET("", middleware.RoleOnly("admin", "customer"), h.GetAllUserOrders)
	order.GET("/:orderID", middleware.RoleOnly("admin", "customer"), h.GetOrderDetail)

	order.POST("/:orderID/shipment", middleware.RoleOnly("admin"), h.CreateShipment)
	order.PUT("/:orderID/shipment", middleware.RoleOnly("admin"), h.UpdateShipmentStatus)
	order.GET("/:orderID/shipment", middleware.RoleOnly("admin", "customer"), h.GetShipmentInfo)
}
