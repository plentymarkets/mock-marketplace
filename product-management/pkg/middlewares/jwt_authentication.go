package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//headerToken := c.GetHeader("Token")
		//
		//if headerToken == "" {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"error": "missing header token",
		//	})
		//	c.Abort()
		//	return
		//}
		//
		//_, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		//
		//if err != nil {
		//	log.Fatalln(err)
		//}
		c.Next()
	}
}
