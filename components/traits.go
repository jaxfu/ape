package components

const (
	TRAIT_TYPE_STRING    TraitType = "STRING"
	TRAIT_TYPE_INT       TraitType = "INT"
	TRAIT_TYPE_UINT      TraitType = "UINT"
	TRAIT_TYPE_FLOAT     TraitType = "FLOAT"
	TRAIT_TYPE_BOOL      TraitType = "BOOL"
	TRAIT_TYPE_UNDEFINED TraitType = "UNDEFINED"
)

// switch ttype {
// case TRAIT_TYPE_STRING:
// case TRAIT_TYPE_INT:
// case TRAIT_TYPE_UINT:
// case TRAIT_TYPE_FLOAT:
// case TRAIT_TYPE_BOOL:
// case TRAIT_TYPE_UNDEFINED:
// }

type TraitsMap map[string]Trait

type TraitType string

type TraitTypesConstraint interface {
	string |
		int |
		uint |
		float64 |
		bool
}

type TraitMetadata struct {
	Key  string
	Type TraitType
	Raw  string
}

type Trait interface {
	Meta() TraitMetadata
}

type TraitValue[T TraitTypesConstraint] struct {
	TraitMetadata
	Value T
}

type (
	TraitString = TraitValue[string]
	TraitInt    = TraitValue[int]
	TraitUint   = TraitValue[uint]
	TraitFloat  = TraitValue[float64]
	TraitBool   = TraitValue[bool]
)

// switch trait.(type) {
// case TraitString:
// case TraitInt:
// case TraitUint:
// case TraitFloat:
// case TraitBool:
// }

func (meta TraitMetadata) Meta() TraitMetadata {
	return meta
}
