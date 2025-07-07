package assembler

import (
	"github.com/jaxfu/ape/compiler/internal/shared"
)

func processNode(node shared.Node, ctx AssemblyCtx) AssemblyCtx {
	switch n := node.(type) {
	case shared.NodeComponent:
		ctx = processNodeComponent(n, ctx)
	case shared.NodeEnumMember:
		ctx = processNodeEnumMember(n, ctx)
	case shared.NodeConstraint:
		ctx = processNodeConstraint(n, ctx)
	default:
		ctx.Error = shared.NewSyntaxError(node.Meta().Position, "unrecognized node type")
	}

	return ctx
}

func processNodeComponent(node shared.NodeComponent, ctx AssemblyCtx) AssemblyCtx {
	// 	ctype := classifyComponentNode(node)
	// 	id := strings.ToLower(node.Name)
	// 	comp := components.NewComponent(ctype, id)
	// 	ctx.Comps[comp.Meta().ComponentId] = comp
	// 	ctx.Stack.Push(comp.Meta().ComponentId)
	//
	// 	return ctx
	// }
	//
	// func processNodeEnumMember(node shared.NodeEnumMember, ctx AssemblyCtx) AssemblyCtx {
	// 	id, ln := ctx.Stack.Curr()
	// 	if ln <= 0 {
	// 		ctx.Error = shared.NewSyntaxError(node.Meta().Position, "orphaned enum member")
	// 		return ctx
	// 	}
	//
	// 	parent, ok := ctx.Comps[id]
	// 	if ok &&
	// 		parent.Meta().Type == components.COMPTYPE_ENUM {
	// 	} else {
	// 		ctx.Error = shared.NewSyntaxError(node.Meta().Position, "orphaned enum member")
	// 		return ctx
	// 	}

	return ctx
}

func processNodeConstraint(node shared.NodeConstraint, ctx AssemblyCtx) AssemblyCtx {
	return ctx
}

func processNodeEnumMember(node shared.NodeEnumMember, ctx AssemblyCtx) AssemblyCtx {
	return ctx
}
