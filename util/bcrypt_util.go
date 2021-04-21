package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword from string
// bcrypt.DefaultCost = 10
func HashPassword(password string, rounds int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(bytes), err
}

// CheckPasswordHash compare encrypt
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
