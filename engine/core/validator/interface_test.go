package validator

import (
	"testing"

	"github.com/jaxfu/ape/components"
)

var invalidObj components.Object = components.Object{
	ComponentMetadata: components.ComponentMetadata{},
	Props:             components.PropsMap{},
}

var validObj components.Object = components.Object{
	ComponentMetadata: components.ComponentMetadata{
		ComponentType: components.COMPONENT_TYPE_OBJECT,
		ComponentId:   "objects.TEST",
		Name:          "TEST",
		IsRoot:        true,
	},
}

func TestInterface(t *testing.T) {
	validator := NewValidator()

	if err := validator.ValidateComponent(invalidObj); err == nil {
		t.Errorf("error expected")
		return
	}

	if err := validator.ValidateComponent(validObj); err != nil {
		t.Errorf("error: %+v\n", err)
		return
	}
}
