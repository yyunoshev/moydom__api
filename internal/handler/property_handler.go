package handler

import (
	"github.com/gin-gonic/gin"
	"moydom_api/internal/service"
	"net/http"
)

type PropertyHandler struct {
	propertyService *service.PropertyService
}

func NewPropertyHandler(propertyService *service.PropertyService) *PropertyHandler {
	return &PropertyHandler{propertyService: propertyService}
}

func (h *PropertyHandler) GetAll(c *gin.Context) {
	page, pageSize := getPaginationParams(c)
	properties, total, err := h.propertyService.GetAllProperties(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, paginatedResponse{
		Data:       properties,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	})
}
