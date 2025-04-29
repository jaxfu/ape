package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateObject(obj components.Object) error {
	if err := v.validateProps(obj.Props); err != nil {
		return fmt.Errorf("error validating props: %+v", err)
	}

	return nil
}
