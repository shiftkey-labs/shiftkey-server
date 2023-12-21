package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shiftkey-labs/shiftkey-api/pkg/controller"
	"github.com/shiftkey-labs/shiftkey-api/pkg/db"
)

func main() {

	db.Connect()

	// Run these two if you're running the server for the first time or have to migrate the db
	// db.Migrate()
	// db.Seed(db.DB)

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", func(ctx *gin.Context) {
		users, err := controller.GetAllUsers()
		fmt.Print(users)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
