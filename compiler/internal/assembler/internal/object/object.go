package object

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/props"
	asmshared "github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/compiler/internal/shared"
)

func AssembleObject(parsedObj parser.ParsedObject) (shared.CompiledObject, error) {
	metadata, err := asmshared.AssembleComponentMetadata(
		parsedObj.ComponentMetadata,
		parsedObj.Context,
	)
	if err != nil {
		return shared.CompiledObject{}, fmt.Errorf("shared.AssembleComponentMetadata: %+v", err)
	}

	propsMap := map[string]shared.CompiledProp{}
	for k, v := range parsedObj.Props {
		v.Context.ParentId = &metadata.ComponentId

		prop, err := props.AssembleProp(v)
		if err != nil {
			return shared.CompiledObject{},
				fmt.Errorf("Assemler.AssembleProp: %+v", err)
		}

		propsMap[k] = prop
	}

	return shared.CompiledObject{
		CompiledComponentMetadata: metadata,
		Props:                     propsMap,
	}, nil
}
