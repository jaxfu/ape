package body

import (
	"fmt"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/props"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	compshared "github.com/jaxfu/ape/compiler/internal/shared"

	"github.com/jaxfu/ape/components"
)

func AssembleMessageBody(body parser.ParsedMessageBody) (compshared.CompiledBody, error) {
	metadata, err := shared.AssembleComponentMetadata(body.Metadata, body.Context)
	if err != nil {
		return compshared.CompiledBody{}, fmt.Errorf("Assember.AssembleComponentMetadata: %+v", err)
	}

	switch body.BodyType {
	case components.MESSAGE_BODY_TYPE_REF:
		if body.Ref == nil {
			return compshared.CompiledBody{}, fmt.Errorf("no content on message body")
		}

		return compshared.CompiledBody{
			BodyType: components.MESSAGE_BODY_TYPE_REF,
			Ref:      *body.Ref,
		}, nil
	case components.MESSAGE_BODY_TYPE_PROPS:
		if body.Props == nil {
			return compshared.CompiledBody{},
				fmt.Errorf("no content on message body")
		}

		props, err := assembleMessageBodyProps(*body.Props, metadata.ComponentId)
		if err != nil {
			return compshared.CompiledBody{},
				fmt.Errorf(
					"Assembler.assembleMessageBodyProps: %+v",
					err,
				)
		}

		return compshared.CompiledBody{
			ComponentMetadata: metadata,
			BodyType:          components.MESSAGE_BODY_TYPE_PROPS,
			Props:             props,
		}, nil
	}

	return compshared.CompiledBody{}, fmt.Errorf("invalid message body type %s", body.BodyType)
}

func assembleMessageBodyProps(parsedProps parser.ParsedProps, parentId string) (map[string]compshared.CompiledProp, error) {
	propsMap := map[string]compshared.CompiledProp{}

	for k, v := range parsedProps {
		v.Context.ParentId = &parentId
		prop, err := props.AssembleProp(v)
		if err != nil {
			return nil,
				fmt.Errorf("Assemler.AssembleProp: %+v", err)
		}

		propsMap[k] = prop
	}

	return propsMap, nil
}
