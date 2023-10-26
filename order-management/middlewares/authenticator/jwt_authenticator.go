package authenticator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"order-management/providers"
	"os"
)

type JwtAuthenticator struct {
	AuthenticationServiceUrl string
}

func (authenticator JwtAuthenticator) NewAuthenticator(authenticatorServiceUrl string) Authenticator {
	if authenticatorServiceUrl == "" {
		panic("authenticator service url is required")
	}

	authenticator.AuthenticationServiceUrl = authenticatorServiceUrl
	return authenticator
}

func (authenticator JwtAuthenticator) Authenticate() gin.HandlerFunc { // can we not do:and remove the above?
	return func(c *gin.Context) { // func (authenticator JwtAuthenticator) Authenticate(authenticatorServiceUrl string) gin.HandlerFunc
		userEmail := c.GetHeader("email")
		userPassword := c.GetHeader("password")
		authenticationApiKey := c.GetHeader("authenticationApiKey")

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

		token, err := providers.FetchToken(authenticator.AuthenticationServiceUrl, userEmail, userPassword, authenticationApiKey)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		authentication, err := jwt.Parse(token.Token, func(t *jwt.Token) (interface{}, error) { // What is this code doing?
			_, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid header token",
				})
				c.Abort()
				return "", errors.New("something went wrong")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if !authentication.Valid {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid header token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
