package shared

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/shared"

	"github.com/jaxfu/ape/components"
)

const (
	KEY_NAME         string = "name"
	KEY_CATEGORY     string = "category"
	KEY_DESCRIPTION  string = "description"
	KEY_TYPE         string = "type"
	KEY_ARRAY        string = "array"
	KEY_PROPS        string = "props"
	KEY_REQUEST      string = "request"
	KEY_RESPONSES    string = "responses"
	KEY_HEADERS      string = "headers"
	KEY_MESSAGE_BODY string = "body"
	KEY_STATUS_CODE  string = "status_code"
)

// TODO: validate name at parsing (if required)
func ParseComponentMetadata(fields map[string]any, compType components.ComponentType, isRoot bool) (ParsedComponentMetadata, error) {
	metadata := ParsedComponentMetadata{}

	// check if name required
	mapName, exists, err := GetStringFromMap(
		fields,
		KEY_NAME,
	)
	if exists {
		if err != nil {
			return ParsedComponentMetadata{}, fmt.Errorf("error parsing %s: %+v", KEY_NAME, err)
		}
		metadata.Name = &mapName
	} else if isRoot {
		return ParsedComponentMetadata{}, fmt.Errorf("name missing")
	} else if _, ok := typesNameRequired[compType]; ok {
		return ParsedComponentMetadata{}, fmt.Errorf("name missing, required for type %s", compType)
	} else {
		typeName := typeIdChildrenNames[compType]
		metadata.Name = &typeName
	}

	category, _, err := GetStringFromMap(
		fields,
		KEY_CATEGORY,
	)
	if err == nil {
		metadata.Category = &category
	}

	description, _, err := GetStringFromMap(
		fields,
		KEY_DESCRIPTION,
	)
	if err == nil {
		metadata.Description = &description
	}

	return metadata, nil
}

var componentTypes = components.ComponentTypes.Types()

var typeIdNames = map[components.ComponentType]string{
	componentTypes.PROP:         "props",
	componentTypes.OBJECT:       "objects",
	componentTypes.ROUTE:        "routes",
	componentTypes.MESSAGE_BODY: "bodies",
	componentTypes.REQUEST:      "requests",
	componentTypes.RESPONSE:     "responses",
}

var typeIdChildrenNames = map[components.ComponentType]string{
	componentTypes.PROP:         "prop",
	componentTypes.OBJECT:       "object",
	componentTypes.ROUTE:        "route",
	componentTypes.MESSAGE_BODY: "body",
	componentTypes.REQUEST:      "request",
	componentTypes.RESPONSE:     "response",
}

var typesNameRequired = map[components.ComponentType]*any{
	componentTypes.PROP:     nil,
	componentTypes.RESPONSE: nil,
}

type ParsedComponentMetadata struct {
	Name        *string `json:"name,omitempty" toml:"name,omitempty"`
	Category    *string `json:"category,omitempty" toml:"category,omitempty"`
	Description *string `json:"description,omitempty" toml:"description,omitempty"`
}
type Context = shared.CompilationContext
