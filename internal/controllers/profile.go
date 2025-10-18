package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-auth-service/internal/db"
	"github.com/rujool11/chirp-auth-service/internal/models"
)

func GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID := userIDVal.(int)
	query := `SELECT id, username, email, bio, likes_count, followers_count,
			following_count, tweets_count, created_at FROM users WHERE id=$1`

	var user models.User
	err := db.DB.QueryRow(c, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Bio,
		&user.LikesCount,
		&user.FollowersCount,
		&user.FollowingCount,
		&user.TweetsCount,
		&user.CreatedAt,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func UpdateBio(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID := userIDVal.(int)
	var input struct {
		Bio string `json:"bio" binding:"required"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind to JSON"})
		return
	}

	query := `UPDATE users SET bio=$1 WHERE id=$2`
	_, err = db.DB.Exec(c, query, input.Bio, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update bio"})
		return
	}

	c.JSON(200, gin.H{"message": "Updated bio successfully", "bio": input.Bio})

}

func UpdatePassword(c *gin.Context) {

}
