package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateRequest(req components.Request) error {
	if err := v.ValidateComponent(req.Body); err != nil {
		return fmt.Errorf("error validating body on request %s: %+v", req.ComponentId, err)
	}

	return nil
}
