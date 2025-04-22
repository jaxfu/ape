package props

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/props/constraints"
	"github.com/jaxfu/ape/engine/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	compShared "github.com/jaxfu/ape/engine/compiler/internal/shared"
)

func AssembleProp(parsedProp parser.ParsedProp) (compShared.CompiledProp, error) {
	if parsedProp.PropMetadata.PropType == "" {
		return compShared.CompiledProp{}, fmt.Errorf(
			"no prop type given",
		)
	}

	metadata, err := shared.AssembleComponentMetadata(
		parsedProp.ComponentMetadata,
		parsedProp.Context,
	)
	if err != nil {
		return compShared.CompiledProp{}, fmt.Errorf("Assembler.AssembleComponentMetadata: %+v", err)
	}

	constraints, err := constraints.AssembleConstraints(
		parsedProp.PropMetadata.PropType,
		parsedProp.Constraints,
	)
	if err != nil {
		return compShared.CompiledProp{}, fmt.Errorf("Assembler.AssembleOpts: %+v", err)
	}

	isArr := parsedProp.PropMetadata.IsArray != nil && *parsedProp.PropMetadata.IsArray

	return compShared.CompiledProp{
		ComponentMetadata: metadata,
		PropMetadata: components.PropMetadata{
			PropType: parsedProp.PropMetadata.PropType,
			IsArray:  isArr,
		},
		Constraints: constraints,
	}, nil
}
