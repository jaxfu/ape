package constraints

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func assembleUintConstraints(fields map[string]any) (components.PropConstraints, error) {
	constraints := components.PropConstraintsUint{}

	size, exists, err := extractFromMap[int64](SIZE, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", SIZE, err)
	} else if exists {
		asUint := uint(size)
		constraints.Size = &asUint
	}

	min, exists, err := extractFromMap[int64](MIN, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MIN, err)
	} else if exists {
		asUint := uint(min)
		constraints.Min = &asUint
	}

	max, exists, err := extractFromMap[int64](MAX, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MAX, err)
	} else if exists {
		asUint := uint(max)
		constraints.Max = &asUint
	}

	return constraints, nil
}
