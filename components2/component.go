package components2

const (
	COMPTYPE_STANDARD  ComponentType = "COMPTYPE_STANDARD"
	COMPTYPE_ARRAY     ComponentType = "COMPTYPE_ARRAY"
	COMPTYPE_UNDEFINED ComponentType = "COMPTYPE_UNDEFINED"
)

// pointer bag
type Component struct {
	Metadata ComponentMetadata
	Standard *ComponentStandard
}

type ComponentMetadata struct {
	Type        ComponentType
	ComponentId string
	Children    map[string]*Component
	Parent      *Component
}

type ComponentStandard struct {
	Constraints []Constraint
}

type ComponentType string

func NewComponent(id string, ctype ComponentType) Component {
	return Component{
		Metadata: ComponentMetadata{
			Type:        ctype,
			ComponentId: id,
			Children:    map[string]*Component{},
			Parent:      nil,
		},
		Standard: nil,
	}
}
