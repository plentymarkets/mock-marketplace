package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product-management/pkg/client"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Token")

		if headerToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing header token"})
			c.Abort()
			return
		}

		// Find User by token
		authRequest := client.NewAuthTokenClient("auth/validate", "TestToken")
		response, err := authRequest.ValidateToken()

		if err != nil || response.StatusCode != http.StatusOK {
			log.Printf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Authenticate"})
			c.Abort()
			return
		}

		c.Next()
	}
}
