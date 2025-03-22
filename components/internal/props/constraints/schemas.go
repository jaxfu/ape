package constraints

import (
	"github.com/jaxfu/ape/components/internal/refs"
)

type PropConstraints = any

// Ref
type PropConstraintsRef struct {
	Reference refs.Reference `json:"target_id" toml:"target_id"`
}

// Int
type PropConstraintsInt struct {
	Size *uint `json:"size,omitempty"`
	Min  *int  `json:"min,omitempty"`
	Max  *int  `json:"max,omitempty"`
}

// Uint
type PropConstraintsUint struct {
	Size *uint `json:"size,omitempty"`
	Min  *uint `json:"min,omitempty"`
	Max  *uint `json:"max,omitempty"`
}

// Float
type PropConstraintsFloat struct {
	Precision *string  `json:"precision,omitempty"`
	Min       *float64 `json:"min,omitempty"`
	Max       *float64 `json:"max,omitempty"`
}

// Text
type PropConstraintsText struct {
	MinLength *uint   `json:"min_length,omitempty" mapstructure:"min_length,omitempty"`
	MaxLength *uint   `json:"max_length,omitempty" mapstructure:"max_length,omitempty"`
	Regex     *string `json:"regex,omitempty"`
	Alnum     *bool   `json:"alnum,omitempty"`
	Alpha     *bool   `json:"alpha,omitempty"`
	Num       *bool   `json:"num,omitempty"`
}

// Blob
type PropConstraintsBlob struct {
	MinSize *uint `json:"min_size,omitempty" mapstructure:"min_size,omitempty"`
	MaxSize *uint `json:"max_size,omitempty" mapstructure:"max_size,omitempty"`
}
