package store

import (
	"fmt"

	"github.com/jaxfu/ape/engine/core/bus"
	"github.com/jaxfu/ape/engine/core/store/internal/categories"
	"github.com/jaxfu/ape/engine/core/store/internal/components"
)

type Store struct {
	Components ComponentStore
	Categories CategoryStore
	Events     <-chan bus.Event
}

type (
	ComponentStore = components.ComponentStore
	CategoryStore  = categories.CategoryStore
)

func NewStore(chin <-chan bus.Event) *Store {
	return &Store{
		Components: components.NewComponentStore(),
		Categories: categories.NewCategoryStore(),
		Events:     chin,
	}
}

func (s *Store) Start() {
	for event := range s.Events {
		fmt.Printf("store: %+v\n", event)
	}
}
