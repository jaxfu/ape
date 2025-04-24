package components

type ComponentMetadata struct {
	ComponentType ComponentType
	ComponentId   ComponentId
	Name          string
	IsRoot        bool
	ParentId      *ComponentId
	Category      *CategoryId
	Description   *string
}

func (meta ComponentMetadata) Metadata() ComponentMetadata {
	return meta
}

type (
	ComponentId = string
	CategoryId  = string
)

type ComponentType = string

// TODO: use map
const (
	COMPONENT_TYPE_OBJECT       ComponentType = "OBJECT"
	COMPONENT_TYPE_PROP         ComponentType = "PROP"
	COMPONENT_TYPE_ROUTE        ComponentType = "ROUTE"
	COMPONENT_TYPE_REQUEST      ComponentType = "REQUEST"
	COMPONENT_TYPE_RESPONSE     ComponentType = "RESPONSE"
	COMPONENT_TYPE_MESSAGE_BODY ComponentType = "MESSAGE_BODY"
)

type AllComponents struct {
	Ids           map[ComponentId]Component
	Props         map[ComponentId]Prop
	Objects       map[ComponentId]Object
	Routes        map[ComponentId]Route
	Requests      map[ComponentId]Request
	Response      map[ComponentId]Response
	MessageBodies map[ComponentId]MessageBody
}

type Component interface {
	Metadata() ComponentMetadata
}

func ComponentSwitch(comp Component) {
	switch comp.(type) {
	case Prop:
	case Object:
	case Route:
	case MessageBody:
	case Request:
	case Response:
	}
}
