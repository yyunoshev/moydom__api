package factory

import (
	"gorm.io/gorm"
	"moydom_api/internal/handler"
	"moydom_api/internal/repository"
	"moydom_api/internal/service"
)

func InitPropertyModule(db *gorm.DB) *handler.PropertyHandler {
	propertyRepo := repository.NewPropertyRepository(db)
	propertyService := service.NewPropertyService(propertyRepo)
	return handler.NewPropertyHandler(propertyService)

}
