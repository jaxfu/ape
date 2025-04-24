package internal

import (
	"fmt"
	"github.com/jaxfu/ape/compiler"

	"github.com/jaxfu/ape/components"
)

func (l *Linker) linkProps(props map[string]compiler.CompiledProp) (
	map[components.ComponentId]components.Prop,
	error,
) {
	linked := map[components.ComponentId]components.Prop{}

	if len(props) > 0 {
		for k, v := range props {
			meta, err := l.LinkComponent(v.ComponentMetadata)
			if err != nil {
				return nil, fmt.Errorf("Linker.LinkComponent: %+v", err)
			}

			linked[k] = components.Prop{
				ComponentMetadata: meta,
				PropMetadata:      v.PropMetadata,
				Constraints:       v.Constraints,
			}
		}
	}

	return linked, nil
}
