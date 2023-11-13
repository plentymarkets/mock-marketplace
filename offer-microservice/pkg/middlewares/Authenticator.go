package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(AuthenticationServiceUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		authenticated, err := authenticationRequest(AuthenticationServiceUrl, token)

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

func authenticationRequest(url string, token string) (bool, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	request.Header.Add("Token", token)

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}
