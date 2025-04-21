package objects

import (
	"github.com/jaxfu/ape/components/internal/props"
	"github.com/jaxfu/ape/components/internal/shared"
)

type Object struct {
	ComponentMetadata shared.ComponentMetadata
	Props             props.PropsMap `json:"props,omitempty" toml:"props,omitempty"`
}

func (obj Object) GetMetadata() shared.ComponentMetadata {
	return obj.ComponentMetadata
}
