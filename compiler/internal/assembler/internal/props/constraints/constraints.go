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

var propTypes = components.PropTypes.Types()

func AssembleConstraints(
	propType components.PropType,
	fields map[string]any,
) (
	components.PropConstraints,
	error,
) {
	switch propType {
	case propTypes.REF:
		return assembleRefConstraints(fields)
	case propTypes.INT:
		return assembleIntConstraints(fields)
	case propTypes.UINT:
		return assembleUintConstraints(fields)
	case propTypes.FLOAT:
		return assembleFloatConstraints(fields)
	case propTypes.TEXT:
		return assembleTextConstraints(fields)
	case propTypes.BLOB:
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
