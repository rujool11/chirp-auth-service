package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-auth-service/internal/db"
)

func main() {
	// initialize DB connection
	db.InitDB()
	defer db.DB.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from chirp-auth-service",
		})
	})

	r.Run(":8001")
}
