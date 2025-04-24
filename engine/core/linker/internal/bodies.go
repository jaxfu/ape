package internal

import (
	"fmt"
	"github.com/jaxfu/ape/compiler"

	"github.com/jaxfu/ape/components"
)

func (l *Linker) linkBodies(objects map[string]compiler.CompiledBody) (map[components.ComponentId]components.MessageBody, error) {
	linked := map[components.ComponentId]components.MessageBody{}

	if len(objects) > 0 {
		for k, v := range objects {
			body, err := l.linkBody(v)
			if err != nil {
				return nil, fmt.Errorf("Linker.linkBody: %+v", err)
			}
			linked[k] = body
		}
	}

	return linked, nil
}

func (l *Linker) linkBody(body compiler.CompiledBody) (components.MessageBody, error) {
	meta, err := l.LinkComponent(body.ComponentMetadata)
	if err != nil {
		return components.MessageBody{}, fmt.Errorf("Linker.LinkComponent: %+v", err)
	}

	// discriminator
	var ref components.ReferenceTag
	props := map[components.ComponentId]components.Prop{}
	switch body.BodyType {
	case components.MESSAGE_BODY_TYPE_REF:
		ref, err = l.linkRef(body.Ref)
		if err != nil {
			return components.MessageBody{}, fmt.Errorf("Linker.linkRef: %+v", err)
		}
	case components.MESSAGE_BODY_TYPE_PROPS:
		props, err = l.linkProps(body.Props)
		if err != nil {
			return components.MessageBody{}, fmt.Errorf("Linker.linkProps: %+v", err)
		}
	}

	return components.MessageBody{
		ComponentMetadata: meta,
		BodyType:          body.BodyType,
		Ref:               ref,
		Props:             props,
	}, nil
}
