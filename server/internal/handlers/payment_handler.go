package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService}
}

func (h *PaymentHandler) HandlePaymentNotification(c *gin.Context) {
	var notif dto.MidtransNotificationRequest
	if !utils.BindAndValidateJSON(c, &notif) {
		return
	}

	if err := h.paymentService.HandlePaymentNotification(notif); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to process payment notification", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment succesfully"})
}

func (h *PaymentHandler) GetAllUserPayments(c *gin.Context) {
	var param dto.PaymentQueryParam

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

	result, pagination, err := h.paymentService.GetAllUserPayments(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": pagination,
	})
}
