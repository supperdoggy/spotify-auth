package utils

import "golang.org/x/crypto/bcrypt"

func CachePassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashed, err
}
