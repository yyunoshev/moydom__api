package handler

import (
	"github.com/gin-gonic/gin"
	"moydom_api/internal/domain"
	"moydom_api/internal/service"
	"net/http"
	"strconv"
)

type FiltersHandler struct {
	filtersService *service.FiltersService
}

func NewFiltersHandler(filtersService *service.FiltersService) *FiltersHandler {
	return &FiltersHandler{filtersService: filtersService}
}

// Districts

func (h *FiltersHandler) GetDistricts(c *gin.Context) {
	districts, err := h.filtersService.GetDistricts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": districts})
}

func (h *FiltersHandler) AddDistrict(c *gin.Context) {
	var input domain.FilterDistrict
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if district, err := h.filtersService.AddDistrict(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": district})
	}
}

func (h *FiltersHandler) UpdateDistrict(c *gin.Context) {
	districtID := c.Param("id")
	id, err := strconv.Atoi(districtID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var input domain.FilterDistrict
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if district, err := h.filtersService.UpdateDistrict(id, input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": district})
	}
}

func (h *FiltersHandler) DeleteDistrict(c *gin.Context) {
	districtID := c.Param("id")
	id, err := strconv.Atoi(districtID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.filtersService.DeleteDistrict(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// MicroDistricts

func (h *FiltersHandler) GetMicrodistricts(c *gin.Context) {
	microdistricts, err := h.filtersService.GetMicrodistricts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": microdistricts})
}

func (h *FiltersHandler) AddMicrodistrict(c *gin.Context) {
	var input domain.FilterMicrodistrict
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if microdistrict, err := h.filtersService.AddMicrodistrict(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": microdistrict})
	}
}

func (h *FiltersHandler) UpdateMicrodistrict(c *gin.Context) {
	microdistrictID := c.Param("id")
	id, err := strconv.Atoi(microdistrictID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var input domain.FilterMicrodistrict
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if microdistrict, err := h.filtersService.UpdateMicrodistrict(id, input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": microdistrict})
	}
}

func (h *FiltersHandler) DeleteMicrodistrict(c *gin.Context) {
	microdistrictID := c.Param("id")
	id, err := strconv.Atoi(microdistrictID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.filtersService.DeleteMicrodistrict(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
