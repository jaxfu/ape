package generators

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/generators/internal"
)

type GeneratorsInterface interface {
	Typescript(components.AllComponents) ([]byte, error)
}

func NewGenerators() GeneratorsInterface {
	return internal.DefaultGenerators()
}
