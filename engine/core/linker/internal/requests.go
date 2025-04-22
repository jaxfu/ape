package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler"
)

func (l *Linker) linkRequests(routes map[string]compiler.CompiledRequest) (
	map[components.ComponentId]components.Request,
	error,
) {
	linked := map[components.ComponentId]components.Request{}

	if len(routes) > 0 {
		for k, v := range routes {
			request, err := l.linkRequest(v)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkRequest: %+v", err)
			}

			linked[k] = request
		}
	}

	return linked, nil
}

func (l *Linker) linkRequest(request compiler.CompiledRequest) (components.Request, error) {
	meta, err := l.LinkComponent(request.ComponentMetadata)
	if err != nil {
		return components.Request{}, fmt.Errorf("Linker.LinkComponent: %+v", err)
	}

	var body *components.MessageBody = nil
	if request.Body != nil {
		linked, err := l.linkBody(*request.Body)
		if err != nil {
			return components.Request{}, fmt.Errorf("Linker.linkBody: %+v", err)
		}
		body = &linked
	}

	return components.Request{
		ComponentMetadata: meta,
		Headers:           request.Headers,
		Body:              body,
	}, nil
}
