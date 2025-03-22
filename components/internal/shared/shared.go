package shared

type ComponentMetadata struct {
	ComponentId   ComponentId
	ComponentType ComponentType
	Context       ComponentContext
	Name          string
	Category      *CategoryId
	Description   *string
}

type ComponentContext struct {
	ComponentType ComponentType
	IsRoot        bool
	Name          *string
	ParentId      *ComponentId
}

type CategoryId struct {
	Validated *string
	Display   string
}

type ComponentId struct {
	Validated *string
	Display   string
}

type (
	ComponentType = string
)
