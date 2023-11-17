package external_router

import (
	"os"
	"strings"
)

type ExternalRouter struct {
	routes map[string]ExternalRoute
}

type ExternalRoute struct {
	Url    string
	Method string
}

func NewExternalRouter() ExternalRouter {
	var externalRouter ExternalRouter
	externalRouter.routes = make(map[string]ExternalRoute)

	externalRouter.AddRoute("validate-token", os.Getenv("AUTHENTICATOR_MICROSERVICE_URL")+"/user/validation", "GET")
	externalRouter.AddRoute("get-offer-by-id", os.Getenv("OFFER_MICROSERVICE_URL")+"/internal/offers/{offerId}", "GET")

	return externalRouter
}

func (externalRouter *ExternalRouter) AddRoute(name string, url string, method string) {
	route := ExternalRoute{
		Url:    url,
		Method: method,
	}

	externalRouter.routes[name] = route
}

func (externalRouter *ExternalRouter) GetRoute(name string, parameters map[string]string) ExternalRoute {
	route := externalRouter.routes[name]

	if parameters != nil {
		route = replacePlaceholders(parameters, route)
	}

	return route
}

func replacePlaceholders(params map[string]string, route ExternalRoute) ExternalRoute {
	for k, v := range params {
		placeholder := "{" + k + "}"
		route.Url = strings.Replace(route.Url, placeholder, v, -1)
	}

	return route
}
