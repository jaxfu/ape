package body

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"

	"github.com/jaxfu/ape/components"
)

func ParseMessageBody(rawBodyMap map[string]any, isRoot bool) (*ParsedMessageBody, error) {
	metadata, err := shared.ParseComponentMetadata(rawBodyMap, components.COMPONENT_TYPE_MESSAGE_BODY, isRoot)
	if err != nil {
		return nil, fmt.Errorf("Parser.ParseComponentMetadata: %+v", err)
	}

	rawBody, ok := rawBodyMap[shared.KEY_MESSAGE_BODY]
	if !ok {
		return nil, nil
	}

	if str, ok := rawBody.(string); ok {
		return &ParsedMessageBody{
			Metadata: metadata,
			BodyType: components.MESSAGE_BODY_TYPE_REF,
			Ref:      &str,
			Context: shared.Context{
				ComponentType: components.COMPONENT_TYPE_MESSAGE_BODY,
				Name:          metadata.Name,
				IsRoot:        isRoot,
			},
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
		Metadata: metadata,
		BodyType: components.MESSAGE_BODY_TYPE_PROPS,
		Props:    &parsedProps,
		Context: shared.Context{
			ComponentType: components.COMPONENT_TYPE_MESSAGE_BODY,
			Name:          metadata.Name,
			IsRoot:        isRoot,
		},
	}, nil
}

type ParsedMessageBody struct {
	Metadata shared.ParsedComponentMetadata
	BodyType components.MessageBodyType
	Ref      *string
	Props    *prop.ParsedProps
	Context  shared.Context
}
