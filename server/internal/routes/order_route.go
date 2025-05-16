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
	order.POST("/:orderID/confirm-delivery", middleware.RoleOnly("customer"), h.ConfirmOrderDelivered)

	order.GET("", middleware.RoleOnly("admin", "customer"), h.GetAllUserOrders)
	order.GET("/:orderID", middleware.RoleOnly("admin", "customer"), h.GetOrderDetail)
	order.GET("/:orderID/shipment", middleware.RoleOnly("admin", "customer"), h.GetShipmentByOrderID)
	order.PATCH("/orders/:orderID/cancel", middleware.RoleOnly("admin", "customer"), h.CancelOrder)

	order.POST("/:orderID/shipment", middleware.RoleOnly("admin"), h.CreateShipment)
}
