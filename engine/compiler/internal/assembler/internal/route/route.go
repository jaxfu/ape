package route

import (
	"fmt"

	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	compshared "github.com/jaxfu/ape/engine/compiler/internal/shared"
)

func AssembleRoute(parsedRoute parser.ParsedRoute) (compshared.CompiledRoute, error) {
	metadata, err := shared.AssembleComponentMetadata(
		parsedRoute.ComponentMetadata,
		parsedRoute.Context,
	)
	if err != nil {
		return compshared.CompiledRoute{}, fmt.Errorf("shared.AssembleComponentMetadata: %+v", err)
	}

	routeMetadata, err := assembleRouteMetadata(parsedRoute.RouteMetadata)
	if err != nil {
		return compshared.CompiledRoute{}, fmt.Errorf("Assembler.AssembleRouteMetadata: %+v", err)
	}

	request := compshared.CompiledRequest{}
	if parsedRoute.Request != nil {
		parsedRoute.Request.Context.ParentId = &metadata.ComponentId

		var err error
		request, err = AssembleRequest(*parsedRoute.Request)
		if err != nil {
			return compshared.CompiledRoute{}, fmt.Errorf("Assembler.AssembleRequest: %+v", err)
		}
	}

	responses, err := AssembleResponses(*parsedRoute.Responses, metadata.ComponentId)
	if err != nil {
		return compshared.CompiledRoute{}, fmt.Errorf("Assembler.AssembleResponses: %+v", err)
	}

	return compshared.CompiledRoute{
		ComponentMetadata: metadata,
		RouteMetadata:     routeMetadata,
		Request:           request,
		Responses:         responses,
	}, nil
}

func assembleRouteMetadata(metadata parser.ParsedRouteMetadata) (compshared.CompiledRouteMetadata, error) {
	url := ""
	if metadata.Url == "" {
		return compshared.CompiledRouteMetadata{}, fmt.Errorf("no url given")
	}
	url = metadata.Url

	method := ""
	if metadata.Method != nil && *metadata.Method != "" {
		method = *metadata.Method
	}

	return compshared.CompiledRouteMetadata{
		Url:    url,
		Method: method,
	}, nil
}
