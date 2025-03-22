package object

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

func ParseObject(scannedComp scanner.ScannedComponent, ctx components.ComponentContext) (ParsedObject, error) {
	metadata, err := shared.ParseComponentMetadata(scannedComp.Fields)
	if err != nil {
		return ParsedObject{}, fmt.Errorf("Parser.parseComponentMetadata: %+v", err)
	}

	parsedProps := prop.ParsedProps{}
	rawScannedPropsMap, ok := scannedComp.Fields[shared.KEY_PROPS]
	if ok {
		scannedPropsMap, ok := rawScannedPropsMap.(map[string]any)
		if ok {
			parsedProps, err = prop.ParseProps(scannedPropsMap)
			if err != nil {
				return ParsedObject{}, fmt.Errorf("Parser.ParseProps: %+v", err)
			}
		} else {
			return ParsedObject{}, fmt.Errorf("invalid type for Props: %+v", scannedPropsMap)
		}
	}

	return ParsedObject{
		ComponentMetadata: metadata,
		Props:             parsedProps,
		Context: components.ComponentContext{
			ComponentType: ctx.ComponentType,
			IsRoot:        ctx.IsRoot,
			Name:          metadata.Name,
		},
	}, nil
}
