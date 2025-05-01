package components

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/pkg/enum"
)

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
	Components  = map[ComponentId]Component
)

type (
	ComponentType           = string
	ComponentTypesInterface struct {
		PROP         ComponentType
		OBJECT       ComponentType
		ROUTE        ComponentType
		REQUEST      ComponentType
		RESPONSE     ComponentType
		MESSAGE_BODY ComponentType
		UNDEFINED    ComponentType
	}
)

var ComponentTypes = enum.Enum[ComponentType, ComponentTypesInterface]{
	TypeList: ComponentTypesImpl,
	MatchMap: map[string]ComponentType{},
}

var ComponentTypesImpl = ComponentTypesInterface{
	PROP:         "PROP",
	OBJECT:       "OBJECT",
	ROUTE:        "ROUTE",
	REQUEST:      "REQUEST",
	RESPONSE:     "RESPONSE",
	MESSAGE_BODY: "MESSAGE_BODY",
	UNDEFINED:    "UNDEFINED",
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

func GenerateComponentId(meta ComponentMetadata) (ComponentId, error) {
	if meta.ComponentType == "" {
		return "", fmt.Errorf("no component type given")
	}

	idParams := []string{}
	if meta.IsRoot {
		if strings.TrimSpace(meta.Name) == "" {
			return "", fmt.Errorf("no name given for root component")
		}

		if meta.Category != nil {
			idParams = append(idParams, *meta.Category)
		}

		typeName, ok := typePluralNames[meta.ComponentType]
		if !ok {
			return "", fmt.Errorf("invalid type %s", meta.ComponentType)
		}

		idParams = append(idParams, typeName)
		idParams = append(idParams, meta.Name)

	} else { // if not root
		if meta.ParentId == nil {
			return "", fmt.Errorf("no parentId given for child")
		}
		if strings.ToLower(*meta.ParentId) == "" {
			return "", fmt.Errorf("no parentId given for child")
		}

		idParams = append(idParams, *meta.ParentId)

		if meta.Name == "" {
			name := typeNames[meta.ComponentType]
			idParams = append(idParams, name)
		} else {
			idParams = append(idParams, meta.Name)
		}
	}

	return strings.Join(idParams, "."), nil
}

var typeNames = map[ComponentType]string{
	ComponentTypesImpl.PROP:         "prop",
	ComponentTypesImpl.OBJECT:       "object",
	ComponentTypesImpl.ROUTE:        "route",
	ComponentTypesImpl.MESSAGE_BODY: "body",
	ComponentTypesImpl.REQUEST:      "request",
	ComponentTypesImpl.RESPONSE:     "response",
}

var typePluralNames = map[ComponentType]string{
	ComponentTypesImpl.PROP:         "props",
	ComponentTypesImpl.OBJECT:       "objects",
	ComponentTypesImpl.ROUTE:        "routes",
	ComponentTypesImpl.MESSAGE_BODY: "bodies",
	ComponentTypesImpl.REQUEST:      "requests",
	ComponentTypesImpl.RESPONSE:     "responses",
}
