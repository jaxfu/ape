package object

import (
	"fmt"

	prop2 "github.com/jaxfu/ape/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/scanner"

	"github.com/jaxfu/ape/components"
)

func ParseObject(scannedComp scanner.ScannedComponent, isRoot bool) (ParsedObject, error) {
	metadata, err := shared.ParseComponentMetadata(
		scannedComp.Fields,
		components.ComponentTypes.Types().OBJECT,
		isRoot,
	)
	if err != nil {
		return ParsedObject{}, fmt.Errorf("Parser.parseComponentMetadata: %+v", err)
	}

	parsedProps := prop2.ParsedProps{}
	rawScannedPropsMap, ok := scannedComp.Fields[shared.KEY_PROPS]
	if ok {
		scannedPropsMap, ok := rawScannedPropsMap.(map[string]any)
		if ok {
			parsedProps, err = prop2.ParseProps(scannedPropsMap)
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
		Context: shared.Context{
			ComponentType: components.ComponentTypes.Types().OBJECT,
			Name:          metadata.Name,
			IsRoot:        isRoot,
		},
	}, nil
}
