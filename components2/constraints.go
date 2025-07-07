package components2

const (
	VALTYPE_STRING    ValueType = "VALTYPE_STRING"
	VALTYPE_INT       ValueType = "VALTYPE_INT"
	VALTYPE_UINT      ValueType = "VALTYPE_UINT"
	VALTYPE_FLOAT     ValueType = "VALTYPE_FLOAT"
	VALTYPE_BOOL      ValueType = "VALTYPE_BOOL"
	VALTYPE_UNDEFINED ValueType = "VALTYPE_UNDEFINED"
)

type Constraint struct {
	Key   string
	Value ConstraintValue
}

type ConstraintValue interface {
	Type() ValueType
}

type ValueType string
