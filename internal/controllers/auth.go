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
		Password string `json:"password" binding:"required"` // min 6 chars
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

	// input struct for incoming payload
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind to JSON"})
		return
	}

	query := `SELECT id, username, email, password_hash, bio, likes_count, followers_count,
			following_count, tweets_count, created_at FROM users WHERE email=$1`

	var user models.User

	err = db.DB.QueryRow(c, query, input.Email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Bio,
		&user.LikesCount,
		&user.FollowersCount,
		&user.FollowingCount,
		&user.TweetsCount,
		&user.CreatedAt,
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// validate password
	if !utils.ValidatePassword(user.PasswordHash, input.Password) {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// generate token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{"user": user, "token": token})
}

func DeleteUser(c *gin.Context) {

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// convert to int
	userID := userIDVal.(int)

	query := `DELETE FROM users WHERE id=$1`

	// Exec returs res and err, res gives info about query execution
	// err is set only if query itself fails
	res, err := db.DB.Exec(c, query, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// get rowsAffected from response
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfuly deleted user"})

}
