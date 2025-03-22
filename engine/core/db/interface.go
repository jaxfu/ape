package db

import (
	"database/sql"

	"github.com/jaxfu/ape/engine/core/db/internal"
)

type DbInterface interface {
	Conn() *sql.DB
	InsertComponent(id string, content []byte) error
	GetComponent(id string) (bool, string, error)
}

func NewDb(dbPath, initSqlPath string) (DbInterface, error) {
	return internal.DefaultDb(dbPath, initSqlPath)
}
