package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-auth-service/internal/db"
)

func main() {
	// initialize DB connection
	db.InitDB()
	db.CreateUserTableIfDoesNotExist()
	defer db.DB.Close()

	// gin.Default() is the router, and gin.Context() is the current HTTP req and res
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from chirp-auth-service",
		})
	})

	r.Run(":8001")
}
