package authenticator

import "github.com/gin-gonic/gin"

type FakeAuthenticator struct{}

func (authenticator FakeAuthenticator) NewAuthenticator(string) Authenticator {
	return authenticator
}

func (authenticator FakeAuthenticator) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
