package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
	"user-microservice/pkg/middleware/token"
	"user-microservice/pkg/models"
	"user-microservice/pkg/repositories"
)

type Credentials struct {
	AuthenticationApiKey string `json:"authenticationApiKey"`
	Email                string `json:"email"`
	Password             string `json:"password"`
}

func RetrieveToken(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials Credentials

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
			return
		}

		if credentials.AuthenticationApiKey == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing api key",
			})
			return
		}

		if credentials.AuthenticationApiKey != os.Getenv("AUTHENTICATION_API_KEY") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			return
		}

		if credentials.Email == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing email",
			})
			return
		}

		if credentials.Password == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing password",
			})
			return
		}

		userRepository := repositories.NewRepository(databaseConnection)
		user, err := userRepository.FindOneByField("email", credentials.Email)

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			return
		}

		if user.Password != credentials.Password {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			return
		}

		if tokenIsNotExpired(user) {
			c.JSON(http.StatusOK, map[string]string{
				"token": user.Token,
			})
			return
		}

		timeNow := time.Now()
		println(timeNow.String())
		token, timestamp, err := token.Generate()

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not generate token",
			})
			return
		}

		user.Token = token
		user.TokenExpiration = timestamp
		userRepository.UpdateUser(user)

		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func tokenIsNotExpired(user *models.User) bool {
	return time.Now().Unix() < user.TokenExpiration.Unix() && !user.TokenExpiration.IsZero()
}
