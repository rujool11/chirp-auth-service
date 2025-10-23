package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-auth-service/internal/db"
	"github.com/rujool11/chirp-auth-service/internal/models"
)

func GetAllUsers(c *gin.Context) {
	query := `SELECT id, username, email, bio, likes_count, followers_count, following_count, tweets_count, created_at
			FROM users`

	rows, err := db.DB.Query(c, query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not fetch users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.LikesCount, &user.FollowersCount, &user.FollowingCount, &user.TweetsCount, &user.CreatedAt); err != nil {
			continue
		}

		users = append(users, user)
	}

	c.JSON(200, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	query := `SELECT id, username, email, bio, likes_count, followers_count, following_count, tweets_count, created_at
			FROM users
			WHERE id=$1`

	var user models.User

	err = db.DB.QueryRow(c, query, id).Scan(
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
		c.JSON(500, gin.H{"error": "Could not fetch user"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	query := `SELECT id, username, email, bio, likes_count, followers_count, following_count, tweets_count, created_at
			FROM users
			WHERE username=$1`

	var user models.User

	err := db.DB.QueryRow(c, query, username).Scan(
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
		c.JSON(500, gin.H{"error": "Could not fetch user"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}
