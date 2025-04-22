package internal

import (
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/object"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/route"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
)

type Assembler struct{}

func DefaultAssembler() *Assembler {
	return &Assembler{}
}

func (asm *Assembler) AssembleProp(parsedProp parser.ParsedProp) (shared.CompiledProp, error) {
	return props.AssembleProp(parsedProp)
}

func (asm *Assembler) AssembleObject(parsedObj parser.ParsedObject) (shared.CompiledObject, error) {
	return object.AssembleObject(parsedObj)
}

func (asm *Assembler) AssembleRoute(parsedRoute parser.ParsedRoute) (shared.CompiledRoute, error) {
	return route.AssembleRoute(parsedRoute)
}
