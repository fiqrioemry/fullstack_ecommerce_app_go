package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AddressRoutes(r *gin.Engine, handler *handlers.AddressHandler) {
	address := r.Group("/api/user")
	address.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	address.GET("/addresses", handler.GetAddresses)
	address.POST("/addresses", handler.AddAddress)
	address.PUT("/addresses/:id", handler.UpdateAddress)
	address.DELETE("/addresses/:id", handler.DeleteAddress)
	address.PATCH("/addresses/:id/main", handler.SetMainAddress)
}
