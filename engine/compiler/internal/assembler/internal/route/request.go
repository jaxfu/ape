package route

import (
	"fmt"
	"maps"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/body"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
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

	var messageBody *components.MessageBody = nil
	if parsedReq.Body != nil {
		parsedReq.Body.Context = components.ComponentContext{
			ComponentType: components.COMPONENT_TYPE_MESSAGE_BODY,
			IsRoot:        false,
			ParentId:      &metadata.ComponentId,
		}
		body, err := body.AssembleMessageBody(*parsedReq.Body)
		if err != nil {
			return components.Request{}, fmt.Errorf("Assembler.AssembleMessageBody: %+v", err)
		}
		messageBody = &body
	}

	return components.Request{
		ComponentMetadata: metadata,
		Headers:           headers,
		Body:              messageBody,
	}, nil
}
