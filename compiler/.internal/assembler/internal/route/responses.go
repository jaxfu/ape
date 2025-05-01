package route

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/components"
)

// TODO: assemble responses
func AssembleResponses(
	parsedRes parser.ParsedResponses,
	routeId string,
) (
	components.ResponsesMap,
	error,
) {
	responses := components.ResponsesMap{}

	for k, v := range parsedRes {
		parentId := fmt.Sprintf(
			"%s.responses", routeId,
		)
		v.Context.ParentId = &parentId
		response, err := AssembleResponse(v)
		if err != nil {
			return nil, fmt.Errorf("Assembler.AssembleResponse: %+v", err)
		}

		responses[k] = response
	}

	return responses, nil
}

func AssembleResponse(parsedRes parser.ParsedResponse) (
	components.Response,
	error,
) {
	metadata, err := shared.AssembleComponentMetadata(parsedRes.Metadata, parsedRes.Context)
	if err != nil {
		return components.Response{}, fmt.Errorf("Assembler.AssembleComponentMetadata: %+v", err)
	}

	parsedRes.Body.Context.ParentId = &metadata.ComponentId
	body, err := body.AssembleMessageBody(*parsedRes.Body)
	if err != nil {
		return components.Response{}, fmt.Errorf("Assembler.AssembleMessageBody: %+v", err)
	}

	return components.Response{
		ComponentMetadata: metadata,
		StatusCode:        *parsedRes.StatusCode,
		Body:              &body,
	}, nil
}
