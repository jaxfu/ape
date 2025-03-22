package components

import (
	"github.com/jaxfu/ape/components/internal/routes"
)

type (
	Route         = routes.Route
	RouteMetadata = routes.RouteMetadata
	Request       = routes.Request
	HttpMethod    = routes.HttpMethod
	HeadersMap    = routes.HeadersMap
	Response      = routes.Response
	ResponsesMap  = routes.ResponsesMap
)

const (
	ROUTE_METHOD_GET  HttpMethod = "GET"
	ROUTE_METHOD_POST HttpMethod = "POST"
)
