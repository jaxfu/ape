package categories

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal/categories/internal"
)

type CategoryStore interface {
	Store(components.CategoryId)
	Exists(components.ComponentId) bool
	All() CategoriesMap
}

func NewCategoryStore() CategoryStore {
	return internal.NewCategoryStore()
}

type CategoriesMap = internal.CategoriesMap
