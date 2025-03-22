package components

import "github.com/jaxfu/ape/components/internal/shared"

type (
	ComponentMetadata = shared.ComponentMetadata
	ComponentContext  = shared.ComponentContext
	ComponentType     = shared.ComponentType
	ComponentId       = shared.ComponentId
	CategoryId        = shared.CategoryId
)

const (
	COMPONENT_TYPE_OBJECT       ComponentType = "OBJECT"
	COMPONENT_TYPE_PROP         ComponentType = "PROP"
	COMPONENT_TYPE_ROUTE        ComponentType = "ROUTE"
	COMPONENT_TYPE_REQUEST      ComponentType = "REQUEST"
	COMPONENT_TYPE_RESPONSE     ComponentType = "RESPONSE"
	COMPONENT_TYPE_MESSAGE_BODY ComponentType = "MESSAGE_BODY"
)

type AllComponents struct {
	Props       []Prop
	Objects     []Object
	Routes      []Route
	Requests    []Request
	Response    []Response
	MessageBody []MessageBody
}
