package sql

import (
	"fmt"

	"github.com/jaxfu/ape/pkg/schemas"
)

type SqlHandler struct{}

func NewSqlHandler() SqlHandler {
	return SqlHandler{}
}

func (s SqlHandler) GenerateCreateTable(obj schemas.Object) (string, error) {
	var cols string
	for i, row := range obj.Props {
		if i < len(obj.Props)-1 {
			cols += fmt.Sprintf("%s %s,\n", row.Name, row.Type)
		} else {
			cols += fmt.Sprintf("%s %s\n", row.Name, row.Type)
		}
	}

	outStr := fmt.Sprintf("CREATE TABLE %s\n(\n%s);\n", obj.Name, cols)

	return outStr, nil
}
