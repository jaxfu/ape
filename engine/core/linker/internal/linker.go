package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler"
	"github.com/jaxfu/ape/engine/core/store"
)

func NewLinker(store *store.Store) *Linker {
	return &Linker{Store: store}
}

type Linker struct {
	Store *store.Store
}

func (l *Linker) LinkAll(compiled compiler.CompiledComponents) (components.AllComponents, error) {
	props, err := l.linkProps(compiled.Props)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkProps: %+v", err)
	}
	objects, err := l.linkObjects(compiled.Objects)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkObjects: %+v", err)
	}
	routes, err := l.linkRoutes(compiled.Routes)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkRoutes: %+v", err)
	}
	bodies, err := l.linkBodies(compiled.Bodies)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkBodies: %+v", err)
	}
	requests, err := l.linkRequests(compiled.Requests)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkRequests: %+v", err)
	}
	responses, err := l.linkResponses(compiled.Responses)
	if err != nil {
		return components.AllComponents{}, fmt.Errorf("Linker.linkResponses: %+v", err)
	}

	return components.AllComponents{
		Props:         props,
		Objects:       objects,
		Routes:        routes,
		Requests:      requests,
		Response:      responses,
		MessageBodies: bodies,
	}, nil
}

func (l *Linker) LinkRef(id string) bool {
	return l.Store.Components.Exists(id)
}

func (l *Linker) LinkComponent(
	metadata compiler.CompiledComponentMetadata,
) (
	components.ComponentMetadata,
	error,
) {
	// TODO: link category
	if ok := l.Store.Components.Exists(metadata.ComponentId); ok {
		return components.ComponentMetadata{}, fmt.Errorf("id %s already exists", metadata.ComponentId)
	}

	return components.ComponentMetadata{
		ComponentType: metadata.ComponentType,
		ComponentId:   metadata.ComponentId,
		Name:          metadata.Name,
		IsRoot:        metadata.IsRoot,
		ParentId:      metadata.ParentId,
		Category:      metadata.Category,
		Description:   metadata.Description,
	}, nil
}
