package components

import (
	"github.com/jaxfu/ape/components/internal/props"
)

type (
	Prop         = props.Prop
	PropsMap     = props.PropsMap
	PropType     = props.PropType
	PropMetadata = props.PropMetadata
)

const (
	PROP_TYPE_INT   PropType = "INT"
	PROP_TYPE_UINT  PropType = "UINT"
	PROP_TYPE_FLOAT PropType = "FLOAT"
	PROP_TYPE_TEXT  PropType = "TEXT"
	PROP_TYPE_BOOL  PropType = "BOOL"
	PROP_TYPE_BLOB  PropType = "BLOB"
	PROP_TYPE_MAP   PropType = "MAP"
	PROP_TYPE_REF   PropType = "REF"
)

var ParsePropType = props.ParseType
