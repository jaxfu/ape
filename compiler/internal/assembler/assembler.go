package assembler

import (
	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/ape/components"
)

func Assemble(ast shared.Ast) ([]components.Component, []error) {
	comps := make([]components.Component, 0, 128)
	errs := make([]error, 0, 64)

	for _, node := range ast {
		switch node.Meta().Type {
		case shared.NODETYPE_COMPONENT:
			comp, err := assembleComponent(node)
			if err != nil {
				errs = append(errs, err)
			}
			comps = append(comps, comp)
			// case shared.NODETYPE_COMMENT:
			// case shared.NODETYPE_CONSTRAINT:
			// case shared.NODETYPE_ENUM_MEMBER:
			// case shared.NODETYPE_ARRAY:
			// case shared.NODETYPE_EMPTYLINE:
			// case shared.NODETYPE_UNDEFINED:
			// default:
		}
	}

	return comps, errs
}

func assembleComponent(node shared.Node) (components.Component, error) {
	comp := components.Object{}

	return comp, nil
}
