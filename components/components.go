package components

const (
	COMPONENT_TYPE_OBJECT    ComponentType = "OBJECT"
	COMPONENT_TYPE_ENUM      ComponentType = "ENUM"
	COMPONENT_TYPE_ARRAY     ComponentType = "ARRAY"
	COMPONENT_TYPE_REFERENCE ComponentType = "REFERENCE"

	COMPONENT_TYPE_STRING ComponentType = "STRING"
	COMPONENT_TYPE_BLOB   ComponentType = "BLOB"
	COMPONENT_TYPE_INT    ComponentType = "INT"
	COMPONENT_TYPE_UINT   ComponentType = "UINT"
	COMPONENT_TYPE_FLOAT  ComponentType = "FLOAT"
	COMPONENT_TYPE_BOOL   ComponentType = "BOOL"

	COMPONENT_TYPE_UNDEFINED ComponentType = "UNDEFINED"
)

// switch ctype {
// case COMPONENT_TYPE_OBJECT:
// case COMPONENT_TYPE_ENUM:
// case COMPONENT_TYPE_ARRAY:
// case COMPONENT_TYPE_REFERENCE:
// case COMPONENT_TYPE_STRING:
// case COMPONENT_TYPE_BLOB:
// case COMPONENT_TYPE_INT:
// case COMPONENT_TYPE_UINT:
// case COMPONENT_TYPE_FLOAT:
// case COMPONENT_TYPE_BOOL:
// case COMPONENT_TYPE_UNDEFINED:
// default:
// }

type ComponentMap map[string]Component

type Component interface {
	Meta() ComponentMetadata
}

type ComponentType string

type ComponentMetadata struct {
	Type        ComponentType
	ComponentId string
	ParentId    string
	Traits      *TraitsMap
}

type ComponentObject struct {
	ComponentMetadata
	Children []*Component
}

type ComponentReference struct {
	ComponentMetadata
}

type ComponentArray struct {
	ComponentMetadata
}

type ComponentEnum struct {
	ComponentMetadata
}

type ComponentString struct {
	ComponentMetadata
}

type ComponentBlob struct {
	ComponentMetadata
}

type ComponentInt struct {
	ComponentMetadata
}

type ComponentUint struct {
	ComponentMetadata
}

type ComponentFloat struct {
	ComponentMetadata
}

type ComponentBool struct {
	ComponentMetadata
}

type ComponentUndefined struct {
	ComponentMetadata
}

// switch v := c.(type) {
// case ComponentObject:
// case ComponentReference:
// case ComponentArray:
// case ComponentEnum:
// case ComponentString:
// case ComponentBlob:
// case ComponentInt:
// case ComponentUint:
// case ComponentFloat:
// case ComponentBool:
// case ComponentUndefined:
// default:
// }

func NewComponent(
	ctype ComponentType,
	id, parentId string,
) Component {
	meta := ComponentMetadata{
		Type:        ctype,
		ComponentId: id,
		ParentId:    parentId,
		Traits:      new(TraitsMap),
	}
	componentTypesMap := map[ComponentType]Component{
		COMPONENT_TYPE_OBJECT:    ComponentObject{ComponentMetadata: meta},
		COMPONENT_TYPE_ENUM:      ComponentEnum{ComponentMetadata: meta},
		COMPONENT_TYPE_ARRAY:     ComponentArray{ComponentMetadata: meta},
		COMPONENT_TYPE_REFERENCE: ComponentReference{ComponentMetadata: meta},
		COMPONENT_TYPE_STRING:    ComponentString{ComponentMetadata: meta},
		COMPONENT_TYPE_BLOB:      ComponentBlob{ComponentMetadata: meta},
		COMPONENT_TYPE_INT:       ComponentInt{ComponentMetadata: meta},
		COMPONENT_TYPE_UINT:      ComponentUint{ComponentMetadata: meta},
		COMPONENT_TYPE_FLOAT:     ComponentFloat{ComponentMetadata: meta},
		COMPONENT_TYPE_BOOL:      ComponentBool{ComponentMetadata: meta},
	}

	comp, ok := componentTypesMap[ctype]
	if ok {
		return comp
	}

	return nil
}

func (cm ComponentMetadata) Meta() ComponentMetadata {
	return cm
}
