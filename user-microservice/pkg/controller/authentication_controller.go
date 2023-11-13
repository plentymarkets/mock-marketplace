package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			return
		}

		authentication, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if !authentication.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": "valid token",
		})
	}
}
