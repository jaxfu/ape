package route

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

func AssembleRoute(parsedRoute parser.ParsedRoute) (components.Route, error) {
	metadata, err := shared.AssembleComponentMetadata(
		parsedRoute.ComponentMetadata,
		parsedRoute.Context,
	)
	if err != nil {
		return components.Route{}, fmt.Errorf("shared.AssembleComponentMetadata: %+v", err)
	}

	routeMetadata, err := assembleRouteMetadata(parsedRoute.RouteMetadata)
	if err != nil {
		return components.Route{}, fmt.Errorf("Assembler.AssembleRouteMetadata: %+v", err)
	}

	request := components.Request{}
	if parsedRoute.Request != nil {
		var err error
		parsedRoute.Request.Context = components.ComponentContext{
			ComponentType: components.COMPONENT_TYPE_REQUEST,
			IsRoot:        false,
			ParentId:      &metadata.ComponentId,
		}
		request, err = AssembleRequest(*parsedRoute.Request)
		if err != nil {
			return components.Route{}, fmt.Errorf("Assembler.AssembleRequest: %+v", err)
		}
	}

	responses, err := AssembleResponses(*parsedRoute.Responses, metadata.ComponentId)
	if err != nil {
		return components.Route{}, fmt.Errorf("Assembler.AssembleResponses: %+v", err)
	}

	return components.Route{
		ComponentMetadata: metadata,
		RouteMetadata:     routeMetadata,
		Request:           request,
		Responses:         responses,
	}, nil
}

func assembleRouteMetadata(metadata parser.ParsedRouteMetadata) (components.RouteMetadata, error) {
	url := ""
	if metadata.Url == "" {
		return components.RouteMetadata{}, fmt.Errorf("no url given")
	}
	url = metadata.Url

	method := ""
	if metadata.Method != nil && *metadata.Method != "" {
		method = *metadata.Method
	}

	return components.RouteMetadata{
		Url:    url,
		Method: method,
	}, nil
}
