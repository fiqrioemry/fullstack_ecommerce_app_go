package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService}
}

func (h *PaymentHandler) HandlePaymentNotifications(c *gin.Context) {
	h.paymentService.WebhookNotifications(c)
}

func (h *PaymentHandler) GetAllUserPayments(c *gin.Context) {
	var param dto.PaymentQueryParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query param"})
		return
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
