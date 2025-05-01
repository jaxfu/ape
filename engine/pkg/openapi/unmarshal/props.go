package unmarshal

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/pkg/openapi/shared"
)

func convertProp(
	name string,
	schema openapi3.Schema,
	schemaType shared.SchemaType,
) (components.Component, error) {
	prop := components.Prop{
		ComponentMetadata: components.ComponentMetadata{
			ComponentType: componentTypes.PROP,
			Name:          name,
		},
	}

	schemaTypes := shared.SchemaTypes.Types()
	switch schemaType {
	case schemaTypes.INTEGER:
		constraints := convertIntegerConstraints(schema)
		if constraints.ConstraintType() == components.PropConstraintTypes.UINT {
			prop.PropMetadata.PropType = propTypes.UINT
		} else {
			prop.PropMetadata.PropType = propTypes.INT
		}
		prop.Constraints = constraints

	case schemaTypes.NUMBER:
		constraints := convertFloatConstraints(schema)
		prop.PropMetadata.PropType = propTypes.FLOAT
		prop.Constraints = constraints

	case schemaTypes.STRING:
		constraints := convertTextConstraints(schema)
		if constraints.ConstraintType() == components.PropConstraintTypes.BLOB {
			prop.PropMetadata.PropType = propTypes.BLOB
		} else {
			prop.PropMetadata.PropType = propTypes.TEXT
		}
		prop.Constraints = constraints

	case schemaTypes.BOOLEAN:
		constraints := components.PropConstraintsBool{
			PropConstraintsMetadata: components.PropConstraintsMetadata{
				PropConstraintType: components.PropConstraintTypes.BOOL,
			},
		}
		prop.Constraints = constraints
		prop.PropMetadata.PropType = propTypes.BOOL
	}

	return prop, nil
}
