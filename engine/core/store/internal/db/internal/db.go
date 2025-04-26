package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jaxfu/ape/engine/core/store/internal/db/generated"
)

func DefaultDb(dbPath, initSqlPath string) (*Db, error) {
	conn, err := initDb(dbPath, initSqlPath)
	if err != nil {
		return nil, err
	}

	return &Db{
		DbConn: conn,
		Query:  generated.New(conn),
	}, nil
}

type Db struct {
	DbConn *sql.DB
	Query  *generated.Queries
}

func (db *Db) Conn() *sql.DB {
	return db.DbConn
}

func (db *Db) InsertComponent(id string, content []byte) error {
	if err := db.Query.InsertComponentMetadata(
		context.Background(),
		generated.InsertComponentMetadataParams{
			ComponentID: id,
			Content:     content,
		},
	); err != nil {
		return err
	}

	return nil
}

func (db *Db) GetComponent(id string) (bool, []byte, error) {
	content, err := db.Query.GetComponentByComponentId(context.Background(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, []byte{}, nil
		}
		return true, []byte{}, fmt.Errorf("Query.GetComponentByComponentId: %+v", err)
	}

	return true, content, nil
}
