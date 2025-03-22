package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/generators/internal/langs/typescript"
)

type Generators struct{}

func DefaultGenerators() *Generators {
	return &Generators{}
}

func (g *Generators) Typescript(components components.AllComponents) ([]byte, error) {
	bytes, err := typescript.Typescript(components)
	if err != nil {
		return []byte{}, fmt.Errorf("Generators.Typescript: %+v", err)
	}

	return bytes, nil
}
