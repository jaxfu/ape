package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (cs *ComponentStore) All() (components.AllComponents, error) {
	comps := components.AllComponents{
		Props:         map[components.ComponentId]components.Prop{},
		Objects:       map[components.ComponentId]components.Object{},
		Routes:        map[components.ComponentId]components.Route{},
		Requests:      map[components.ComponentId]components.Request{},
		Response:      map[components.ComponentId]components.Response{},
		MessageBodies: map[components.ComponentId]components.MessageBody{},
	}

	cs.Mutex.Lock()
	for k, v := range cs.Components {
		cpy, err := DeepCopy(&v)
		if err != nil {
			return comps, fmt.Errorf("ComponentStore.DeepCopy for %s: %+v", k, err)
		}

		switch v.GetMetadata().ComponentType {
		case components.COMPONENT_TYPE_PROP:
			cast, err := cast[components.Prop](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.Props[k] = cast
		case components.COMPONENT_TYPE_OBJECT:
			cast, err := cast[components.Object](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.Objects[k] = cast
		case components.COMPONENT_TYPE_ROUTE:
			cast, err := cast[components.Route](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.Routes[k] = cast
		case components.COMPONENT_TYPE_MESSAGE_BODY:
			cast, err := cast[components.MessageBody](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.MessageBodies[k] = cast
		case components.COMPONENT_TYPE_REQUEST:
			cast, err := cast[components.Request](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.Requests[k] = cast
		case components.COMPONENT_TYPE_RESPONSE:
			cast, err := cast[components.Response](cpy)
			if err != nil {
				return comps, fmt.Errorf("error casting for %s: %+v", k, err)
			}
			comps.Response[k] = cast
		}
	}
	cs.Mutex.Unlock()

	return comps, nil
}

func cast[D any](src any) (D, error) {
	cast, ok := src.(D)
	if !ok {
		return *new(D), fmt.Errorf("could not cast")
	}
	return cast, nil
}
