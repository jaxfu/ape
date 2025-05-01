package unmarshal

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/pkg/openapi/shared"
)

// Schemas
// 1. Classify Schemas
// 2. Convert non-recursive props
// 3. Recursively convert array + object

func unmarshalSchemas(schemas openapi3.Schemas) (
	components.Components,
	error,
) {
	comps := components.Components{}

	for k, v := range schemas {
		if v == nil {
			return nil, fmt.Errorf("null schema")
		}
		if v.Value == nil {
			return nil, fmt.Errorf("null schema")
		}
		schema := *v.Value

		// get type
		schemaType, err := classifySchema(schema)
		if err != nil {
			return nil, fmt.Errorf("error classifying schema '%s': %+v", k, err)
		}

		// convert non-recursive props
		converted, err := schemaRouter[schemaType](k, schema, schemaType)
		if err != nil {
			return nil, fmt.Errorf("error converting schema '%s': %+v", k, err)
		}

		comps[k] = converted
	}

	return comps, nil
}

func classifySchema(schema openapi3.Schema) (shared.SchemaType, error) {
	if schema.Type == nil || len(*schema.Type) == 0 {
		return shared.SchemaTypes.Types().UNDEFINED, fmt.Errorf("schema '%s' missing type", schema.Title)
	}
	schemaTypeStr := strings.ToLower((*schema.Type)[0])
	schemaTypeStr = strings.TrimSpace(schemaTypeStr)
	schemaType := shared.SchemaTypes.Match(schemaTypeStr)

	return schemaType, nil
}

var schemaRouter = map[shared.SchemaType]func(string, openapi3.Schema, shared.SchemaType) (components.Component, error){
	shared.SchemaTypes.Types().INTEGER: convertProp,
	shared.SchemaTypes.Types().NUMBER:  convertProp,
	shared.SchemaTypes.Types().STRING:  convertProp,
	shared.SchemaTypes.Types().BOOLEAN: convertProp,
	// TODO: array + object schemas
	shared.SchemaTypes.Types().ARRAY: convertArray,
	// shared.SchemaTypes.Types().OBJECT: convertProp,
}
