package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateComponentMetadata(meta components.ComponentMetadata) error {
	if v := isValid(meta.ComponentId); !v {
		return fmt.Errorf("component id empty")
	}
	if v := isValid(meta.Name); !v {
		return fmt.Errorf("name empty")
	}
	if !isValid(meta.ComponentType) {
		return fmt.Errorf("missing component type")
	}

	return nil
}
