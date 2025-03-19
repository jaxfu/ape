package sql

import "github.com/jaxfu/ape/pkg/schemas"

type SqlHandlerInterface interface {
	GenerateCreateTable(obj schemas.Object) (string, error)
}
