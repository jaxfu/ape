package route

import (
	"fmt"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	compshared "github.com/jaxfu/ape/compiler/internal/shared"
	"maps"

	"github.com/jaxfu/ape/components"
)

func AssembleRequest(parsedReq parser.ParsedRequest) (compshared.CompiledRequest, error) {
	metadata, err := shared.AssembleComponentMetadata(
		parser.ParsedComponentMetadata{},
		parsedReq.Context,
	)
	if err != nil {
		return compshared.CompiledRequest{},
			fmt.Errorf(
				"Assembler.AssembleComponentMetadata: %+v",
				err,
			)
	}

	headers := components.HeadersMap{}
	if parsedReq.Headers != nil {
		maps.Copy(headers, *parsedReq.Headers)
	}

	var messageBody *compshared.CompiledBody = nil
	if parsedReq.Body != nil {
		parsedReq.Body.Context.ParentId = &metadata.ComponentId
		body, err := body.AssembleMessageBody(*parsedReq.Body)
		if err != nil {
			return compshared.CompiledRequest{}, fmt.Errorf("Assembler.AssembleMessageBody: %+v", err)
		}
		messageBody = &body
	}

	return compshared.CompiledRequest{
		ComponentMetadata: metadata,
		Headers:           headers,
		Body:              messageBody,
	}, nil
}
