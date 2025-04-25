package assembler

import (
	"github.com/jaxfu/ape/compiler/internal/assembler/internal"
	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/components"
)

type Assembler interface {
	AssembleProp(parser.ParsedProp) (components.Prop, error)
	AssembleObject(parser.ParsedObject) (components.Object, error)
	AssembleRoute(parser.ParsedRoute) (components.Route, error)
}

func NewAssembler() Assembler {
	return internal.DefaultAssembler()
}
