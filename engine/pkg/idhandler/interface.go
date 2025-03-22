package idhandler

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/pkg/idhandler/internal"
)

type IdHandler interface {
	Generate(GenerateIdParams) (components.ComponentId, error)
}

func NewIdHandler() IdHandler {
	return internal.DefaultIdHandler()
}

type GenerateIdParams = internal.GenerateIdParams
