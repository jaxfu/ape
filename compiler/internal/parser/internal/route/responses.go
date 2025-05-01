package route

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/parser/internal/body"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"

	"github.com/jaxfu/ape/components"
)

func ParseResponses(rawResponses map[string]any) (*ParsedResponsesMap, error) {
	responses := ParsedResponsesMap{}

	for k, v := range rawResponses {
		asMap, ok := v.(map[string]any)
		asMap["name"] = k
		if !ok {
			return nil, fmt.Errorf("invalid response format for %s: %+v", k, v)
		}

		parsedRes, err := ParseResponse(
			asMap,
			false,
		)
		if err != nil {
			return nil, fmt.Errorf("Parser.ParseResponse: %+v", err)
		}

		responses[k] = parsedRes
	}

	return &responses, nil
}

func ParseResponse(rawResponse map[string]any, isRoot bool) (ParsedResponse, error) {
	var statusCode *uint = nil
	if rawStatusCode, ok := rawResponse[shared.KEY_STATUS_CODE]; ok {
		asInt64, ok := rawStatusCode.(int64)
		if !ok {
			return ParsedResponse{}, fmt.Errorf("invalid status code format: %+v", rawStatusCode)
		}
		asUint := uint(asInt64)
		statusCode = &asUint
	}

	metadata, err := shared.ParseComponentMetadata(
		rawResponse,
		components.ComponentTypes.Types().RESPONSE,
		isRoot,
	)
	if err != nil {
		return ParsedResponse{}, fmt.Errorf("Parser.ParseComponentMetadata: %+v", err)
	}

	parsedBody, err := body.ParseMessageBody(rawResponse, false)
	if err != nil {
		return ParsedResponse{}, fmt.Errorf("Parser.ParseMessageBody: %+v", err)
	}

	return ParsedResponse{
		Metadata:   metadata,
		StatusCode: statusCode,
		Body:       parsedBody,
		Context: shared.Context{
			ComponentType: components.ComponentTypes.Types().RESPONSE,
			Name:          metadata.Name,
			IsRoot:        isRoot,
		},
	}, nil
}

type ParsedResponse struct {
	Metadata   shared.ParsedComponentMetadata
	StatusCode *uint
	Body       *body.ParsedMessageBody
	Context    shared.Context
}

type ParsedResponsesMap = map[string]ParsedResponse
