package components

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal/components/internal"
)

type ComponentStore interface {
	Store(Component) error
	Load(components.ComponentId) (Component, error)
	Exists(components.ComponentId) bool
	All() (components.AllComponents, error)
}

func NewComponentStore() ComponentStore {
	return internal.NewComponentStore()
}

type Component = internal.Component
