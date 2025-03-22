package body

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
)

func ParseMessageBody(rawBodyMap map[string]any, ctx components.ComponentContext) (*ParsedMessageBody, error) {
	metadata, err := shared.ParseComponentMetadata(rawBodyMap)
	if err != nil {
		return nil, fmt.Errorf("Paresr.ParseComponentMetadata: %+v", err)
	}

	ctx.IsRoot = false
	ctx.ComponentType = components.COMPONENT_TYPE_MESSAGE_BODY

	rawBody, ok := rawBodyMap[shared.KEY_MESSAGE_BODY]
	if !ok {
		return nil, nil
	}

	if str, ok := rawBody.(string); ok {
		return &ParsedMessageBody{
			Metadata: metadata,
			BodyType: components.MESSAGE_BODY_TYPE_REF,
			Ref:      &str,
			Context:  ctx,
		}, nil
	}

	props, ok := rawBody.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid body format: %+v", rawBody)
	}
	parsedProps, err := prop.ParseProps(props)
	if err != nil {
		return nil, fmt.Errorf("Parser.ParseProps: %+v", err)
	}

	return &ParsedMessageBody{
		BodyType: components.MESSAGE_BODY_TYPE_PROPS,
		Props:    &parsedProps,
		Context:  ctx,
	}, nil
}

type ParsedMessageBody struct {
	Metadata shared.ParsedComponentMetadata
	BodyType components.MessageBodyType
	Ref      *string
	Props    *prop.ParsedProps
	Context  components.ComponentContext
}
