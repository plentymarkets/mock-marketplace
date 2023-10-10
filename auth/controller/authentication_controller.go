package controller

import (
	"auth/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		if apiKey != os.Getenv("JWT_SECRET") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}

		// implement logic to create user in database
		// implement logic to read user from database

		// init token, sep function created
		token, err := middleware.CreateJWT()
		// if error at func above
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not generate token",
			})
			c.Abort()
			return
		}

		// implement logic to save JWT token in database for each user

		// if ok -> save to database
		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
