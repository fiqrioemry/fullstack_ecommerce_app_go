// internal/routes/user_route.go
package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine, handler *handlers.AdminHandler) {
	admin := r.Group("/api/admin/dashboard")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))

	admin.GET("/customers", handler.GetAllCustomer)
	admin.GET("/customers/:id", handler.GetCustomerDetail)
	admin.GET("/summary", handler.GetDashboardStats)
	admin.GET("/revenue", handler.GetRevenueStats)

}
