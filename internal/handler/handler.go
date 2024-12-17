package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type paginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

func getPaginationParams(c *gin.Context) (page, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))           // Номер страницы по умолчанию 1
	pageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10")) // Размер страницы по умолчанию 10
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return page, pageSize
}
