package props

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props/constraints"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
)

func AssembleProp(parsedProp parser.ParsedProp) (components.Prop, error) {
	if parsedProp.PropMetadata.PropType == "" {
		return components.Prop{}, fmt.Errorf(
			"no type given for prop %s",
			parsedProp.ComponentMetadata.Name,
		)
	}

	// TODO: pass isSolo down
	metadata, err := shared.AssembleComponentMetadata(
		parsedProp.ComponentMetadata,
		parsedProp.Context,
	)
	if err != nil {
		return components.Prop{}, fmt.Errorf("Assembler.AssembleComponentMetadata: %+v", err)
	}

	constraints, err := constraints.AssembleConstraints(
		parsedProp.PropMetadata.PropType,
		parsedProp.Constraints,
	)
	if err != nil {
		return components.Prop{}, fmt.Errorf("Assembler.AssembleOpts: %+v", err)
	}

	isArr := parsedProp.PropMetadata.IsArray != nil && *parsedProp.PropMetadata.IsArray

	return components.Prop{
		ComponentMetadata: metadata,
		PropMetadata: components.PropMetadata{
			PropType: parsedProp.PropMetadata.PropType,
			IsArray:  isArr,
		},
		Constraints: constraints,
	}, nil
}
