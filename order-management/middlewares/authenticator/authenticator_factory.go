package authenticator

import "fmt"

func CreateAuthenticator(driver string) (AuthenticatorInterface, error) {
	var authenticator AuthenticatorInterface
	var err error

	switch driver {
	case "fake":
		authenticator = &FakeAuthenticator{}
	case "jwt":
		authenticator = &JwtAuthenticator{}
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}

	if err != nil {
		return nil, err
	}

	return authenticator, nil
}
