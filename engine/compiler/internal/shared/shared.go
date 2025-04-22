package shared

import "github.com/jaxfu/ape/components"

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

func (ccm CompiledComponentMetadata) Id() string {
	return ccm.ComponentId
}

type CompilationContext struct {
	ComponentType components.ComponentType
	Name          *string
	IsRoot        bool
	ParentId      *string
}
