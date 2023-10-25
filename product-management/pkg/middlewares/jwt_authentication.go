package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Token")

		if headerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing header token",
			})
			c.Abort()
			return
		}

		// Find User by token

		// Validate Token to Auth
		// token, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

		//if err != nil {
		//	log.Fatalln(err)
		//}

		// If:   token is valid save user to memory
		// Else: Return Unvalid Token

		c.Next()
	}
}
