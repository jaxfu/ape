package assembler

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

type Assembler interface {
	AssembleProp(parser.ParsedProp) (components.Prop, error)
	AssembleObject(parser.ParsedObject) (components.Object, error)
	AssembleRoute(parser.ParsedRoute) (components.Route, error)
}

func NewAssembler() Assembler {
	return internal.DefaultAssembler()
}
