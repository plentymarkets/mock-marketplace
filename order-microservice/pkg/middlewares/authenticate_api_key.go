package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthenticateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("apiKey")

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing apiKey",
			})
			c.Abort()
			return
		}

		if apiKey != os.Getenv("API_KEY") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid secret",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
