package main

import (
	"auth/migrate"
	"github.com/joho/godotenv"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
}

func createJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString([]byte(("JWT_SECRET")))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenStr, nil
}

func main() {
	migrate.Migrate()
}

//Route zur Authentifizierung -> routes.go
//Controller -> AuthController der quasi sagt ob man authentifiziert wurde.
//Entities -> Models -> User (Username, Password, Token)
//Gorm -> ORM -> User
//Gin -> HTTP Framework
