package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/logger"
)

type authenticateToken struct {
	SellerId string `json:"sellerId"`
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		externalRouter := external_router.NewExternalRouter()
		authenticationServiceRoute := externalRouter.GetRoute("validate-token", nil)

		token := c.GetHeader("Token")

		if token == "" {
			logger.Log("missing token", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			return
		}

		authentication, err := authenticationRequest(authenticationServiceRoute, token)

		if err != nil {
			logger.Log("could not authenticate token", nil)
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.Set("sellerId", authentication.SellerId)
		c.Next()
	}
}

func authenticationRequest(route external_router.ExternalRoute, token string) (*authenticateToken, error) {
	client := &http.Client{}

	req, err := http.NewRequest(route.Method, route.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var authentication authenticateToken
	err = json.Unmarshal(body, &authentication)
	if err != nil {
		return nil, err
	}

	return &authentication, nil
}
