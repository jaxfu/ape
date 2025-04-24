package internal

import (
	"fmt"

	"github.com/jaxfu/ape/compiler"

	"github.com/jaxfu/ape/components"
)

func (l *Linker) linkResponses(responses map[string]compiler.CompiledResponse) (
	map[components.ComponentId]components.Response,
	error,
) {
	linked := map[components.ComponentId]components.Response{}

	if len(responses) > 0 {
		for k, v := range responses {
			response, err := l.linkResponse(v)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkResponse: %+v", err)
			}

			linked[k] = response
		}
	}

	return linked, nil
}

func (l *Linker) linkResponse(response compiler.CompiledResponse) (components.Response, error) {
	meta, err := l.LinkComponent(response.CompiledComponentMetadata)
	if err != nil {
		return components.Response{}, fmt.Errorf("Linker.LinkComponent: %+v", err)
	}

	var body *components.MessageBody = nil
	if response.Body != nil {
		linked, err := l.linkBody(*response.Body)
		if err != nil {
			return components.Response{}, fmt.Errorf("Linker.linkBody: %+v", err)
		}
		body = &linked
	}

	return components.Response{
		ComponentMetadata: meta,
		StatusCode:        response.StatusCode,
		Body:              body,
	}, nil
}
