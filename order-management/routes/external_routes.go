package routes

import "os"

func GetExternalRoutesConfig() map[string]string {
	authenticationMicroService := os.Getenv("EXTERNAL_ROUTES_HOST")

	externalRoutes := map[string]string{
		"authenticationService": authenticationMicroService + "/auth/validate",
	}

	return externalRoutes
}
