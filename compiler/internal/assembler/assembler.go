package assembler

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/golp/stack"
)

const (
	COMPS_PREALLOC uint = 64
	STACK_PREALLOC uint = 16
)

type AssemblyCtx struct {
	Comps components.ComponentMap
	Stack stack.Stack[string]
	Error error
}

type ComponentStack stack.Stack[string]

// TODO: wip assembler
func Assemble(ast shared.Ast) (components.ComponentMap, error) {
	ctx := AssemblyCtx{
		Comps: components.ComponentMap{},
		Stack: stack.Stack[string]{},
	}

	for _, node := range ast {
		ctx = processNode(node, ctx)
		if ctx.Error != nil {
			return nil, fmt.Errorf(
				"ctx:%+v\nnode: %+v\n%+v\n",
				ctx,
				node,
				ctx.Error,
			)
		}
	}

	return ctx.Comps, nil
}
