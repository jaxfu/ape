package unmarshal

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/pkg/dev"
	"github.com/jaxfu/ape/engine/pkg/openapi/shared"
)

func convertArray(
	name string,
	schema openapi3.Schema,
	schemaType shared.SchemaType,
) (components.Component, error) {
	curr := schema
	for curr.Items != nil {
		if curr.Items.Value != nil {
			dev.PrettyPrint(curr.Items.Value)
			curr = *curr.Items.Value
		}
	}

	return components.Prop{}, nil
}
