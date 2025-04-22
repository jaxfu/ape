package store

import (
	"github.com/jaxfu/ape/engine/core/store/internal/categories"
	"github.com/jaxfu/ape/engine/core/store/internal/components"
)

type Store struct {
	Components components.ComponentStore
	Categories categories.CategoryStore
}

func NewStore() *Store {
	return &Store{
		Components: components.NewComponentStore(),
		Categories: categories.NewCategoryStore(),
	}
}
