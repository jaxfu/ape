package store

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/bus"
	"github.com/jaxfu/ape/engine/core/store/internal/categories"
	internalcomps "github.com/jaxfu/ape/engine/core/store/internal/components"
)

type Store struct {
	Components ComponentStore
	Categories CategoryStore
	Events     <-chan bus.Event
}

type Manifest struct {
	Components map[string]*components.Component
}

type (
	ComponentStore = internalcomps.ComponentStore
	CategoryStore  = categories.CategoryStore
)

func NewStore(chin <-chan bus.Event) *Store {
	return &Store{
		Components: internalcomps.NewComponentStore(),
		Categories: categories.NewCategoryStore(),
		Events:     chin,
	}
}

func (s *Store) Start() {
	for event := range s.Events {
		fmt.Printf("store event recieved: %+v\n", event.EventType)
		fmt.Printf("%+v\n", event.Component)
	}
}

func (s *Store) CreateComponent(event bus.Event) {
	// validate
}
