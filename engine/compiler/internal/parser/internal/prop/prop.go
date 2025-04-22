package prop

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

func ParseProp(scannedComp scanner.ScannedComponent, isRoot bool) (ParsedProp, error) {
	metadata, err := shared.ParseComponentMetadata(scannedComp.Fields, components.COMPONENT_TYPE_PROP, isRoot)
	if err != nil {
		return ParsedProp{}, fmt.Errorf("Parser.parseComponentMetadata: %+v", err)
	}

	propTypeVal, _, err := shared.GetStringFromMap(scannedComp.Fields, shared.KEY_TYPE)
	if err != nil {
		return ParsedProp{}, fmt.Errorf("error finding %s: %+v", shared.KEY_TYPE, err)
	}
	propType, err := components.ParsePropType(propTypeVal)
	if err != nil {
		return ParsedProp{}, fmt.Errorf("parsePropType: %+v", err)
	}

	var isArr *bool = nil
	arrayVal, ok := scannedComp.Fields[shared.KEY_ARRAY]
	if ok {
		asStr, ok := arrayVal.(bool)
		if !ok {
			return ParsedProp{}, fmt.Errorf("invalid format for array, expected bool")
		}
		isArr = &asStr
	}

	delete(scannedComp.Fields, shared.KEY_NAME)
	delete(scannedComp.Fields, shared.KEY_CATEGORY)
	delete(scannedComp.Fields, shared.KEY_DESCRIPTION)
	delete(scannedComp.Fields, shared.KEY_TYPE)

	return ParsedProp{
		ComponentMetadata: metadata,
		PropMetadata: ParsedPropMetadata{
			PropType: propType,
			IsArray:  isArr,
		},
		Constraints: scannedComp.Fields,
		Context: shared.Context{
			ComponentType: components.COMPONENT_TYPE_PROP,
			Name:          metadata.Name,
			IsRoot:        isRoot,
		},
	}, nil
}
