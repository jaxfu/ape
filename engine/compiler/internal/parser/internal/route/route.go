package route

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

func ParseRoute(scannedComp scanner.ScannedComponent, ctx components.ComponentContext) (ParsedRoute, error) {
	metadata, err := shared.ParseComponentMetadata(scannedComp.Fields)
	if err != nil {
		return ParsedRoute{}, fmt.Errorf("Parser.parseComponentMetadata: %+v", err)
	}
	ctx.Name = metadata.Name

	routeMetadata, err := parseRouteMetadata(scannedComp.Fields)
	if err != nil {
		return ParsedRoute{}, fmt.Errorf("Parser.parseRouteMetadata: %+v", err)
	}

	requestCtx := components.ComponentContext{
		ComponentType: components.COMPONENT_TYPE_REQUEST,
		IsRoot:        false,
	}
	request, err := ParseRequest(scannedComp.Fields, requestCtx)
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
		Context:           ctx,
	}, nil
}

type ParsedRoute struct {
	ComponentMetadata shared.ParsedComponentMetadata
	RouteMetadata     ParsedRouteMetadata
	Responses         *ParsedResponsesMap
	Request           *ParsedRequest
	Context           components.ComponentContext
}
