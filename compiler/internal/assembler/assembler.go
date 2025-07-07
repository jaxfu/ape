package assembler

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
	components "github.com/jaxfu/ape/components2"
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

func classifyComponentNode(node shared.NodeComponent) components.ComponentType {
	if node.IsReference {
		return components.COMPTYPE_REFERENCE
	} else if strings.ToLower(node.Type) == shared.KEYWORD_DECLARE_ENUM {
		return components.COMPTYPE_ENUM
	} else if strings.ToLower(node.Type) == shared.KEYWORD_DECLARE_ARRAY {
		return components.COMPTYPE_ARRAY
	} else {
		return components.COMPTYPE_STANDARD
	}
}
