package store

import (
	"fmt"

	"github.com/jaxfu/ape/engine/core/store/internal/components"
)

type Store struct {
	Components components.ComponentStore
}

func Get[T any](store components.ComponentStore, id string) (T, error) {
	instance := new(T)

	found, err := store.GetOne(id)
	if err != nil {
		return *instance, fmt.Errorf("Store.GetOne: %+v", err)
	}

	err = found.Bind(instance)
	if err != nil {
		return *instance, fmt.Errorf("StoredComponent.Bind: %+v", err)
	}

	return *instance, nil
}

func NewStore() *Store {
	return &Store{
		Components: components.NewComponentStore(),
	}
}
