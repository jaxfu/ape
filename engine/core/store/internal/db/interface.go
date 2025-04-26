package db

import (
	"database/sql"

	"github.com/jaxfu/ape/engine/core/store/internal/db/internal"
)

type Db interface {
	Conn() *sql.DB
	InsertComponent(id string, content []byte) error
	GetComponent(id string) (bool, []byte, error)
}

func NewDb(dbPath, initSqlPath string) (Db, error) {
	return internal.DefaultDb(dbPath, initSqlPath)
}
