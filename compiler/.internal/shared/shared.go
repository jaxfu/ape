package shared

import "github.com/jaxfu/ape/components"

type CompilationContext struct {
	ComponentType components.ComponentType
	Name          *string
	IsRoot        bool
	ParentId      *string
}
