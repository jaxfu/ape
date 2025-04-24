package constraints

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

const (
	MIN_LENGTH string = "min_length"
	MAX_LENGTH string = "max_length"
	REGEX      string = "regex"
	ALPHA      string = "alpha"
	ALNUM      string = "alnum"
	NUM        string = "num"
)

func assembleTextConstraints(fields map[string]any) (components.PropConstraints, error) {
	constraints := components.PropConstraintsText{}

	min, exists, err := extractFromMap[int64](MIN_LENGTH, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MIN_LENGTH, err)
	} else if exists {
		asUint := uint(min)
		constraints.MinLength = &asUint
	}

	max, exists, err := extractFromMap[int64](MAX_LENGTH, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", MAX_LENGTH, err)
	} else if exists {
		asUint := uint(max)
		constraints.MaxLength = &asUint
	}

	regex, exists, err := extractFromMap[string](REGEX, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", REGEX, err)
	} else if exists {
		constraints.Regex = &regex
	}

	alpha, exists, err := extractFromMap[bool](ALPHA, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", ALPHA, err)
	} else if exists {
		constraints.Alpha = &alpha
	}

	num, exists, err := extractFromMap[bool](NUM, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", NUM, err)
	} else if exists {
		constraints.Num = &num
	}

	alnum, exists, err := extractFromMap[bool](ALNUM, fields)
	if err != nil {
		return nil, fmt.Errorf("error finding %s: %+v", ALNUM, err)
	} else if exists {
		constraints.Alnum = &alnum
	}

	return constraints, nil
}
