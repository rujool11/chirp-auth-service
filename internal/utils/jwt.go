package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// saved as slice of bytes
var jwtKey = []byte(os.Getenv("JWT_KEY"))

// Claims is the data stored in the token
type Claims struct {
	ID                   int `json:"id"` // data encoded in token
	jwt.RegisteredClaims     // standard claims
}

// create jwt token given user id
func GenerateJWT(ID int) (string, error) {
	expiration := time.Now().Add(24 * 7 * 4 * time.Hour) // validity 4 weeks

	// pointer to a struct
	claims := &Claims{
		ID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration), // sets exp
			IssuedAt:  jwt.NewNumericDate(time.Now()), // sets iat
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HMAC SHA 256 signing
	return token.SignedString(jwtKey)                          // automatically returns string and error
}
