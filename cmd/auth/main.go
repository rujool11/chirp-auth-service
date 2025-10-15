package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rujool11/chirp-auth-service/internal/db"
)

func main() {
	// initialize DB connection
	db.InitDB()
	db.CreateUserTableIfDoesNotExist()
	defer db.DB.Close()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	// gin.Default() is the router, and gin.Context() is the current HTTP req and res
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from chirp-auth-service",
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8001"
		log.Println("Defaulting PORT to 8001")
	}

	r.Run(":" + PORT)
}
