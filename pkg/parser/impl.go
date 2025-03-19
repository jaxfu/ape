package parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/jaxfu/ape/pkg/schemas"
)

type Parser struct{}

func NewParser() ParserInterface {
	return Parser{}
}

func (p Parser) ParseJSON(b []byte) (RawApeObject, error) {
	var os RawApeObject
	err := json.Unmarshal(b, &os)
	if err != nil {
		return os, err
	}
	return os, nil
}

func (p Parser) ParseTOML(b []byte) (RawApeObject, error) {
	var os RawApeObject
	err := toml.Unmarshal(b, &os)
	if err != nil {
		return os, err
	}
	return os, nil
}

func (p Parser) GenerateObject(s RawApeObject) (schemas.Object, error) {
	var obj schemas.Object

	if s.Name == nil {
		return obj, fmt.Errorf("name missing")
	}

	if s.Category == nil {
		obj.Category = "MAIN"
	} else {
		// TODO: cast to category id
		obj.Category = *s.Category
	}
	obj.Name = *s.Name
	obj.ObjectId = strings.ToLower(fmt.Sprintf("%s_%s", obj.Category, *s.Name))
	obj.Description = s.Description

	props, err := ParseProps(s.Props)
	if err != nil {
		return obj, err
	}
	obj.Props = props

	return obj, nil
}

func ParseProps(p RawApeObjectProps) ([]schemas.Prop, error) {
	typeMapping := map[string]string{
		"text": "TEXT",
		"int":  "INTEGER",
	}
	props := []schemas.Prop{}

	for name, p := range p {
		value, exists := p["type"]
		if !exists {
			return props, fmt.Errorf("no type found on prop '%s'", name)
		}
		str, ok := value.(string)
		if !ok {
			return props, fmt.Errorf("invalid type value: %+v on prop '%s'", value, name)
		}

		validType, exists := typeMapping[str]
		if !exists {
			return props, fmt.Errorf("invalid type value: %s on prop '%s'", value, name)
		}

		props = append(props, schemas.Prop{
			Name:         name,
			Type:         validType,
			IntFields:    &schemas.PropTypeIntegerFields{},
			StringFields: &schemas.PropTypeStringFields{},
		})
	}

	return props, nil
}
