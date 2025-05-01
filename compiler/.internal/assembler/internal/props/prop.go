package props

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/props/constraints"
	"github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"

	"github.com/jaxfu/ape/components"
)

func AssembleProp(parsedProp parser.ParsedProp) (components.Prop, error) {
	if parsedProp.PropMetadata.PropType == "" {
		return components.Prop{}, fmt.Errorf(
			"no prop type given",
		)
	}

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

func AssembleProps(props parser.ParsedProps, parentId *string) (components.PropsMap, error) {
	propsMap := components.PropsMap{}
	for k, v := range props {
		v.Context.ParentId = parentId

		prop, err := AssembleProp(v)
		if err != nil {
			return nil,
				fmt.Errorf("Assembler.AssembleProp: %+v", err)
		}

		propsMap[k] = prop
	}

	return propsMap, nil
}
