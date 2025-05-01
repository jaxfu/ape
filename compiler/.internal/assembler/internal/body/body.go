package body

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/props"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"

	"github.com/jaxfu/ape/components"
)

func AssembleMessageBody(body parser.ParsedMessageBody) (components.MessageBody, error) {
	metadata, err := shared.AssembleComponentMetadata(body.Metadata, body.Context)
	if err != nil {
		return components.MessageBody{}, fmt.Errorf("Assember.AssembleComponentMetadata: %+v", err)
	}

	switch body.BodyType {
	case components.MESSAGE_BODY_TYPE_REF:
		if body.Ref == nil {
			return components.MessageBody{}, fmt.Errorf("no content on message body")
		}

		return components.MessageBody{
			BodyType: components.MESSAGE_BODY_TYPE_REF,
			Ref:      *body.Ref,
		}, nil
	case components.MESSAGE_BODY_TYPE_PROPS:
		if body.Props == nil {
			return components.MessageBody{},
				fmt.Errorf("no content on message body")
		}

		propsMap, err := props.AssembleProps(*body.Props, &metadata.ComponentId)
		if err != nil {
			return components.MessageBody{},
				fmt.Errorf(
					"Assembler.assembleMessageBodyProps: %+v",
					err,
				)
		}

		return components.MessageBody{
			ComponentMetadata: metadata,
			BodyType:          components.MESSAGE_BODY_TYPE_PROPS,
			Props:             propsMap,
		}, nil
	}

	return components.MessageBody{}, fmt.Errorf("invalid message body type %s", body.BodyType)
}
