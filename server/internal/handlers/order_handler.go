package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.CheckoutRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	resp, err := h.service.Checkout(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) GetAllUserOrders(c *gin.Context) {
	var param dto.OrderQueryParam
	role := utils.MustGetRole(c)
	userID := utils.MustGetUserID(c)

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query param"})
		return
	}

	if param.Page == 0 {
		param.Page = 1
	}
	if param.Limit == 0 {
		param.Limit = 10
	}

	result, pagination, err := h.service.GetAllOrders(userID, role, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": pagination,
	})
}

func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("orderID")
	order, err := h.service.GetOrderDetail(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateShipment(c *gin.Context) {
	orderID := c.Param("orderID")
	var req dto.CreateShipmentRequest

	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	result, err := h.service.CreateShipment(orderID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *OrderHandler) GetShipmentByOrderID(c *gin.Context) {
	orderID := c.Param("orderID")

	result, err := h.service.GetShipmentByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Shipment not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *OrderHandler) ConfirmOrderDelivered(c *gin.Context) {
	orderID := c.Param("orderID")
	userID := utils.MustGetUserID(c)

	result, err := h.service.ConfirmOrderDelivered(orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *OrderHandler) CheckShippingCost(c *gin.Context) {
	var req dto.ShippingCostRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	const originCityID = 52
	const originProvinceID = 2

	costs, err := utils.EstimateShippingRates(
		originProvinceID,
		originCityID,
		req.DestinationProvinceID,
		req.DestinationCityID,
		req.Weight,
		req.Courier,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to calculate shipping cost",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courier": req.Courier,
		"costs":   costs,
	})
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	orderID := c.Param("orderID")

	resp, err := h.service.CancelOrder(orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
