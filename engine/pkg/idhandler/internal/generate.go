package internal

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/components"
)

func DefaultIdHandler() IdHandler {
	return IdHandler{}
}

type IdHandler struct{}

// TODO: return string
func (idh IdHandler) Generate(params GenerateIdParams) (components.ComponentId, error) {
	// if name required, need name
	if _, ok := typesNameRequired[params.ComponentType]; ok {
		if params.Name == nil || *params.Name == "" {
			return components.ComponentId{}, fmt.Errorf("no name given for root component")
		}
	}

	idParams := []string{}
	if params.IsRoot {
		if params.Name == nil || *params.Name == "" {
			return components.ComponentId{}, fmt.Errorf("no name given for root component")
		}
		if params.Category != nil {
			idParams = append(idParams, *params.Category)
		}
		typeName, ok := typeIdNames[params.ComponentType]
		if !ok {
			return components.ComponentId{}, fmt.Errorf("invalid type %s", params.ComponentType)
		}

		idParams = append(idParams, typeName)
		idParams = append(idParams, *params.Name)

	} else { // if not root
		if params.ParentId == nil {
			return components.ComponentId{}, fmt.Errorf("no parentId given for child")
		}
		if params.ParentId.Display == "" {
			return components.ComponentId{}, fmt.Errorf("empty parentId given for child %s", *params.Name)
		}
		idParams = append(idParams, params.ParentId.Display)

		if params.Name != nil {
			if *params.Name == "" {
				name := typeIdChildrenNames[params.ComponentType]
				idParams = append(idParams, name)
			} else {
				idParams = append(idParams, *params.Name)
			}
		} else {
			name := typeIdChildrenNames[params.ComponentType]
			idParams = append(idParams, name)
		}
	}

	return components.ComponentId{
		Display: strings.Join(idParams, "."),
	}, nil
}

type GenerateIdParams struct {
	components.ComponentContext
	Category *string
}

var typeIdNames = map[components.ComponentType]string{
	components.COMPONENT_TYPE_PROP:         "props",
	components.COMPONENT_TYPE_OBJECT:       "objects",
	components.COMPONENT_TYPE_ROUTE:        "routes",
	components.COMPONENT_TYPE_MESSAGE_BODY: "bodies",
	components.COMPONENT_TYPE_REQUEST:      "requests",
	components.COMPONENT_TYPE_RESPONSE:     "responses",
}

var typeIdChildrenNames = map[components.ComponentType]string{
	components.COMPONENT_TYPE_PROP:         "prop",
	components.COMPONENT_TYPE_OBJECT:       "object",
	components.COMPONENT_TYPE_ROUTE:        "route",
	components.COMPONENT_TYPE_MESSAGE_BODY: "body",
	components.COMPONENT_TYPE_REQUEST:      "request",
	components.COMPONENT_TYPE_RESPONSE:     "response",
}

var typesNameRequired = map[components.ComponentType]*any{
	components.COMPONENT_TYPE_PROP:     nil,
	components.COMPONENT_TYPE_RESPONSE: nil,
}
