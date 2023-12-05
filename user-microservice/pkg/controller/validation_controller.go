package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"time"
	tokenGenerator "user-microservice/pkg/middleware/token"
	"user-microservice/pkg/repositories"
)

func Validate(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			return
		}

		authentication, err := parseJwt(token)

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

		userRepository := repositories.NewRepository(databaseConnection)
		user, err := userRepository.FindOneByField("token", token)

		if user.TokenExpiration.Before(time.Now()) && user.RefreshTokenExpiration.After(time.Now()) {
			token, timestamp, refreshTimestamp, err := tokenGenerator.Generate()

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
				return
			}

			user.Token = token
			user.TokenExpiration = timestamp
			user.RefreshTokenExpiration = refreshTimestamp
			userRepository.UpdateUser(user)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
				return
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": "token doesn't belong to any user",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"sellerId": strconv.Itoa(user.SellerID),
		})
	}
}

func parseJwt(token string) (*jwt.Token, error) {
	authentication, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return authentication, err
}
