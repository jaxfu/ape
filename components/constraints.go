package components

const (
	CONSTRAINT_VALUE_TYPE_STRING    ConstraintValueType = "STRING"
	CONSTRAINT_VALUE_TYPE_INT       ConstraintValueType = "INT"
	CONSTRAINT_VALUE_TYPE_UINT      ConstraintValueType = "UINT"
	CONSTRAINT_VALUE_TYPE_FLOAT     ConstraintValueType = "FLOAT"
	CONSTRAINT_VALUE_TYPE_BOOL      ConstraintValueType = "BOOL"
	CONSTRAINT_VALUE_TYPE_UNDEFINED ConstraintValueType = "UNDEFINED"
)

// bag
type Constraint struct {
	Key       string
	RawValue  string
	ValType   ConstraintValueType
	ValString string
	ValInt    int
	ValUint   uint
	ValFloat  float64
	ValBool   bool
}

type ConstraintValueType string
