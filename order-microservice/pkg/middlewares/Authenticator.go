package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/routes/external-router"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		externalRouter := external_router.NewExternalRouter()
		authenticationServiceRoute := externalRouter.GetRoute("validate-token", nil)

		token := c.GetHeader("Token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
			return
		}

		authenticated, err := authenticationRequest(authenticationServiceRoute, token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if !authenticated {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
			return
		}

		c.Next()
	}
}

func authenticationRequest(authenticationServiceRoute external_router.ExternalRoute, token string) (bool, error) {
	client := &http.Client{}

	request, err := http.NewRequest(authenticationServiceRoute.Method, authenticationServiceRoute.Url, nil)
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
