package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var Secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

// TODO: Relocate to forum-service
//func ValidateToken(tokenString string) (*jwt.Token, error) {
//	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return Secret, nil
//	})
//}
