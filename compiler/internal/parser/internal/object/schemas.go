package object

import (
	"github.com/jaxfu/ape/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
)

type ParsedObject struct {
	ComponentMetadata shared.ParsedComponentMetadata
	Props             prop.ParsedProps
	Context           shared.Context
}
