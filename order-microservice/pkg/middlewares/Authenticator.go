package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
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

		authenticated, err := authenticationRequest(AuthenticationServiceUrl, authenticator.Token)

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

	data := map[string]string{"token": token}
	jsonData, err := json.Marshal(data)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	request.Header.Add("content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}
