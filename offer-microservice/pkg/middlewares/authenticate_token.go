package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Authentication struct {
	SellerId string `json:"sellerId"`
}

func AuthenticateToken(AuthenticationServiceUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		authentication, err := authenticationRequest(AuthenticationServiceUrl, token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("sellerId", authentication.SellerId)
		c.Next()
	}
}

func authenticationRequest(url string, token string) (*Authentication, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Token", token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var authentication Authentication
	err = json.Unmarshal(body, &authentication)
	if err != nil {
		return nil, err
	}

	return &authentication, nil
}
