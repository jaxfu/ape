package props

import (
	"github.com/jaxfu/ape/components/internal/props/constraints"
	"github.com/jaxfu/ape/components/internal/shared"
)

type Prop struct {
	ComponentMetadata shared.ComponentMetadata
	PropMetadata      PropMetadata
	Constraints       constraints.PropConstraints
}

func (p Prop) GetMetadata() shared.ComponentMetadata {
	return p.ComponentMetadata
}

type PropsMap map[string]Prop

type PropMetadata struct {
	PropType PropType
	IsArray  bool
}
