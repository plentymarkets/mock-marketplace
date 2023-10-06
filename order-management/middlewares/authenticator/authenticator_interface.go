package authenticator

type AuthenticatorInterface interface {
	Authenticate() bool
}
