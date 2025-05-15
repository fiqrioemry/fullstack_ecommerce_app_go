package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, h *handlers.CartHandler) {
	cart := r.Group("/api/cart", middleware.AuthRequired(), middleware.RoleOnly("customer"))

	cart.GET("", h.GetCart)
	cart.POST("", h.AddToCart)
	cart.DELETE("", h.ClearCart)
	cart.DELETE("/:productId", h.RemoveItem)
	cart.PUT("/:productId", h.UpdateQuantity)
	cart.PATCH("/:productId/checked", h.ToggleChecked)

}
