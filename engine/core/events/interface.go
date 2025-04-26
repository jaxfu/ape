package events

import (
	"github.com/jaxfu/ape/engine/core/events/internal"
	"github.com/jaxfu/ape/engine/core/store"
)

type EventProcessor interface {
	Start()
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
)

var EventTypes = internal.EventTypes
