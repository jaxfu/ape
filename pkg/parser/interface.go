package parser

import "github.com/jaxfu/ape/pkg/schemas"

type ParserInterface interface {
	ParseJSON([]byte) (RawApeObject, error)
	ParseTOML([]byte) (RawApeObject, error)
	GenerateObject(RawApeObject) (schemas.Object, error)
}
