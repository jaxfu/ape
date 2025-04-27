package internal

import (
	"github.com/jaxfu/ape/engine/core/events"
)

func DefaultApi(bus *events.Bus) *Api {
	return &Api{
		Bus: bus,
	}
}

type Api struct {
	Bus *events.Bus
}
