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

// Regions

func (h *FiltersHandler) GetRegions(c *gin.Context) {
	regions, err := h.filtersService.GetRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": regions})
}

func (h *FiltersHandler) AddRegion(c *gin.Context) {
	var input domain.FilterRegion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if region, err := h.filtersService.AddRegion(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": region})
	}
}

func (h *FiltersHandler) UpdateRegion(c *gin.Context) {
	regionID := c.Param("id")
	id, err := strconv.Atoi(regionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var input domain.FilterRegion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if region, err := h.filtersService.UpdateRegion(id, input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": region})
	}
}

func (h *FiltersHandler) DeleteRegion(c *gin.Context) {
	regionID := c.Param("id")
	id, err := strconv.Atoi(regionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.filtersService.DeleteRegion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// PropertyCategories

func (h *FiltersHandler) GetPropertyCategories(c *gin.Context) {
	categories, err := h.filtersService.GetPropertyCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (h *FiltersHandler) AddPropertyCategory(c *gin.Context) {
	var input domain.FilterPropertyCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if category, err := h.filtersService.AddPropertyCategory(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": category})
	}
}

func (h *FiltersHandler) UpdatePropertyCategory(c *gin.Context) {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var input domain.FilterPropertyCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if category, err := h.filtersService.UpdatePropertyCategory(id, input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": category})
	}
}

func (h *FiltersHandler) DeletePropertyCategory(c *gin.Context) {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.filtersService.DeletePropertyCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
