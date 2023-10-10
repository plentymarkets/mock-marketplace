package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenStr, nil
}
