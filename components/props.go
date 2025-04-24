package components

import (
	"fmt"
	"strings"
)

type Prop struct {
	ComponentMetadata
	PropMetadata PropMetadata
	Constraints  PropConstraints
}
type PropMetadata struct {
	PropType PropType
	IsArray  bool
}

type PropsMap map[string]Prop

func ParsePropType(s string) (PropType, error) {
	t, ok := propTypesKeysMap[strings.ToLower(s)]
	if !ok || len(t) < 1 {
		return "", fmt.Errorf("invalid prop type value '%s'", s)
	}

	return t, nil
}

type (
	PropType           = string
	PropTypesInterface struct {
		UNDEFINED PropType `json:"undefined"`
		INT       PropType `json:"int"`
		UINT      PropType `json:"uint"`
		FLOAT     PropType `json:"float"`
		TEXT      PropType `json:"text"`
		BOOL      PropType `json:"bool"`
		BLOB      PropType `json:"blob"`
		MAP       PropType `json:"map"`
		REF       PropType `json:"ref"`
	}
)

var PropTypes = PropTypesInterface{
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

var propTypesKeysMap = map[string]PropType{
	"undefined": PropTypes.UNDEFINED,
	"int":       PropTypes.INT,
	"uint":      PropTypes.UINT,
	"float":     PropTypes.FLOAT,
	"text":      PropTypes.TEXT,
	"bool":      PropTypes.BOOL,
	"blob":      PropTypes.BLOB,
	"map":       PropTypes.MAP,
	"ref":       PropTypes.REF,
}
