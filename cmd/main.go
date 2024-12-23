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
	userHandler, authMiddleware := factory.InitUserModule(db)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/property", propertyHandler.GetAll)
	router.POST("/auth/signup", userHandler.CreateUser)
	router.POST("/auth/login", userHandler.Login)
	router.GET("/user/profile", authMiddleware.CheckAuth(), userHandler.GetUserProfile)
	err := router.Run(cfg.ServerPort)
	if err != nil {
		log.Fatalf("fail to start server: %v", err)
	}
}
