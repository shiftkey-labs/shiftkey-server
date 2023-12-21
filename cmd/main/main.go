package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shiftkey-labs/shiftkey-api/pkg/controller"
)

func main() {

	router := gin.Default()

	router.GET("/users", controller.GetDummyUsers)

	// router.GET("/", "test")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
