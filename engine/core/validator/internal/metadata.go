package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateComponentMetadata(meta components.ComponentMetadata) error {
	if _, err := components.ParsePropType(string(meta.ComponentType)); err != nil {
		return fmt.Errorf("invalid component type %s", meta.ComponentType)
	}
	if v := isValid(meta.ComponentId); !v {
		return fmt.Errorf("component id empty")
	}
	if v := isValid(meta.Name); !v {
		return fmt.Errorf("name empty")
	}

	return nil
}
