package routes

import "os"

func GetExternalRoutesConfig() map[string]string {
	externalRoutes := map[string]string{
		"authenticationService": os.Getenv("EXTERNAL_ROUTES_HOST") + "/authenticate",
	}

	return externalRoutes
}
