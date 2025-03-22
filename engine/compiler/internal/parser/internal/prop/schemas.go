package prop

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
)

type ParsedProp struct {
	ComponentMetadata shared.ParsedComponentMetadata
	PropMetadata      ParsedPropMetadata
	Constraints       map[string]any
	Context           components.ComponentContext
}
type ParsedProps = map[string]ParsedProp

type ParsedPropMetadata struct {
	PropType string
	IsArray  *bool
}
