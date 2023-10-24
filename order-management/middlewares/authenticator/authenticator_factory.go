package authenticator

import "fmt"

func NewAuthenticator(driver string) (Authenticator, error) {
	var authenticator Authenticator
	var err error

	switch driver {
	case "fake":
		authenticator = &FakeAuthenticator{}
	case "jwt":
		authenticator = &JwtAuthenticator{}
	default:
		authenticator = nil
		err = fmt.Errorf("unknown driver: %s", driver)
	}

	return authenticator, err
}
