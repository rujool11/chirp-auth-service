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
			created_at TIMESTAMP DEFAULT NOW()
		); `

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating table: %v", err)
	}

	log.Printf("users table is ready")

}
