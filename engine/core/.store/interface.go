package store

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal"
	"github.com/jaxfu/ape/engine/core/store/repo"
)

type StoreInterface interface {
	Add(any) error
	Get(string) (any, error)
	GetObjects() ([]components.Object, error)
	Sync(components.AllComponents) error
}

func NewStore(r repo.Repository) StoreInterface {
	return internal.DefaultStore(r)
}
