package shared

import "github.com/jaxfu/ape/components"

type CompiledComponents struct {
	Props     []CompiledProp
	Objects   []CompiledObject
	Routes    []components.Route
	Bodies    []CompiledBody
	Requests  []CompiledRequest
	Responses []CompiledResponse
}

type (
	CompiledProp     struct{}
	CompiledObject   struct{}
	CompiledRoute    struct{}
	CompiledBody     struct{}
	CompiledRequest  struct{}
	CompiledResponse struct{}
)
