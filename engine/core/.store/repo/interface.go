package repo

import (
	"database/sql"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/db/sql/generated"
	"github.com/jaxfu/ape/engine/core/store/repo/internal"
)

type Repository interface {
	StoreObject(components.Object) error
	// TODO: change to name:obj map
	GetObjects() ([]components.Object, error)
	StoreProp(components.Prop) error
	GetProps() ([]components.Prop, error)
	GetQueryClient() *generated.Queries
}

func NewRepository(db *sql.DB) Repository {
	return internal.DefaultRepository(db)
}
