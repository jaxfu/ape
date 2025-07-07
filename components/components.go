package components

const (
	COMPTYPE_STANDARD  ComponentType = "STANDARD"
	COMPTYPE_REFERENCE ComponentType = "REFERENCE"
	COMPTYPE_ENUM      ComponentType = "ENUM"
	COMPTYPE_ARRAY     ComponentType = "ARRAY"
	COMPTYPE_UNDEFINED ComponentType = "UNDEFINED"
)

// switch ctype {
// case COMPTYPE_STANDARD:
// case COMPTYPE_REFERENCE:
// case COMPTYPE_ENUM:
// case COMPTYPE_ARRAY:
// case COMPTYPE_UNDEFINED:
// }

type ComponentMap map[string]Component

type Component any

type ComponentType string

type ComponentTypesConstraint interface {
	ComponentStandard |
		ComponentReference |
		ComponentEnum |
		ComponentArray
}

// switch comp.(type) {
// case ComponentStandard:
// case ComponentReference:
// case ComponentEnum:
// case ComponentArray:
// }

type ComponentMetadata struct {
	Type        ComponentType
	ComponentId string
	ParentId    string
}

type ComponentStandard struct {
	Metadata    ComponentMetadata
	Constraints map[string]Constraint
}

type ComponentReference struct {
	Metadata ComponentMetadata
}

type EnumMember struct {
	EnumId string
	Key    string
}

type ComponentEnum struct {
	Metadata ComponentMetadata
	Members  map[string]EnumMember
}

type ComponentArray struct {
	Metadata ComponentMetadata
}

func NewComponent[T ComponentTypesConstraint](
	ctype ComponentType,
	id, parentId string,
) Component {
	meta := ComponentMetadata{
		Type:        ctype,
		ComponentId: id,
		ParentId:    parentId,
	}

	switch ctype {
	case COMPTYPE_STANDARD:
		return ComponentStandard{
			Metadata:    meta,
			Constraints: map[string]Constraint{},
		}
	case COMPTYPE_REFERENCE:
		return ComponentReference{
			Metadata: meta,
		}
	case COMPTYPE_ENUM:
		return ComponentEnum{
			Metadata: meta,
			Members:  map[string]EnumMember{},
		}
	case COMPTYPE_ARRAY:
		return ComponentArray{
			Metadata: meta,
		}
	default:
		return nil
	}
}
