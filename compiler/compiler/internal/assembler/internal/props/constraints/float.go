package constraints

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

const PRECISION string = "precision"

func assembleFloatConstraints(fields map[string]any) (components.PropConstraints, error) {
	constraints := components.PropConstraintsFloat{}

	precision, exists, err := extractFromMap[string](PRECISION, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", PRECISION, err)
	} else if exists {
		constraints.Precision = &precision
	}

	min, exists, err := extractFromMap[float64](MIN, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MIN, err)
	} else if exists {
		constraints.Min = &min
	}

	max, exists, err := extractFromMap[float64](MAX, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MAX, err)
	} else if exists {
		constraints.Max = &max
	}

	return constraints, nil
}
