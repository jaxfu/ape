package assembler

import (
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
)

type Assembler interface {
	AssembleProp(parser.ParsedProp) (shared.CompiledProp, error)
	AssembleObject(parser.ParsedObject) (shared.CompiledObject, error)
	AssembleRoute(parser.ParsedRoute) (shared.CompiledRoute, error)
}

func NewAssembler() Assembler {
	return internal.DefaultAssembler()
}
