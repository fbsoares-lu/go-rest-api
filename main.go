package main

import (
	"fmt"
	"rest-go-api/database"
	"rest-go-api/models"
	"rest-go-api/routes"
)

func main() {
	database.Connection()
	fmt.Println("Server is Running...")
	routes.HandleRequest()

	database.DB.AutoMigrate(&models.Personality{})
}
