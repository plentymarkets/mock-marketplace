package controller

import (
	"auth/middleware"
	"auth/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func Auth(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("ApiKey")

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing api key",
			})
			c.Abort()
			return
		}

		if apiKey != os.Getenv("JWT_SECRET") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}

		userEmail := c.GetHeader("userEmail")
		userPassword := c.GetHeader("userPassword")

		userRepository := repositories.NewRepository(databaseConnection)
		user := userRepository.GetUserByEmail(userEmail)

		if user.Password != userPassword {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			c.Abort()
			return
		}

		if user.TokenExpiration.After(time.Now()) {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "token is still valid",
			})
		}

		// init token, sep function created
		token, timestamp, err := middleware.CreateJWT()
		// if error at func above
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not generate token",
			})
			c.Abort()
			return
		}

		user.Token = token
		user.TokenExpiration = timestamp
		userRepository.UpdateUser(user)

		// if ok -> save to database
		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
