package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler"
)

// TODO: wip link route
func (l *Linker) linkRoutes(routes map[string]compiler.CompiledRoute) (
	map[components.ComponentId]components.Route,
	error,
) {
	linked := map[components.ComponentId]components.Route{}

	if len(routes) > 0 {
		for k, v := range routes {
			meta, err := l.LinkComponent(v.ComponentMetadata)
			if err != nil {
				return nil, fmt.Errorf("Linker.LinkComponent: %+v", err)
			}
			routeMeta, err := l.linkRouteMetadata(v.RouteMetadata)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkRouteMetadata: %+v", err)
			}
			request, err := l.linkRequest(v.Request)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkRequest: %+v", err)
			}
			responses, err := l.linkResponses(v.Responses)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkResponses: %+v", err)
			}

			linked[k] = components.Route{
				ComponentMetadata: meta,
				RouteMetadata:     routeMeta,
				Request:           request,
				Responses:         responses,
			}
		}
	}

	return linked, nil
}

// TODO: link url/method
func (l *Linker) linkRouteMetadata(meta compiler.CompiledRouteMetadata) (components.RouteMetadata, error) {
	return components.RouteMetadata{
		Url:    meta.Url,
		Method: meta.Method,
	}, nil
}
