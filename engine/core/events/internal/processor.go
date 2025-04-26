package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store"
)

func InitEventProcessor(events <-chan Event, store *store.Store) *EventProcessor {
	return &EventProcessor{
		Store:  store,
		Events: events,
	}
}

type EventProcessor struct {
	Store  *store.Store
	Events <-chan Event
}

func (eh *EventProcessor) Start() {
	for event := range eh.Events {
		fmt.Printf("processor: %+v", event.EventType)
		switch event.EventType {
		case EventTypes.CREATE:
			if err := eh.createComponent(event.Component); err != nil {
				fmt.Printf("EventProcessor.createComponent: %+v", err)
			}
		}
	}
}

func (eh *EventProcessor) createComponent(comp components.Component) error {
	if err := eh.Store.Components.Store(comp); err != nil {
		return fmt.Errorf("error storing component: %+v", err)
	}

	return nil
}

type Event struct {
	EventType EventType
	Component components.Component
}

type (
	EventType           = string
	EventTypesInterface struct {
		CREATE EventType
	}
)

var EventTypes = EventTypesInterface{
	CREATE: "CREATE",
}
