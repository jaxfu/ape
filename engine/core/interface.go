package core

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/jaxfu/ape/engine/core/events"
	"github.com/jaxfu/ape/engine/core/server"
	"github.com/jaxfu/ape/engine/core/store"
)

const (
	DB_NAME     string = "ape.db"
	INIT_DB_SQL string = "core/store/internal/db/sql/schemas.sql"
	BASE_URL    string = "localhost"
	PORT        uint   = 5000
)

type Core struct {
	Store          *store.Store
	Server         server.Server
	Bus            events.Bus
	EventProcessor events.EventProcessor
}

func InitCore() (*Core, error) {
	initFp, err := filepath.Abs(INIT_DB_SQL)
	if err != nil {
		log.Printf("illegal filepath '%s': %+v\n", initFp, err)
	}
	store := store.NewStore(DB_NAME, initFp)
	bus := events.Bus{
		Events: make(chan events.Event, 5),
	}

	server, err := server.NewServer(
		BASE_URL,
		PORT,
		&bus,
	)
	if err != nil {
		return nil, fmt.Errorf("server.NewServer: %+v", err)
	}

	return &Core{
		Store:          store,
		Server:         server,
		Bus:            bus,
		EventProcessor: events.NewEventProcessor(bus.Events, store),
	}, nil
}

func (c *Core) Start(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.Server.Start(ctx)
		fmt.Println("server stopped")
	}()
	go func() {
		defer wg.Done()
		c.EventProcessor.Start(ctx)
		fmt.Println("event processor stopped")
	}()

	<-ctx.Done()
	fmt.Println("Shutting down core...")
	wg.Wait()
	c.cleanup()
	fmt.Println("Core Shutdown")
}

func (c *Core) cleanup() {
	c.Store.Db.Conn().Close()
}
