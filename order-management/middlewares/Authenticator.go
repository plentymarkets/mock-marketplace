package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-management/providers"
)

type Authenticator struct {
	Token string `json:"token"`
}

func Authenticate(AuthenticationServiceUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authenticator Authenticator

		if err := c.ShouldBindJSON(&authenticator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
			c.Abort()
			return
		}

		if authenticator.Token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		authenticated, err := providers.Authentication(AuthenticationServiceUrl, authenticator.Token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if !authenticated {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
