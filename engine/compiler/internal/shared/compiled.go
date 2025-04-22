package shared

import "github.com/jaxfu/ape/components"

type (
	CompiledProps = map[string]CompiledProp
	CompiledProp  struct {
		ComponentMetadata CompiledComponentMetadata
		PropMetadata      components.PropMetadata
		Constraints       components.PropConstraints
	}
)

type CompiledObject struct {
	ComponentMetadata CompiledComponentMetadata
	Props             CompiledProps
}

type CompiledBody struct {
	ComponentMetadata CompiledComponentMetadata
	BodyType          components.MessageBodyType
	Ref               components.ReferenceTag
	Props             CompiledProps
}

type CompiledRouteMetadata struct {
	Url    string                `json:"url,omitempty" toml:"url,omitempty"`
	Method components.HttpMethod `json:"method,omitempty" toml:"method,omitempty"`
}
type CompiledRoute struct {
	ComponentMetadata CompiledComponentMetadata
	RouteMetadata     CompiledRouteMetadata
	Request           CompiledRequest
	Responses         CompiledResponses
}

type CompiledRequest struct {
	ComponentMetadata CompiledComponentMetadata
	Headers           components.HeadersMap
	Body              *CompiledBody
}

type CompiledResponse struct {
	ComponentMetadata CompiledComponentMetadata
	StatusCode        uint          `json:"status_code" toml:"status_code"`
	Body              *CompiledBody `json:"body" toml:"body"`
}

type CompiledResponses = map[string]CompiledResponse
