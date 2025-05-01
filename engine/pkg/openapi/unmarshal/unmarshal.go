package unmarshal

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jaxfu/ape/components"
)

var (
	componentTypes = components.ComponentTypes.Types()
	propTypes      = components.PropTypes.Types()
)

func Unmarshal(fp string) (components.Components, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(fp)
	if err != nil {
		return nil, fmt.Errorf("loader.LoadFromFile: %+v", err)
	}

	err = doc.Components.Validate(loader.Context)
	if err != nil {
		return nil, fmt.Errorf("error validating OpenApi file: %+v", err)
	}

	// Schemas
	comps, err := unmarshalSchemas(doc.Components.Schemas)
	if err != nil {
		return nil, fmt.Errorf("convertSchemas: %+v", err)
	}

	return comps, nil
}
