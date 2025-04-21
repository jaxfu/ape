package components

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal/components/internal"
)

type ComponentStore interface {
	Add(Component) error
	GetOne(string) (DefaultStoredComponent, error)
	GetAll() map[string]any
}

type (
	Component              = internal.Component
	DefaultStoredComponent = internal.StoredComponent
)

func NewComponentStore() ComponentStore {
	return internal.NewComponentStore()
}

type StoredComponent interface {
	Info() (string, components.ComponentType, []byte)
	Bind(dest any) error
}
