package bus

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

const (
	BUFLEN int = 1
)

type Bus struct {
	Events     chan Event
	Dispatches Dispatches
}

func NewBus() *Bus {
	return &Bus{
		Events: NewChannel[Event](BUFLEN),
		Dispatches: Dispatches{
			Store: make(chan Event),
		},
	}
}

type Component = components.Component

func (bus *Bus) Start() {
	for event := range bus.Events {
		fmt.Printf("Bus: %+v\n", event.EventType)
		bus.CreateComponent(event)
	}
}

// TODO: wip validate/send component
func (bus *Bus) CreateComponent(event Event) error {
	bus.Dispatches.Store <- event
	return nil
}
