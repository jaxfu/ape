package routes

import (
	"github.com/jaxfu/ape/components/internal/routes/internal"
	"github.com/jaxfu/ape/components/internal/routes/internal/body"
	"github.com/jaxfu/ape/components/internal/routes/internal/request"
	"github.com/jaxfu/ape/components/internal/routes/internal/responses"
	"github.com/jaxfu/ape/components/internal/shared"
)

type Route struct {
	ComponentMetadata shared.ComponentMetadata
	RouteMetadata     internal.RouteMetadata
	Request           Request      `json:"request" toml:"request"`
	Responses         ResponsesMap `json:"responses" toml:"responses"`
}

func (r Route) GetMetadata() shared.ComponentMetadata {
	return r.ComponentMetadata
}

type (
	Request         = request.Request
	Response        = responses.Response
	ResponsesMap    = responses.ResponsesMap
	MessageBody     = body.MessageBody
	MessageBodyType = body.MessageBodyType
	HttpMethod      = internal.HttpMethod
	RouteMetadata   = internal.RouteMetadata
	HeadersMap      = request.HeadersMap
)

const (
	MESSAGE_BODY_TYPE_REF   = body.REF
	MESSAGE_BODY_TYPE_PROPS = body.PROPS
)

const (
	HTTP_METHOD_GET  HttpMethod = "GET"
	HTTP_METHOD_POST HttpMethod = "POST"
)
