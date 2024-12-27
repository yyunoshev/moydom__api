package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"moydom_api/config"
	_ "moydom_api/docs"
	"moydom_api/internal/factory"
	"moydom_api/pkg/database"
)

// @title           Мой дом
// @version         1.0
// @description     API для сервиса Мой дом

// @host      localhost:8080
// @BasePath  /
func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg.DatabaseURL)

	propertyHandler := factory.InitPropertyModule(db)
	filtersHandler := factory.InitFiltersModule(db)
	userHandler, authMiddleware := factory.InitUserModule(db)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/property", propertyHandler.GetAll)
	router.POST("/auth/signup", userHandler.CreateUser)
	router.POST("/auth/login", userHandler.Login)
	router.GET("/user/profile", authMiddleware.CheckAuth(), userHandler.GetUserProfile)

	router.GET("/admin/districts", authMiddleware.CheckAuth(), filtersHandler.GetDistricts)
	router.POST("/admin/districts", authMiddleware.CheckAuth(), filtersHandler.AddDistrict)
	router.PATCH("/admin/districts/:id", authMiddleware.CheckAuth(), filtersHandler.UpdateDistrict)
	router.DELETE("/admin/districts/:id", authMiddleware.CheckAuth(), filtersHandler.DeleteDistrict)

	router.GET("/admin/microdistricts", authMiddleware.CheckAuth(), filtersHandler.GetMicrodistricts)
	router.POST("/admin/microdistricts", authMiddleware.CheckAuth(), filtersHandler.AddMicrodistrict)
	router.PATCH("/admin/microdistricts/:id", authMiddleware.CheckAuth(), filtersHandler.UpdateMicrodistrict)
	router.DELETE("/admin/microdistricts/:id", authMiddleware.CheckAuth(), filtersHandler.DeleteMicrodistrict)

	router.GET("/admin/regions", authMiddleware.CheckAuth(), filtersHandler.GetRegions)
	router.POST("/admin/regions", authMiddleware.CheckAuth(), filtersHandler.AddRegion)
	router.PATCH("/admin/regions/:id", authMiddleware.CheckAuth(), filtersHandler.UpdateRegion)
	router.DELETE("/admin/regions/:id", authMiddleware.CheckAuth(), filtersHandler.DeleteRegion)

	err := router.Run(cfg.ServerPort)
	if err != nil {
		log.Fatalf("fail to start server: %v", err)
	}
}
