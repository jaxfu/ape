package prop

import (
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
)

type ParsedProp struct {
	ComponentMetadata shared.ParsedComponentMetadata
	PropMetadata      ParsedPropMetadata
	Constraints       map[string]any
	Context           shared.Context
}
type ParsedProps = map[string]ParsedProp

type ParsedPropMetadata struct {
	PropType string
	IsArray  *bool
}
