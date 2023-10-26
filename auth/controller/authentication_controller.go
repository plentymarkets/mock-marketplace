package controller

import (
	"auth/middleware"
	"auth/models"
	"auth/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

// Auth
// TODO: This function uses the header instead of the body. Is this the best approach for handling sensitive data?
// TODO: Errors are very specific. Are they TOO specific?
// TODO: Passwords are not hashed. This sould be implemented in the future.
func Auth(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticationApiKey := c.GetHeader("authenticationApiKey")

		if authenticationApiKey == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing api key",
			})
			c.Abort()
			return
		}

		if authenticationApiKey != os.Getenv("AUTHENTICATION_API_KEY") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}

		email := c.GetHeader("email")

		if email == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing email",
			})
			c.Abort()
			return
		}

		password := c.GetHeader("password")

		if password == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing password",
			})
			c.Abort()
			return
		}

		userRepository := repositories.NewRepository(databaseConnection)
		user := userRepository.GetUserByEmail(email)

		if user.Password != password {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			c.Abort()
			return
		}

		if tokenIsntExpired(user) {
			c.JSON(http.StatusOK, map[string]string{
				"token": user.Token,
			})
			return
		}

		timeNow := time.Now()
		println(timeNow.String())
		token, timestamp, err := middleware.CreateJWT()

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

		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})

		c.Done()
	}
}

func tokenIsntExpired(user models.User) bool {
	return time.Now().Unix() < user.TokenExpiration.Unix() && !user.TokenExpiration.IsZero()
}
