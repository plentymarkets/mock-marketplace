package authenticator

type FakeAuthenticator struct{}

func (FakeAuthenticator) Authenticate() bool {
	return true
}
