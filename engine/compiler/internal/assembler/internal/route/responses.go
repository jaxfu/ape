package route

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

// TODO: assemble responses
func AssembleResponses(parsedReq parser.ParsedResponses, routeId components.ComponentId) (components.ResponsesMap, error) {
	responses := components.ResponsesMap{}

	for k, v := range parsedReq {
		parentId := components.ComponentId{Display: fmt.Sprintf("%s.responses", routeId.Display)}
		v.Context.ParentId = &parentId
		response, err := AssembleResponse(v)
		if err != nil {
			return nil, fmt.Errorf("Assembler.AssembleResponse: %+v", err)
		}

		responses[k] = response
	}

	return responses, nil
}

func AssembleResponse(parsedRes parser.ParsedResponse) (components.Response, error) {
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
		Name:              *parsedRes.Context.Name,
		StatusCode:        *parsedRes.StatusCode,
		Body:              body,
	}, nil
}
