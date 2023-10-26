package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString != "JWT-SECRET" {
			log.Println("Unauthorized")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
