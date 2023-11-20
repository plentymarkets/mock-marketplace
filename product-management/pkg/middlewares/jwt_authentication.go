package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product-management/pkg/client"
)

type authenticateToken struct {
	SellerId string `json:"sellerId"`
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Token")

		if headerToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing header token"})
			c.Abort()
			return
		}

		// Find User by token
		response, err := client.AuthenticationRequest(headerToken)

		if err != nil {
			log.Printf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "There is a problem with the authentication process "})
			c.Abort()
			return
		}

		c.Set("sellerId", response.SellerId)
		c.Next()
	}
}
