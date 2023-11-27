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
		response, err := client.ValidateToken(headerToken)

		if err != nil {
			log.Printf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "There is a problem with the authentication process "})
			c.Abort()
			return
		}

		if response.StatusCode != http.StatusOK {
			log.Printf("Error code: %s", response.Status)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Authenticate"})
			c.Abort()
			return
		}

		c.Request.Header.Add("seller-id", "1")
		c.Next()
	}
}
