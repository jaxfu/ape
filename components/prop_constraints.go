package components

type PropConstraints interface {
	ConstraintType() PropConstraintType
}

type PropConstraintsMetadata struct {
	PropConstraintType PropConstraintType
}

func (meta PropConstraintsMetadata) ConstraintType() PropConstraintType {
	return meta.PropConstraintType
}

func PropConstraintSwitch(constraint PropConstraints) {
	switch constraint.(type) {
	case PropConstraintsRef:
	case PropConstraintsInt:
	case PropConstraintsUint:
	case PropConstraintsFloat:
	case PropConstraintsText:
	case PropConstraintsBlob:
	}
}

type (
	PropConstraintType           = string
	PropConstraintTypesInterface struct {
		UNDEFINED PropConstraintType `json:"undefined"`
		INT       PropConstraintType `json:"int"`
		UINT      PropConstraintType `json:"uint"`
		FLOAT     PropConstraintType `json:"float"`
		TEXT      PropConstraintType `json:"text"`
		BOOL      PropConstraintType `json:"bool"`
		BLOB      PropConstraintType `json:"blob"`
		MAP       PropConstraintType `json:"map"`
		REF       PropConstraintType `json:"ref"`
	}
)

var PropConstraintTypes = PropConstraintTypesInterface{
	UNDEFINED: "UNDEFINED",
	INT:       "INT",
	UINT:      "UINT",
	FLOAT:     "FLOAT",
	TEXT:      "TEXT",
	BOOL:      "BOOL",
	BLOB:      "BLOB",
	MAP:       "MAP",
	REF:       "REF",
}

var propConstraintTypesKeysMap = map[string]PropConstraintType{
	"undefined": PropConstraintTypes.UNDEFINED,
	"int":       PropConstraintTypes.INT,
	"uint":      PropConstraintTypes.UINT,
	"float":     PropConstraintTypes.FLOAT,
	"text":      PropConstraintTypes.TEXT,
	"bool":      PropConstraintTypes.BOOL,
	"blob":      PropConstraintTypes.BLOB,
	"map":       PropConstraintTypes.MAP,
	"ref":       PropConstraintTypes.REF,
}

// Ref
type PropConstraintsRef struct {
	PropConstraintsMetadata
	Reference ReferenceTag `json:"target_id" toml:"target_id"`
}

// Int
type PropConstraintsInt struct {
	PropConstraintsMetadata
	Size *uint `json:"size,omitempty"`
	Min  *int  `json:"min,omitempty"`
	Max  *int  `json:"max,omitempty"`
}

// Uint
type PropConstraintsUint struct {
	PropConstraintsMetadata
	Size *uint `json:"size,omitempty"`
	Min  *uint `json:"min,omitempty"`
	Max  *uint `json:"max,omitempty"`
}

// Float
type PropConstraintsFloat struct {
	PropConstraintsMetadata
	Precision *string  `json:"precision,omitempty"`
	Min       *float64 `json:"min,omitempty"`
	Max       *float64 `json:"max,omitempty"`
}

// Text
type PropConstraintsText struct {
	PropConstraintsMetadata
	MinLength *uint   `json:"min_length,omitempty"`
	MaxLength *uint   `json:"max_length,omitempty"`
	Regex     *string `json:"regex,omitempty"`
	Alnum     *bool   `json:"alnum,omitempty"`
	Alpha     *bool   `json:"alpha,omitempty"`
	Num       *bool   `json:"num,omitempty"`
}

// Blob
type PropConstraintsBlob struct {
	PropConstraintsMetadata
	MinSize *uint `json:"min_size,omitempty"`
	MaxSize *uint `json:"max_size,omitempty"`
}
