package constraints

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

const (
	SIZE string = "size"
	MIN  string = "min"
	MAX  string = "max"
)

func AssembleConstraints(propType components.PropType, fields map[string]any) (components.PropConstraints, error) {
	switch propType {
	case components.PROP_TYPE_REF:
		return assembleRefConstraints(fields)
	case components.PROP_TYPE_INT:
		return assembleIntConstraints(fields)
	case components.PROP_TYPE_UINT:
		return assembleUintConstraints(fields)
	case components.PROP_TYPE_FLOAT:
		return assembleFloatConstraints(fields)
	case components.PROP_TYPE_TEXT:
		return assembleTextConstraints(fields)
	case components.PROP_TYPE_BLOB:
		return assembleBlobConstraints(fields)
	}

	return nil, fmt.Errorf("unrecognized prop type %s", propType)
}

func extractFromMap[T any](key string, fields map[string]any) (T, bool, error) {
	var value T

	raw, ok := fields[key]
	if !ok {
		return value, false, nil
	}

	asType, ok := raw.(T)
	if !ok {
		return value, true, fmt.Errorf("incrorrect format for %s", key)
	}

	return asType, true, nil
}
