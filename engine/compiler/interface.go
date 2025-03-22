package compiler

import (
	"github.com/jaxfu/ape/engine/compiler/internal"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
)

type Compiler interface {
	File(string, []byte) (CompiledComponents, error)
}

func NewCompiler() Compiler {
	return internal.DefaultCompiler()
}

type (
	CompiledComponents = shared.CompiledComponents
	CompiledProp       = shared.CompiledProp
	CompiledObject     = shared.CompiledObject
	CompiledRoute      = shared.CompiledRoute
	CompiledBody       = shared.CompiledBody
	CompiledRequest    = shared.CompiledRequest
	CompiledResponse   shared.CompiledResponse
)
