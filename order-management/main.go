package main

import (
	"order-management/helper"
	"order-management/middlewares/authenticator"
	"order-management/migrate"
	"order-management/seed"
)

func init() {
	helper.LoadEnvVariables()
}

func main() {
	var authenticator authenticator.AuthenticatorInterface = authenticator.FakeAuthenticator{}

	if !authenticator.Authenticate() {
		return
	}

	migrate.Migrate()
	seed.Seed()
}
