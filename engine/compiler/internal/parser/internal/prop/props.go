package prop

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

func ParseProps(scannedProps map[string]any) (ParsedProps, error) {
	parsedProps := ParsedProps{}

	for k, v := range scannedProps {
		if fields, ok := v.(map[string]any); ok {
			fields[shared.KEY_NAME] = k

			rawProp := scanner.ScannedComponent{
				ComponentType: components.COMPONENT_TYPE_PROP,
				Fields:        fields,
			}

			parsedProp, err := ParseProp(
				rawProp,
				components.ComponentContext{
					IsRoot:        false,
					ComponentType: components.COMPONENT_TYPE_PROP,
					Name:          &k,
				},
			)
			if err != nil {
				return nil, fmt.Errorf("Parser.ParseProp: %+v", err)
			}
			parsedProps[k] = parsedProp
		} else {
			return nil, fmt.Errorf("invalid type for prop %s: %+v", k, fields)
		}
	}

	return parsedProps, nil
}
