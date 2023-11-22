package token

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func Generate() (string, time.Time, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expiration := time.Now().Add(time.Hour)
	claims["expiration"] = expiration.Unix()
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatal(err)
		return "", expiration, err
	}

	return tokenStr, expiration, nil
}
