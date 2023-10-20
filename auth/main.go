package main

import (
	"auth/helper"
	"auth/migrate"
	"auth/routes"
	"auth/seed"
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
	databaseConnection := helper.GetDatabaseConnection()

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection)
}
