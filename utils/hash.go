package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytePassword), err
}
