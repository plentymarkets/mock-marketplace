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
		requestURL := "http://host.docker.internal:3004/auth/validate" // TODO - Remove hardcoded stuff
		response, err := client.GET(requestURL, headerToken)

		if err != nil || response.StatusCode != http.StatusOK {
			if headerToken == "" {
				log.Printf(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Authenticate"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
