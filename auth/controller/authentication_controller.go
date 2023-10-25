package controller

import (
	"auth/middleware"
	"auth/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func Auth(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("ApiKey") // What API key?

		if apiKey == "" { // Combine the 2 ifs in 1 and return a more generic error
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing api key",
			})
			c.Abort()
			return
		}

		if apiKey != os.Getenv("JWT_SECRET") { // What is the JWT_SECRET and what if i want 10 account
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}

		email := c.GetHeader("email")       // And why are they in the header and not the body?
		password := c.GetHeader("password") // Why is password plain text?

		userRepository := repositories.NewRepository(databaseConnection)
		user := userRepository.GetUserByEmail(email) // Error handling

		if user.Password != password { // Why do we check here the credentials
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			c.Abort()
			return
		}

		//if user.TokenExpiration.After(time.Now()) {
		//	c.JSON(http.StatusInternalServerError, map[string]string{
		//		"error": "token is still valid",
		//	})
		//}

		// init token, sep function created
		token, timestamp, err := middleware.CreateJWT() // Why is this a middleware? What is a middleware?
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
