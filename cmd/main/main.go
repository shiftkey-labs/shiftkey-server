package main

import (
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
	db.Migrate()
	db.Seed(db.DB)

	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {
		users, err := controller.GetAllUsers()
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
