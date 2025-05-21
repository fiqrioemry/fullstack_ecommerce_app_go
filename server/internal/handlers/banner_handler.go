package handlers

import (
	"net/http"

	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BannerHandler struct {
	bannerService services.BannerService
}

func NewBannerHandler(bannerService services.BannerService) *BannerHandler {
	return &BannerHandler{bannerService}
}

func (h *BannerHandler) UploadBanner(c *gin.Context) {
	var req dto.BannerRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	uploadedURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
		return
	}
	req.ImageURL = uploadedURL

	if err := h.bannerService.Create(req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create banner", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "banner uploaded successfully"})
}

func (h *BannerHandler) GetAllBanners(c *gin.Context) {
	results, err := h.bannerService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch banner"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *BannerHandler) GetBanner(c *gin.Context) {
	position := c.Param("position")
	results, err := h.bannerService.Get(position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch banner"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.bannerService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete banner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "banner deleted"})
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	bannerID := c.Param("id")
	var req dto.BannerRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.bannerService.Update(uuid.MustParse(bannerID), req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update banner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "banner updated"})
}
