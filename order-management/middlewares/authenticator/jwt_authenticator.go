package authenticator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"order-management/providers"
)

type JwtAuthenticator struct {
	AuthenticationServiceUrl string
}

func (authenticator JwtAuthenticator) NewAuthenticator(authenticatorServiceUrl string) AuthenticatorInterface {
	if authenticatorServiceUrl == "" {
		panic("authenticator service url is required")
	}

	authenticator.AuthenticationServiceUrl = authenticatorServiceUrl
	return authenticator
}

func (authenticator JwtAuthenticator) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetHeader("userEmail")
		userPassword := c.GetHeader("userPassword")

		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing user email",
			})
			c.Abort()
			return
		}

		if userPassword == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing user password",
			})
			c.Abort()
			return
		}

		token, err := providers.FetchToken(authenticator.AuthenticationServiceUrl, userEmail, userPassword)

		authentication, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid header token",
				})
				c.Abort()
				return "", errors.New("something went wrong")
			}

			return []byte(config.Get("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not validate token",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid header token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
