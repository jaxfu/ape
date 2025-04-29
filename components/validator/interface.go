package validator

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/components/validator/internal"
)

type Validator interface {
	ValidateComponent(components.Component) error
}

func NewValidator() Validator {
	return internal.NewValidator()
}
