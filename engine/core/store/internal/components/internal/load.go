package internal

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/components"
)

func (cs *ComponentStore) Load(id components.ComponentId) (Component, error) {
	cs.Mutex.Lock()
	found, ok := cs.Components[strings.ToLower(id)]
	if !ok {
		return nil, fmt.Errorf("id %s not found", id)
	}

	comp, err := DeepCopy(&found)
	if err != nil {
		return nil, fmt.Errorf("ComponentStore.DeepCopy for %s: %+v", id, err)
	}
	cs.Mutex.Unlock()

	return comp, nil
}
