package internal

import (
	object2 "github.com/jaxfu/ape/compiler/internal/parser/internal/object"
	prop2 "github.com/jaxfu/ape/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/route"
	"github.com/jaxfu/ape/compiler/internal/scanner"
)

type Parser struct{}

func DefaultParser() Parser {
	return Parser{}
}

func (p Parser) ParseProp(scannedComp scanner.ScannedComponent, isRoot bool) (prop2.ParsedProp, error) {
	return prop2.ParseProp(scannedComp, isRoot)
}

func (p Parser) ParseObject(scannedComp scanner.ScannedComponent, isRoot bool) (object2.ParsedObject, error) {
	return object2.ParseObject(scannedComp, isRoot)
}

func (p Parser) ParseRoute(scannedComp scanner.ScannedComponent, isRoot bool) (route.ParsedRoute, error) {
	return route.ParseRoute(scannedComp, isRoot)
}
