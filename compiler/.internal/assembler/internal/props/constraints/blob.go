package constraints

import (
	"github.com/jaxfu/ape/components"
)

func assembleBlobConstraints(fields map[string]any) (components.PropConstraints, error) {
	constraints := components.PropConstraintsBlob{}

	// size, exists, err := extractFromMap[int64](SIZE, fields)
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding %s: %+v", SIZE, err)
	// } else if exists {
	// 	asUint := uint(size)
	// 	constraints.Size = &asUint
	// }

	return constraints, nil
}
