package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{cartService}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	items, total, err := h.cartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "total": total})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.CartItemRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	if err := h.cartService.AddToCart(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	productID := c.Param("productId")

	var req dto.UpdateCartItemRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	if err := h.cartService.UpdateQuantity(userID, productID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quantity updated"})
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	productID := c.Param("productId")
	if err := h.cartService.RemoveItem(userID, productID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	if err := h.cartService.ClearCart(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}

func (h *CartHandler) ToggleChecked(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	productID := c.Param("productId")

	if err := h.cartService.ToggleItemChecked(userID, productID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item check status updated"})
}
