package core

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jaxfu/ape/engine/core/bus"
	"github.com/jaxfu/ape/engine/core/db"
	"github.com/jaxfu/ape/engine/core/server"
	"github.com/jaxfu/ape/engine/core/store"
)

const (
	DB_NAME     string = "ape.db"
	INIT_DB_SQL string = "core/db/sql/schemas.sql"
	BASE_URL    string = "localhost"
	PORT        uint   = 5000
	CLIENT_DIR  string = "clients/web/dist"
)

type Core struct {
	Store  *store.Store
	Server server.Server
	Db     *sql.DB
	Bus    *bus.Bus
}

func InitCore() (*Core, error) {
	db, err := db.NewDb(DB_NAME, INIT_DB_SQL)
	if err != nil {
		log.Fatalf("error opening db at %s: %+v\n", DB_NAME, err)
	}

	bus := bus.NewBus()
	go bus.Start()

	clientDir, err := filepath.Abs(CLIENT_DIR)
	if err != nil {
		log.Printf("illegal filepath '%s': %+v\n", clientDir, err)
	}
	server, err := server.NewServer(
		BASE_URL,
		PORT,
		clientDir,
		bus,
	)
	if err != nil {
		return nil, fmt.Errorf("server.NewServer: %+v", err)
	}

	store := store.NewStore(bus.Dispatches.Store)
	go store.Start()

	return &Core{
		Store:  store,
		Server: server,
		Db:     db.Conn(),
		Bus:    bus,
	}, nil
}
