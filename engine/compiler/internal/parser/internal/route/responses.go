package route

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/body"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
)

func ParseResponses(rawResponses map[string]any) (*ParsedResponsesMap, error) {
	responses := ParsedResponsesMap{}

	for k, v := range rawResponses {
		asMap, ok := v.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid response format for %s: %+v", k, v)
		}

		ctx := components.ComponentContext{
			ComponentType: components.COMPONENT_TYPE_RESPONSE,
			IsRoot:        false,
			Name:          &k,
		}
		parsedRes, err := ParseResponse(asMap, ctx)
		if err != nil {
			return nil, fmt.Errorf("Parser.ParseResponse: %+v", err)
		}

		responses[k] = parsedRes
	}

	return &responses, nil
}

func ParseResponse(rawResponse map[string]any, ctx components.ComponentContext) (ParsedResponse, error) {
	var statusCode *uint = nil
	if rawStatusCode, ok := rawResponse[shared.KEY_STATUS_CODE]; ok {
		asInt64, ok := rawStatusCode.(int64)
		if !ok {
			return ParsedResponse{}, fmt.Errorf("invalid status code format: %+v", rawStatusCode)
		}
		asUint := uint(asInt64)
		statusCode = &asUint
	}

	metadata, err := shared.ParseComponentMetadata(rawResponse)
	if err != nil {
		return ParsedResponse{}, fmt.Errorf("Parser.ParseComponentMetadata: %+v", err)
	}

	bodyCtx := components.ComponentContext{
		ComponentType: components.COMPONENT_TYPE_MESSAGE_BODY,
		IsRoot:        false,
	}
	parsedBody, err := body.ParseMessageBody(rawResponse, bodyCtx)
	if err != nil {
		return ParsedResponse{}, fmt.Errorf("Parser.ParseMessageBody: %+v", err)
	}

	return ParsedResponse{
		Metadata:   metadata,
		StatusCode: statusCode,
		Body:       parsedBody,
		Context:    ctx,
	}, nil
}

type ParsedResponse struct {
	Metadata   shared.ParsedComponentMetadata
	StatusCode *uint
	Body       *body.ParsedMessageBody
	Context    components.ComponentContext
}

type ParsedResponsesMap = map[string]ParsedResponse
