package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), error
}

func CheckPasswordHash(hashPassword, password string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return error == nil
}
