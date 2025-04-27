package components

type ComponentMetadata struct {
	ComponentType ComponentType `json:"component_type"`
	ComponentId   ComponentId   `json:"component_id"`
	Name          string        `json:"name"`
	IsRoot        bool          `json:"is_root"`
	ParentId      *ComponentId  `json:"parent_id,omitempty"`
	Category      *CategoryId   `json:"category,omitempty"`
	Description   *string       `json:"description,omitempty"`
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

type Components = map[ComponentId]Component

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
