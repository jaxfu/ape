package route

import (
	"fmt"
	"maps"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"

	"github.com/jaxfu/ape/components"
)

func AssembleRequest(parsedReq parser.ParsedRequest) (components.Request, error) {
	metadata, err := shared.AssembleComponentMetadata(
		parser.ParsedComponentMetadata{},
		parsedReq.Context,
	)
	if err != nil {
		return components.Request{},
			fmt.Errorf(
				"Assembler.AssembleComponentMetadata: %+v",
				err,
			)
	}

	headers := components.HeadersMap{}
	if parsedReq.Headers != nil {
		maps.Copy(headers, *parsedReq.Headers)
	}

	messageBody := components.MessageBody{}
	if parsedReq.Body != nil {
		parsedReq.Body.Context.ParentId = &metadata.ComponentId
		messageBody, err = body.AssembleMessageBody(*parsedReq.Body)
		if err != nil {
			return components.Request{}, fmt.Errorf("Assembler.AssembleMessageBody: %+v", err)
		}
	}

	return components.Request{
		ComponentMetadata: metadata,
		Headers:           headers,
		Body:              messageBody,
	}, nil
}
