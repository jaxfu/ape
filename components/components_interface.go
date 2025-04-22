package components

import (
	"github.com/jaxfu/ape/components/internal/objects"
	"github.com/jaxfu/ape/components/internal/props"
	"github.com/jaxfu/ape/components/internal/routes"
	"github.com/jaxfu/ape/components/internal/shared"
)

type (
	ComponentMetadata = shared.ComponentMetadata
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
	Props         map[ComponentId]Prop
	Objects       map[ComponentId]Object
	Routes        map[ComponentId]Route
	Requests      map[ComponentId]Request
	Response      map[ComponentId]Response
	MessageBodies map[ComponentId]MessageBody
}

type ComponentsBag struct {
	Type     ComponentType
	Prop     *props.Prop
	Object   *objects.Object
	Route    *routes.Route
	Body     *routes.MessageBody
	Request  *routes.Request
	Response *routes.Response
}

// switch template
//
// switch ComponentType {
// case components.COMPONENT_TYPE_PROP:
// case components.COMPONENT_TYPE_OBJECT:
// case components.COMPONENT_TYPE_ROUTE:
// case components.COMPONENT_TYPE_MESSAGE_BODY:
// case components.COMPONENT_TYPE_REQUEST:
// case components.COMPONENT_TYPE_RESPONSE:
// }
