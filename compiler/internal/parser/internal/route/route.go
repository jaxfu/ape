package route

import (
	"fmt"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/scanner"

	"github.com/jaxfu/ape/components"
)

func ParseRoute(scannedComp scanner.ScannedComponent, isRoot bool) (ParsedRoute, error) {
	metadata, err := shared.ParseComponentMetadata(scannedComp.Fields, components.ComponentTypes.ROUTE, isRoot)
	if err != nil {
		return ParsedRoute{}, fmt.Errorf("Parser.parseComponentMetadata: %+v", err)
	}

	routeMetadata, err := parseRouteMetadata(scannedComp.Fields)
	if err != nil {
		return ParsedRoute{}, fmt.Errorf("Parser.parseRouteMetadata: %+v", err)
	}

	request, err := ParseRequest(scannedComp.Fields, false)
	if err != nil {
		return ParsedRoute{}, fmt.Errorf("Parser.ParseRequest: %+v", err)
	}

	var responses *ParsedResponsesMap = nil
	if rawResponse, ok := scannedComp.Fields[shared.KEY_RESPONSES]; ok {
		responseMap, ok := rawResponse.(map[string]any)
		if !ok {
			return ParsedRoute{}, fmt.Errorf("invalid responses format")
		}

		responses, err = ParseResponses(responseMap)
		if err != nil {
			return ParsedRoute{}, fmt.Errorf("Parser.ParseResponses: %+v", err)
		}

	}

	return ParsedRoute{
		ComponentMetadata: metadata,
		RouteMetadata:     routeMetadata,
		Request:           request,
		Responses:         responses,
		Context: shared.Context{
			ComponentType: components.ComponentTypes.ROUTE,
			Name:          metadata.Name,
			IsRoot:        isRoot,
		},
	}, nil
}

type ParsedRoute struct {
	ComponentMetadata shared.ParsedComponentMetadata
	RouteMetadata     ParsedRouteMetadata
	Responses         *ParsedResponsesMap
	Request           *ParsedRequest
	Context           shared.Context
}
