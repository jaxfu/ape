package compiler

import (
	"github.com/jaxfu/ape/compiler/internal"
	"github.com/jaxfu/ape/compiler/internal/shared"
)

type Compiler interface {
	File(string, []byte) (CompiledComponents, error)
}

func NewCompiler() Compiler {
	return internal.DefaultCompiler()
}

type (
	CompiledComponents        = shared.CompiledComponents
	CompiledComponentMetadata = shared.CompiledComponentMetadata
	CompiledProp              = shared.CompiledProp
	CompiledObject            = shared.CompiledObject
	CompiledRoute             = shared.CompiledRoute
	CompiledRouteMetadata     = shared.CompiledRouteMetadata
	CompiledBody              = shared.CompiledBody
	CompiledRequest           = shared.CompiledRequest
	CompiledResponse          = shared.CompiledResponse
)
