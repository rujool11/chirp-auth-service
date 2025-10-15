package models

import "time"

type User struct {
	// struct tags used to indicate JSON key for JSON encoder/decoder
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // ignore this when converting to json
	CreatedAt    time.Time `json:"created_at"`
}
