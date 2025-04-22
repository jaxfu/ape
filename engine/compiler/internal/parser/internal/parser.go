package internal

import (
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/object"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/route"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

type Parser struct{}

func DefaultParser() Parser {
	return Parser{}
}

func (p Parser) ParseProp(scannedComp scanner.ScannedComponent, isRoot bool) (prop.ParsedProp, error) {
	return prop.ParseProp(scannedComp, isRoot)
}

func (p Parser) ParseObject(scannedComp scanner.ScannedComponent, isRoot bool) (object.ParsedObject, error) {
	return object.ParseObject(scannedComp, isRoot)
}

func (p Parser) ParseRoute(scannedComp scanner.ScannedComponent, isRoot bool) (route.ParsedRoute, error) {
	return route.ParseRoute(scannedComp, isRoot)
}
