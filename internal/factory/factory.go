package factory

import (
	"gorm.io/gorm"
	"moydom_api/internal/handler"
	"moydom_api/internal/middleware"
	"moydom_api/internal/repository"
	"moydom_api/internal/service"
)

func InitPropertyModule(db *gorm.DB) *handler.PropertyHandler {
	propertyRepo := repository.NewPropertyRepository(db)
	propertyService := service.NewPropertyService(propertyRepo)
	return handler.NewPropertyHandler(propertyService)
}

func InitUserModule(db *gorm.DB) (*handler.UserHandler, *middleware.AuthMiddleware) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	return handler.NewUserHandler(userService), middleware.NewAuthMiddleware(*userService)
}

func InitFiltersModule(db *gorm.DB) *handler.FiltersHandler {
	filtersRepo := repository.NewFiltersRepository(db)
	filtersService := service.NewFiltersService(filtersRepo)
	return handler.NewFiltersHandler(filtersService)
}
