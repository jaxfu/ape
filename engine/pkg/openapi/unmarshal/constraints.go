package unmarshal

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jaxfu/ape/components"
)

func convertIntegerConstraints(schema openapi3.Schema) components.PropConstraints {
	var max *int = nil
	if schema.Max != nil {
		i := int(*schema.Max)
		max = &i
	}

	var size *uint = nil
	format := strings.ToLower(schema.Format)
	format = strings.TrimSpace(format)
	switch format {
	case "int32":
		uintSize := uint(32)
		size = &uintSize
	case "int64":
		uintSize := uint(64)
		size = &uintSize
	}

	var min *int = nil
	if schema.Min != nil {
		i := int(*schema.Min)
		min = &i

		if *min >= 0 {
			uintMin := uint(*min)

			var uintMax *uint = nil
			if max != nil {
				u := uint(*max)
				uintMax = &u
			}
			return components.PropConstraintsUint{
				PropConstraintsMetadata: components.PropConstraintsMetadata{
					PropConstraintType: components.PropConstraintTypes.UINT,
				},
				Size: size,
				Min:  &uintMin,
				Max:  uintMax,
			}
		}
	}

	return components.PropConstraintsInt{
		PropConstraintsMetadata: components.PropConstraintsMetadata{
			PropConstraintType: components.PropConstraintTypes.INT,
		},
		Size: size,
		Min:  min,
		Max:  max,
	}
}

func convertFloatConstraints(schema openapi3.Schema) components.PropConstraintsFloat {
	var precision *string = nil
	format := strings.ToLower(schema.Format)
	format = strings.TrimSpace(format)
	float := "float"
	double := "double"
	switch format {
	case float:
		precision = &float
	case double:
		precision = &double
	}

	return components.PropConstraintsFloat{
		PropConstraintsMetadata: components.PropConstraintsMetadata{
			PropConstraintType: propTypes.FLOAT,
		},
		Precision: precision,
		Min:       schema.Min,
		Max:       schema.Max,
	}
}

func convertTextConstraints(schema openapi3.Schema) components.PropConstraints {
	format := strings.ToLower(schema.Format)
	format = strings.TrimSpace(format)
	if format == "binary" {
		return components.PropConstraintsBlob{
			PropConstraintsMetadata: components.PropConstraintsMetadata{
				PropConstraintType: components.PropConstraintTypes.BLOB,
			},
		}
	}

	var min *uint
	if schema.MinLength != 0 {
		minUint := uint(schema.MinLength)
		min = &minUint
	}

	var max *uint
	if schema.MinLength != 0 {
		maxUint := uint(*schema.MaxLength)
		max = &maxUint
	}

	return components.PropConstraintsText{
		PropConstraintsMetadata: components.PropConstraintsMetadata{
			PropConstraintType: components.PropConstraintTypes.TEXT,
		},
		MinLength: min,
		MaxLength: max,
	}
}
