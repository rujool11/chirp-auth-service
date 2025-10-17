package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-auth-service/internal/db"
	"github.com/rujool11/chirp-auth-service/internal/models"
	"github.com/rujool11/chirp-auth-service/internal/utils"
)

func RegisterUser(c *gin.Context) {

	// binding: required ensures that the field must be present
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=6"` // min 6 chars
	}

	// convert input JSON to struct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind to JSON"})
		return
	}

	// hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// user creation query, returns record
	query := `INSERT INTO USERS (username, email, password_hash, created_at)
			VALUES ($1, $2, $3, $4) RETURNING id, username, email, bio, likes_count, 
			followers_count, following_count, tweets_count, created_at`

	var user models.User
	// query parameters passed to
	// QueryRow - returns one row, Query - returns multiple rows, Exec - doesnt return anything (eg. simple insert)
	// create row, and save returned values to user struct
	err = db.DB.QueryRow(c, query,
		input.Username,
		input.Email,
		hashedPassword,
		time.Now(),
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Bio,
		&user.LikesCount,
		&user.FollowersCount,
		&user.FollowingCount,
		&user.TweetsCount,
		&user.CreatedAt,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// generate token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// retrurn user and token in response
	c.JSON(200, gin.H{"user": user, "token": token})

}

func LoginUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
