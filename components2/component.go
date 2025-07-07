package components2

const (
	COMPTYPE_STANDARD  ComponentType = "COMPTYPE_STANDARD"
	COMPTYPE_ARRAY     ComponentType = "COMPTYPE_ARRAY"
	COMPTYPE_UNDEFINED ComponentType = "COMPTYPE_UNDEFINED"
)

// component interface to be used by component processors
type Component struct {
	Metadata ComponentMetadata
	Standard *ComponentStandard
}

type ComponentMetadata struct {
	Type        ComponentType
	ComponentId string
	Children    []*Component
	Parent      *Component
}

type ComponentType string

func NewComponent() Component {
	return Component{
		Metadata: NewComponentMetada(),
		Standard: nil,
	}
}

func NewComponentMetada() ComponentMetadata {
	return ComponentMetadata{
		Children: make([]*Component, 0, 16),
		Parent:   nil,
	}
}

type ComponentStandard struct {
	Constraints []Constraint
}
