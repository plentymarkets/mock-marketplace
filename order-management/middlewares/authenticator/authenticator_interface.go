package authenticator

import "github.com/gin-gonic/gin"

type Authenticator interface {
	NewAuthenticator(authenticatorServiceUrl string) Authenticator
	Authenticate() gin.HandlerFunc
}
