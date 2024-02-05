package utils

import (
	jwt "github.com/golang-jwt/jwt"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
)

func GenerateToken(user db.User, secret string) (string, error) {
	// generate token
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"Email":  user.Email,
		"Name":   user.FullName,
		"role":   user.Role,
	}).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
