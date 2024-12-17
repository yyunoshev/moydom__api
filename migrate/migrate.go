package main

import (
	"log"
	"moydom_api/initializers"
	"moydom_api/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Client{})
	if err != nil {
		log.Fatal(err)
	}
}
