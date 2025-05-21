// internal/handlers/user_handler.go
package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *AdminHandler {
	return &AdminHandler{adminService}
}

func (h *AdminHandler) GetAllCustomer(c *gin.Context) {
	var params dto.CustomerQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	customers, pagination, err := h.adminService.GetAllCustomer(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch customers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       customers,
		"pagination": pagination,
	})
}

func (h *AdminHandler) GetCustomerDetail(c *gin.Context) {
	id := c.Param("id")
	user, err := h.adminService.GetCustomerDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *AdminHandler) GetDashboardStats(c *gin.Context) {
	gender := c.Query("gender")
	data, err := h.adminService.GetDashboardStats(gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get dashboard stats"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *AdminHandler) GetRevenueStats(c *gin.Context) {
	rangeType := c.DefaultQuery("range", "daily")
	stats, total, err := h.adminService.GetRevenueStats(rangeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get revenue stats"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"stats": stats,
	})
}
