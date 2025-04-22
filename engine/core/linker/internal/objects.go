package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler"
)

func (l *Linker) linkObjects(objects map[string]compiler.CompiledObject) (map[components.ComponentId]components.Object, error) {
	linked := map[components.ComponentId]components.Object{}

	if len(objects) > 0 {
		for k, v := range objects {
			meta, err := l.LinkComponent(v.ComponentMetadata)
			if err != nil {
				return nil, fmt.Errorf("Linker.LinkComponent: %+v", err)
			}

			props, err := l.linkProps(v.Props)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkProps: %+v", err)
			}

			linked[k] = components.Object{
				ComponentMetadata: meta,
				Props:             props,
			}
		}
	}

	return linked, nil
}
