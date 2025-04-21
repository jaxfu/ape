package body

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	compShared "github.com/jaxfu/ape/engine/compiler/internal/shared"
)

func AssembleMessageBody(body parser.ParsedMessageBody) (compShared.CompiledBody, error) {
	metadata, err := shared.AssembleComponentMetadata(body.Metadata, body.Context)
	if err != nil {
		return compShared.CompiledBody{}, fmt.Errorf("Assember.AssembleComponentMetadata: %+v", err)
	}

	switch body.BodyType {
	case components.MESSAGE_BODY_TYPE_REF:
		if body.Ref == nil {
			return compShared.CompiledBody{}, fmt.Errorf("no content on message body")
		}

		return compShared.CompiledBody{
			BodyType: components.MESSAGE_BODY_TYPE_REF,
			Ref:      *body.Ref,
		}, nil
	case components.MESSAGE_BODY_TYPE_PROPS:
		if body.Props == nil {
			return compShared.CompiledBody{},
				fmt.Errorf("no content on message body")
		}

		props, err := assembleMessageBodyProps(*body.Props, metadata.ComponentId)
		if err != nil {
			return compShared.CompiledBody{},
				fmt.Errorf(
					"Assembler.assembleMessageBodyProps: %+v",
					err,
				)
		}

		return compShared.CompiledBody{
			ComponentMetadata: metadata,
			BodyType:          components.MESSAGE_BODY_TYPE_PROPS,
			Props:             map[string]compShared.CompiledProp{},
		}, nil
	}

	return compShared.CompiledBody{}, fmt.Errorf("invalid message body type %s", body.BodyType)
}

func assembleMessageBodyProps(parsedProps parser.ParsedProps, parentId string) (compShared.CompiledProps, error) {
	propsMap := compShared.CompiledProps{}

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
