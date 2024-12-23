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

// GetAll godoc
// @Summary      Получить список недвижимости
// @Description  Возвращает полный список созданной недвижимости с пагинацией. Это не посты. Позже добавлю фильтры в виде параметров и загрузку постов.
// @Tags         property
// @Accept       json
// @Produce      json
// @Param        page       query     int     false  "Номер страницы"
// @Param        page_size  query     int     false  "Размер страницы"
// @Success      200        {object}  paginatedResponse{data=[]domain.Property}
// @Failure      400        {object}  map[string]string "Ошибка"
// @Router       /property [get]
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
