package internal

import (
	"encoding/json"
	"fmt"

	"github.com/jaxfu/ape/components"
)

func NewStoredComponent(
	id string,
	compType components.ComponentType,
	content []byte,
) StoredComponent {
	return StoredComponent{
		ComponentId:   id,
		ComponentType: compType,
		Content:       content,
	}
}

type StoredComponent struct {
	ComponentId   string
	ComponentType components.ComponentType
	Content       []byte
}

func (sc StoredComponent) Info() (string, components.ComponentType, []byte) {
	return sc.ComponentId, sc.ComponentType, sc.Content
}

func (sc StoredComponent) Bind(dest any) error {
	if err := json.Unmarshal(sc.Content, dest); err != nil {
		return fmt.Errorf("json.Unmarshal: %+v", err)
	}

	return nil
}
