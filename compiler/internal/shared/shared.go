package shared

import "github.com/jaxfu/ape/components"

type CompiledComponent interface {
	Metadata() CompiledComponentMetadata
}

type CompiledComponents struct {
	Props     map[string]CompiledProp
	Objects   map[string]CompiledObject
	Routes    map[string]CompiledRoute
	Bodies    map[string]CompiledBody
	Requests  map[string]CompiledRequest
	Responses map[string]CompiledResponse
}

type CompiledComponentMetadata struct {
	ComponentType components.ComponentType
	Name          string
	ComponentId   string
	IsRoot        bool
	ParentId      *string
	Category      *string
	Description   *string
}

func (meta CompiledComponentMetadata) Metadata() CompiledComponentMetadata {
	return meta
}

type CompilationContext struct {
	ComponentType components.ComponentType
	Name          *string
	IsRoot        bool
	ParentId      *string
}
