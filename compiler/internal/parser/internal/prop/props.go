package prop

import (
	"fmt"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/scanner"

	"github.com/jaxfu/ape/components"
)

func ParseProps(scannedProps map[string]any) (ParsedProps, error) {
	parsedProps := ParsedProps{}

	for k, v := range scannedProps {
		if fields, ok := v.(map[string]any); ok {
			fields[shared.KEY_NAME] = k

			rawProp := scanner.ScannedComponent{
				ComponentType: components.ComponentTypes.PROP,
				Fields:        fields,
			}

			parsedProp, err := ParseProp(
				rawProp,
				false,
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
