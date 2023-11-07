package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

type Authenticator struct {
	Token string `json:"token"`
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authenticator Authenticator

		if err := c.ShouldBindJSON(&authenticator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
			return
		}

		if authenticator.Token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		authentication, err := jwt.Parse(authenticator.Token, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid header token",
				})
				c.Abort()
				return "", errors.New("something went wrong")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if !authentication.Valid {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": "valid token",
		})
	}
}
