package events

import (
	"context"

	"github.com/jaxfu/ape/engine/core/events/internal"
	"github.com/jaxfu/ape/engine/core/store"
)

type EventProcessor interface {
	Start(ctx context.Context)
}

func NewEventProcessor(events <-chan Event, store *store.Store) EventProcessor {
	return internal.InitEventProcessor(events, store)
}

type Bus struct {
	Events chan Event
}

type (
	Event               = internal.Event
	EventType           = internal.EventType
	EventTypesInterface = internal.EventTypesInterface
	Result              = internal.Result
)

var EventTypes = internal.EventTypes
