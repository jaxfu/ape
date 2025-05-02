package compiler

import (
	"github.com/jaxfu/ape/compiler/internal"
	"github.com/jaxfu/ape/components"
)

type Compiler interface {
	File(string, []byte) (components.Components, error)
}

func NewCompiler() Compiler {
	return internal.DefaultCompiler()
}
