package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateResponses(resps components.ResponsesMap) error {
	for _, r := range resps {
		if err := v.ValidateComponent(r); err != nil {
			return fmt.Errorf("error validating response %s", r.ComponentId)
		}
	}

	return nil
}

func (v Validator) validateResponse(resp components.Response) error {
	if err := v.ValidateComponent(resp.Body); err != nil {
		return fmt.Errorf("error validating body on response %s: %+v", resp.ComponentId, err)
	}

	return nil
}
