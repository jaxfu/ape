package internal

import (
	"database/sql"

	"github.com/jaxfu/ape/engine/core/store/repo/db/sql/generated"
)

type Repository struct {
	QueryClient *generated.Queries
}

func DefaultRepository(db *sql.DB) *Repository {
	return &Repository{
		QueryClient: generated.New(db),
	}
}

func (rp *Repository) GetQueryClient() *generated.Queries {
	return rp.QueryClient
}
