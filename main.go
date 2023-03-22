package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"Calorie-Tracking/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())	

	router.POST("/entries/create", routes.AddEntry)
	router.GET("/entries")
}
