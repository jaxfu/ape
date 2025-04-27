package internal

import (
	"context"
	"encoding/json"
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

func (eh *EventProcessor) Start(ctx context.Context) {
	for {
		select {
		case event := <-eh.Events:
			fmt.Printf("processor: %+v\n", event.EventType)
			switch event.EventType {
			case EventTypes.CREATE_COMPONENT:
				if err := eh.createComponent(event.Component); err != nil {
					event.ResultChan <- Result{Error: fmt.Errorf("EventProcessor.createComponent: %+v", err)}
					break
				}
				event.ResultChan <- Result{Error: nil}
			case EventTypes.GET_COMPONENTS:
				fmt.Println("GET_COMPONENTS")

				comps := []components.Component{}
				for _, v := range eh.Store.Components {
					comps = append(comps, v)
				}
				marshalled, err := json.Marshal(comps)
				if err != nil {
					event.ResultChan <- Result{Error: fmt.Errorf("json.Marshal: %+v", err)}
					break
				}
				event.ResultChan <- Result{Bytes: marshalled}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (eh *EventProcessor) createComponent(comp components.Component) error {
	if err := eh.Store.CreateComponent(comp); err != nil {
		return fmt.Errorf("error storing component: %+v", err)
	}

	return nil
}

type Event struct {
	EventType  EventType
	Component  components.Component
	ResultChan chan Result
}

type Result struct {
	Bytes []byte
	Error error
}

type (
	EventType           = string
	EventTypesInterface struct {
		CREATE_COMPONENT EventType
		GET_COMPONENTS   EventType
	}
)

var EventTypes = EventTypesInterface{
	CREATE_COMPONENT: "CREATE_COMPONENT",
	GET_COMPONENTS:   "GET_COMPONENTS",
}
