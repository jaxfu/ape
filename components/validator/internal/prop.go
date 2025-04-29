package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (v Validator) validateProp(prop components.Prop) error {
	parsedType, err := components.ParsePropType(prop.PropMetadata.PropType)
	if err != nil {
		return fmt.Errorf(
			"error validating prop type: %+v",
			err,
		)
	}

	if err := validatePropConstraints(prop, parsedType); err != nil {
		return fmt.Errorf("error validating prop constraints on %s: %+v", prop.Name, err)
	}

	return nil
}

func (v Validator) validateProps(propsMap components.PropsMap) error {
	for k, p := range propsMap {
		if k != p.Name {
			return fmt.Errorf(
				"name mismatch for prop %s: got %s, want %s",
				p.ComponentId,
				p.Name,
				k,
			)
		}

		if err := v.validateProp(p); err != nil {
			return fmt.Errorf(
				"error validating prop %s: %+v",
				p.ComponentId,
				err,
			)
		}
	}

	return nil
}

func validatePropConstraints(prop components.Prop, parsedType components.PropType) error {
	switch prop.Constraints.(type) {

	case components.PropConstraintsRef:
		if parsedType != components.PropConstraintTypes.REF {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.REF,
				parsedType,
			)
		}
	case components.PropConstraintsInt:
		if parsedType != components.PropConstraintTypes.INT {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.INT,
				parsedType,
			)
		}
	case components.PropConstraintsUint:
		if parsedType != components.PropConstraintTypes.UINT {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.UINT,
				parsedType,
			)
		}
	case components.PropConstraintsFloat:
		if parsedType != components.PropConstraintTypes.FLOAT {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.FLOAT,
				parsedType,
			)
		}
	case components.PropConstraintsText:
		if parsedType != components.PropConstraintTypes.TEXT {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.TEXT,
				parsedType,
			)
		}
	case components.PropConstraintsBlob:
		if parsedType != components.PropConstraintTypes.BLOB {
			return fmt.Errorf(
				"incorrect prop type for %s: want %s, got %s",
				prop.ComponentId,
				components.PropConstraintTypes.BLOB,
				parsedType,
			)
		}
	}

	return nil
}
