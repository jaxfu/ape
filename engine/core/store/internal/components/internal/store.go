package internal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jaxfu/ape/components"
)

func NewComponentStore() *ComponentStore {
	return &ComponentStore{
		Components: ComponentsMap{},
		Mutex:      new(sync.Mutex),
	}
}

type ComponentsMap = map[components.ComponentId]components.Component

type ComponentStore struct {
	Mutex      *sync.Mutex
	Components ComponentsMap
}

func (cs *ComponentStore) Store(comp components.Component) error {
	id := strings.ToLower(comp.Metadata().ComponentId)

	cs.Mutex.Lock()
	component, err := DeepCopy(&comp)
	if err != nil {
		return fmt.Errorf("ComponentStore.DeepCopy for %s: %+v", comp.Metadata().ComponentId, err)
	}
	cs.Components[id] = component
	cs.Mutex.Unlock()

	return nil
}

func (cs *ComponentStore) Exists(id components.ComponentId) bool {
	id = strings.ToLower(id)
	cs.Mutex.Lock()
	_, ok := cs.Components[id]
	cs.Mutex.Unlock()
	return ok
}
