package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"moydom_api/config"
	"moydom_api/internal/factory"
	"moydom_api/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg.DatabaseURL)

	propertyHandler := factory.InitPropertyModule(db)
	router := gin.Default()
	router.GET("/property", propertyHandler.GetAll)

	err := router.Run(cfg.ServerPort)
	if err != nil {
		log.Fatalf("fail to start server: %v", err)
	}
}
