package components

import (
	"github.com/jaxfu/ape/pkg/enum"
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

type (
	PropType           = string
	PropTypesInterface struct {
		UNDEFINED  PropType `json:"undefined"`
		INT        PropType `json:"int"`
		UINT       PropType `json:"uint"`
		FLOAT      PropType `json:"float"`
		TEXT       PropType `json:"text"`
		BOOL       PropType `json:"bool"`
		BLOB       PropType `json:"blob"`
		MAP        PropType `json:"map"`
		REF        PropType `json:"ref"`
		COMPONENT  PropType `json:"component"`
		COLLECTION PropType `json:"collection"`
	}
)

var PropTypes = enum.Enum[PropType, PropTypesInterface]{
	TypeList: PropTypesImpl,
	MatchMap: map[string]PropType{},
}

var PropTypesImpl = PropTypesInterface{
	UNDEFINED: "UNDEFINED",
	INT:       "INT",
	UINT:      "UINT",
	FLOAT:     "FLOAT",
	TEXT:      "TEXT",
	BOOL:      "BOOL",
	BLOB:      "BLOB",
	MAP:       "MAP",
	REF:       "REF",
	COMPONENT: "COMPONENT",
}
