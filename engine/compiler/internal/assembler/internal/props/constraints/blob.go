package constraints

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/pkg/dev"
)

func assembleBlobConstraints(fields map[string]any) (components.PropConstraints, error) {
	constraints := components.PropConstraintsBlob{}
	dev.PrettyPrint(fields)

	// size, exists, err := extractFromMap[int64](SIZE, fields)
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding %s: %+v", SIZE, err)
	// } else if exists {
	// 	asUint := uint(size)
	// 	constraints.Size = &asUint
	// }

	return constraints, nil
}
