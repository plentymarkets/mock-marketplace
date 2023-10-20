package authenticator

import "github.com/gin-gonic/gin"

type AuthenticatorInterface interface {
	NewAuthenticator(authenticatorServiceUrl string) AuthenticatorInterface
	Authenticate() gin.HandlerFunc
}
