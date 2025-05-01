package route

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/parser/internal/body"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"

	"github.com/jaxfu/ape/components"
)

func ParseRequest(routeFields map[string]any, isRoot bool) (*ParsedRequest, error) {
	request, ok := routeFields[shared.KEY_REQUEST]
	if !ok {
		return nil, fmt.Errorf("no request object found")
	}
	reqMap, ok := request.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid request format")
	}

	metadata, err := shared.ParseComponentMetadata(
		reqMap,
		components.ComponentTypes.Types().REQUEST,
		isRoot,
	)
	if err != nil {
		return nil, fmt.Errorf("Parser.ParseComponentMetadata:  %+v", err)
	}

	headers, err := parseHeaders(reqMap)
	if err != nil {
		return nil, fmt.Errorf("Parser.parseHeaders: %+v", err)
	}

	body, err := body.ParseMessageBody(reqMap, false)
	if err != nil {
		return nil, fmt.Errorf("Parser.ParseMessageBody: %+v", err)
	}

	ctx := shared.Context{
		ComponentType: components.ComponentTypes.Types().REQUEST,
		Name:          metadata.Name,
		IsRoot:        isRoot,
	}
	return &ParsedRequest{
		ComponentMetadata: metadata,
		Headers:           headers,
		Body:              body,
		Context:           ctx,
	}, nil
}

func parseHeaders(request map[string]any) (*ParsedHeadersMap, error) {
	rawHeaders, ok := request[shared.KEY_HEADERS]
	if !ok {
		return nil, nil
	}
	headersMap, ok := rawHeaders.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("incorrect headers format: %+v", headersMap)
	}

	parsedMap := map[string]string{}
	for k, v := range headersMap {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("invalid header value fo %s: %+v", k, v)
		}
		parsedMap[k] = str
	}

	return &parsedMap, nil
}

type ParsedRequest struct {
	ComponentMetadata shared.ParsedComponentMetadata
	Headers           *ParsedHeadersMap
	Body              *body.ParsedMessageBody
	Context           shared.Context
}

type ParsedHeadersMap = map[string]string
