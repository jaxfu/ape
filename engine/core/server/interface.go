package server

import (
	"context"

	"github.com/jaxfu/ape/engine/core/events"
	"github.com/jaxfu/ape/engine/core/server/internal"
)

type Server interface {
	Start(ctx context.Context) error
}

func NewServer(url string, port uint, clientDirFp string, bus *events.Bus) (Server, error) {
	return internal.NewServer(url, port, clientDirFp, bus)
}
