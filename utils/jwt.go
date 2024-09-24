package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "userId": userId, "exp": time.Now().Add(time.Hour * 2).Unix()})
	return token.SignedString([]byte(key))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, error := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// check the token
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing token")
		}
		return []byte(key), nil
	})
	if error != nil {
		return 0, errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claim, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// email:=claim["email"].(string)
	userId := claim["userId"].(float64)
	fmt.Printf("%T,", userId)
	return int64(userId), nil
}
