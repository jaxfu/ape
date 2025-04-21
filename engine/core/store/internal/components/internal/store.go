package internal

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/jaxfu/ape/components"
)

func NewComponentStore() *ComponentStore {
	return &ComponentStore{
		Components: &sync.Map{},
	}
}

type ComponentStore struct {
	Components *sync.Map
}

type Component interface {
	GetMetadata() components.ComponentMetadata
}

func (cs *ComponentStore) Add(comp Component) error {
	metadata := comp.GetMetadata()
	id := strings.ToLower(metadata.ComponentId.Display)
	compType := metadata.ComponentType

	marshalled, err := json.Marshal(comp)
	if err != nil {
		return fmt.Errorf("json.Marshal: %+v", err)
	}

	cs.Components.Store(
		id,
		NewStoredComponent(id, compType, marshalled),
	)

	return nil
}

func (cs *ComponentStore) GetOne(id string) (StoredComponent, error) {
	id = strings.ToLower(id)

	found, ok := cs.Components.Load(id)
	if !ok {
		return StoredComponent{}, fmt.Errorf("no component found with id %s", id)
	}

	comp, ok := found.(StoredComponent)
	if !ok {
		return StoredComponent{}, fmt.Errorf("component has invalid format")
	}

	return comp, nil
}

func (cs *ComponentStore) GetAll() map[string]any {
	comps := map[string]any{}

	cs.Components.Range(rangeFunc(comps))

	return comps
}

func rangeFunc(comps map[string]any) func(any, any) bool {
	return func(key any, value any) bool {
		keyStr, _ := key.(string)
		comp, _ := value.(StoredComponent)
		comps[keyStr] = comp

		return true
	}
}
