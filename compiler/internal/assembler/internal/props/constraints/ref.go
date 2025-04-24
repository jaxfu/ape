package constraints

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

const TARGET string = "target"

func assembleRefConstraints(fields map[string]any) (components.PropConstraints, error) {
	refStr, exists, err := extractFromMap[string](TARGET, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", TARGET, err)
	} else if !exists {
		return nil, fmt.Errorf("ref missing target")
	}

	return components.PropConstraintsRef{
		Reference: refStr,
	}, nil
}
