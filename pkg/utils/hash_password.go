package utils

import "golang.org/x/crypto/bcrypt"

const (
	cost = 15
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}
