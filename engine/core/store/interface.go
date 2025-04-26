package store

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal/categories"
	storecomps "github.com/jaxfu/ape/engine/core/store/internal/components"
	"github.com/jaxfu/ape/engine/core/validator"
)

type Store struct {
	Components ComponentStore
	Categories CategoryStore
}

type (
	ComponentStore = storecomps.ComponentStore
	CategoryStore  = categories.CategoryStore
)

func NewStore() *Store {
	store := Store{
		Components: storecomps.NewComponentStore(),
		Categories: categories.NewCategoryStore(),
	}

	return &store
}

func (s *Store) CreateComponent(comp components.Component) error {
	// validate
	if err := validator.NewValidator().ValidateComponent(comp); err != nil {
		return fmt.Errorf("error validating: %+v", err)
	}

	if err := s.Components.Store(comp); err != nil {
		return fmt.Errorf("ComponentStore.Store: %+v", err)
	}

	fmt.Println("component stored")
	return nil
}
