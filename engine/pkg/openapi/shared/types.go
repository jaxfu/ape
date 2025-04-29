package shared

import "strings"

type Enum[T ~string, S any] interface {
	Match(string) T
	Types() S
}

var SchemaTypes Enum[SchemaType, SchemaTypesInterface] = EnumSchemaTypes{
	TypesInterface: Types,
	MatchMap:       TypesMatchMap,
}

type EnumSchemaTypes struct {
	TypesInterface SchemaTypesInterface
	MatchMap       map[string]SchemaType
}

func (s EnumSchemaTypes) Types() SchemaTypesInterface {
	return s.TypesInterface
}

func (s EnumSchemaTypes) Match(src string) SchemaType {
	src = strings.ToLower(src)
	src = strings.TrimSpace(src)

	found, ok := s.MatchMap[src]
	if !ok {
		return s.Types().UNDEFINED
	}

	return found
}

type (
	SchemaType           = string
	SchemaTypesInterface struct {
		UNDEFINED SchemaType
		STRING    SchemaType
		NUMBER    SchemaType
		INTEGER   SchemaType
		BOOLEAN   SchemaType
		ARRAY     SchemaType
		OBJECT    SchemaType
	}
)

var Types = SchemaTypesInterface{
	STRING:  "string",
	NUMBER:  "number",
	INTEGER: "integer",
	BOOLEAN: "boolean",
	ARRAY:   "array",
	OBJECT:  "object",
}

var TypesMatchMap = map[string]SchemaType{
	"string":  Types.STRING,
	"number":  Types.NUMBER,
	"integer": Types.INTEGER,
	"boolean": Types.BOOLEAN,
	"array":   Types.ARRAY,
	"object":  Types.OBJECT,
}
