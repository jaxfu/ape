package internal

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/object"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/route"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

type Assembler struct{}

func DefaultAssembler() *Assembler {
	return &Assembler{}
}

func (asm *Assembler) AssembleProp(parsedProp parser.ParsedProp) (components.Prop, error) {
	return props.AssembleProp(parsedProp)
}

func (asm *Assembler) AssembleObject(parsedObj parser.ParsedObject) (components.Object, error) {
	return object.AssembleObject(parsedObj)
}

func (asm *Assembler) AssembleRoute(parsedRoute parser.ParsedRoute) (components.Route, error) {
	return route.AssembleRoute(parsedRoute)
}
