package shared

type ComponentMetadata struct {
	ComponentType ComponentType
	ComponentId   ComponentId
	Name          string
	IsRoot        bool
	ParentId      *ComponentId
	Category      *CategoryId
	Description   *string
}

type (
	ComponentType = string
	ComponentId   = string
	CategoryId    = string
)
