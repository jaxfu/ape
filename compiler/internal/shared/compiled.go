package shared

import "github.com/jaxfu/ape/components"

type CompiledProp struct {
	CompiledComponentMetadata
	PropMetadata components.PropMetadata
	Constraints  components.PropConstraints
}

type CompiledObject struct {
	CompiledComponentMetadata
	Props map[string]CompiledProp
}

type CompiledBody struct {
	CompiledComponentMetadata
	BodyType components.MessageBodyType
	Ref      components.ReferenceTag
	Props    map[string]CompiledProp
}

type CompiledRouteMetadata struct {
	Url    string                `json:"url,omitempty" toml:"url,omitempty"`
	Method components.HttpMethod `json:"method,omitempty" toml:"method,omitempty"`
}
type CompiledRoute struct {
	CompiledComponentMetadata
	RouteMetadata CompiledRouteMetadata
	Request       CompiledRequest
	Responses     CompiledResponses
}

type CompiledRequest struct {
	CompiledComponentMetadata
	Headers components.HeadersMap
	Body    *CompiledBody
}

type CompiledResponse struct {
	CompiledComponentMetadata
	StatusCode uint          `json:"status_code" toml:"status_code"`
	Body       *CompiledBody `json:"body" toml:"body"`
}

type CompiledResponses = map[string]CompiledResponse
