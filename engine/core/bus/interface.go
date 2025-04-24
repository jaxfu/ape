package bus

import "fmt"

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

func (bus *Bus) Start() {
	for event := range bus.Events {
		fmt.Printf("bus: %+v\n", event)
		bus.Dispatches.Store <- event
	}
}
