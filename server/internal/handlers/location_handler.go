package handlers

import (
	"net/http"
	"strconv"

	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service services.LocationService
}

func NewLocationHandler(service services.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (h *LocationHandler) SearchProvincesByName(c *gin.Context) {
	var query dto.SearchProvinceRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid query"})
		return
	}

	data, err := h.service.SearchProvincesByName(query.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search provinces"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetProvinces(c *gin.Context) {
	data, err := h.service.GetAllProvinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch provinces"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetCitiesByProvinceID(c *gin.Context) {
	idStr := c.Param("provinceId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID"})
		return
	}

	data, err := h.service.GetCitiesByProvinceID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cities"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) SearchCitiesByName(c *gin.Context) {
	var query dto.SearchCityRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid query"})
		return
	}

	data, err := h.service.SearchCitiesByName(query.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search cities"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetDistrictsByCityID(c *gin.Context) {
	idStr := c.Param("cityId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	data, err := h.service.GetDistrictsByCityID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch districts"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetSubdistrictsByDistrictID(c *gin.Context) {
	idStr := c.Param("districtId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid district ID"})
		return
	}

	data, err := h.service.GetSubdistrictsByDistrictID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subdistricts"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetPostalCodesBySubdistrictID(c *gin.Context) {
	idStr := c.Param("subdistrictId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subdistrict ID"})
		return
	}

	data, err := h.service.GetPostalCodesBySubdistrictID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch postal codes"})
		return
	}
	c.JSON(http.StatusOK, data)
}
