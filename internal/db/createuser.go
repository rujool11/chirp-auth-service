package db

import (
	"context"
	"log"
)

func CreateUserTableIfDoesNotExist() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(50) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			bio TEXT DEFAULT '',
			likes_count INT DEFAULT 0,
			followers_count INT DEFAULT 0,
			following_count INT DEFAULT 0,
			tweets_count INT DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating table: %v", err)
	}

	log.Println("Users table is ready")
}
