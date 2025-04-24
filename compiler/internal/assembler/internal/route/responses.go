package route

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	compshared "github.com/jaxfu/ape/compiler/internal/shared"
)

// TODO: assemble responses
func AssembleResponses(parsedRes parser.ParsedResponses, routeId string) (compshared.CompiledResponses, error) {
	responses := compshared.CompiledResponses{}

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

func AssembleResponse(parsedRes parser.ParsedResponse) (compshared.CompiledResponse, error) {
	metadata, err := shared.AssembleComponentMetadata(parsedRes.Metadata, parsedRes.Context)
	if err != nil {
		return compshared.CompiledResponse{}, fmt.Errorf("Assembler.AssembleComponentMetadata: %+v", err)
	}

	parsedRes.Body.Context.ParentId = &metadata.ComponentId
	body, err := body.AssembleMessageBody(*parsedRes.Body)
	if err != nil {
		return compshared.CompiledResponse{}, fmt.Errorf("Assembler.AssembleMessageBody: %+v", err)
	}

	return compshared.CompiledResponse{
		CompiledComponentMetadata: metadata,
		StatusCode:                *parsedRes.StatusCode,
		Body:                      &body,
	}, nil
}
