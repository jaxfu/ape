package shared

import (
	"github.com/jaxfu/ape/pkg/enum"
)

var SchemaTypes = enum.Enum[SchemaType, SchemaTypesInterface]{
	TypeList: Types,
	MatchMap: TypesMatchMap,
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

// type EnumSchemaTypes struct {
// 	TypesInterface SchemaTypesInterface
// 	MatchMap       map[string]SchemaType
// }
//
// func (s EnumSchemaTypes) Types() SchemaTypesInterface {
// 	return s.TypesInterface
// }
//
// func (s EnumSchemaTypes) Match(src string) SchemaType {
// 	src = strings.ToLower(src)
// 	src = strings.TrimSpace(src)
//
// 	found, ok := s.MatchMap[src]
// 	if !ok {
// 		return s.Types().UNDEFINED
// 	}
//
// 	return found
// }
