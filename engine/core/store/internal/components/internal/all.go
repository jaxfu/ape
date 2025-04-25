package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (cs *ComponentStore) All() (components.Components, error) {
	comps := components.Components{}

	cs.Mutex.Lock()
	for k, v := range cs.Components {
		cpy, err := DeepCopy(&v)
		if err != nil {
			return comps, fmt.Errorf("ComponentStore.DeepCopy for %s: %+v", k, err)
		}

		comps[k] = cpy
	}
	cs.Mutex.Unlock()

	return comps, nil
}
