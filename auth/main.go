package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.GetHeader, will give me all headers (secret ones) values
		apiKey := c.GetHeader("ApiKey")
		// if for validation == 'null'
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing api key",
			})
			c.Abort()
			return
		}
		// if for validation == true
		// we hardcode it, in real env we get it from database
		if apiKey != ("JWT_SECRET") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}
		// init token, sep function created
		token, err := createJWT()
		// if error at func above
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not generate token",
			})
			c.Abort()
			return
		}
		// if ok -> response as JSON return
		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
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
	// start default
	r := gin.Default()
	// main route /api
	api := r.Group("/api")
	// main route /api/auth and call func auth
	api.POST("/auth", Auth())
	err := r.Run(":3001")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
