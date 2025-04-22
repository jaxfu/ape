package internal

import (
	"maps"
	"strings"
	"sync"

	"github.com/jaxfu/ape/components"
)

func NewCategoryStore() *CategoryStore {
	return &CategoryStore{
		Categories: new(CategoriesMap),
		Mutex:      new(sync.Mutex),
	}
}

type CategoriesMap = map[components.CategoryId]struct{}

type CategoryStore struct {
	Mutex      *sync.Mutex
	Categories *CategoriesMap
}

type Component interface {
	GetMetadata() components.ComponentMetadata
}

func (cs *CategoryStore) Store(
	id components.CategoryId,
) {
	id = strings.ToLower(id)
	cs.Mutex.Lock()
	(*cs.Categories)[id] = struct{}{}
	cs.Mutex.Unlock()
}

func (cs *CategoryStore) Exists(id components.ComponentId) bool {
	id = strings.ToLower(id)
	cs.Mutex.Lock()
	_, ok := (*cs.Categories)[id]
	cs.Mutex.Unlock()
	return ok
}

func (cs *CategoryStore) All() CategoriesMap {
	categories := CategoriesMap{}

	cs.Mutex.Lock()
	maps.Copy(categories, *cs.Categories)
	cs.Mutex.Unlock()

	return categories
}
