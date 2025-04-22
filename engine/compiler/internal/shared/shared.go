package shared

import "github.com/jaxfu/ape/components"

type CompiledComponents struct {
	Props     []CompiledProp
	Objects   []CompiledObject
	Routes    []CompiledRoute
	Bodies    []CompiledBody
	Requests  []CompiledRequest
	Responses []CompiledResponse
}

type CompiledComponentMetadata struct {
	ComponentType components.ComponentType
	Name          string
	ComponentId   string
	IsRoot        bool
	Category      *string
	Description   *string
}

type CompilationContext struct {
	ComponentType components.ComponentType
	Name          *string
	IsRoot        bool
	ParentId      *string
}
