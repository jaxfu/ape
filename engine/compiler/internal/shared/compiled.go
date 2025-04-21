package shared

import "github.com/jaxfu/ape/components"

type (
	CompiledObject   struct{}
	CompiledRoute    struct{}
	CompiledRequest  struct{}
	CompiledResponse struct{}
)

type (
	CompiledProps = map[string]CompiledProp
	CompiledProp  struct {
		ComponentMetadata CompiledComponentMetadata
		PropMetadata      components.PropMetadata
		Constraints       components.PropConstraints
	}
)

type CompiledBody struct {
	ComponentMetadata CompiledComponentMetadata
	BodyType          components.MessageBodyType
	Ref               components.ReferenceTag
	Props             CompiledProps
}
