package object

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props"
	asmshared "github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

func AssembleObject(parsedObj parser.ParsedObject) (components.Object, error) {
	metadata, err := asmshared.AssembleComponentMetadata(
		parsedObj.ComponentMetadata,
		parsedObj.Context,
	)
	if err != nil {
		return components.Object{}, fmt.Errorf("shared.AssembleComponentMetadata: %+v", err)
	}

	propsMap := components.PropsMap{}
	for k, v := range parsedObj.Props {
		v.Context.ParentId = &metadata.ComponentId

		prop, err := props.AssembleProp(v)
		if err != nil {
			return components.Object{},
				fmt.Errorf("Assemler.AssembleProp: %+v", err)
		}

		propsMap[k] = prop
	}

	return components.Object{
		ComponentMetadata: metadata,
		Props:             propsMap,
	}, nil
}
