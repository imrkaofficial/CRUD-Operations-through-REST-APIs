package main

import (
	"mysql-api/config"
	"mysql-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "demo:demo1234@tcp(localhost:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	config.Connect(dsn) // Initialize and connect to the database

	router := gin.New() // Create a Gin router

	routes.AccRoute(router) // Initialize routes

	if err := router.Run(":3000"); err != nil {
		// Handle any errors that occur when starting the server
		panic(err)
	}
}
