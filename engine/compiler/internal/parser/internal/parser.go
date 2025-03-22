package internal

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/object"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/route"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

type Parser struct{}

func DefaultParser() Parser {
	return Parser{}
}

func (p Parser) ParseProp(scannedComp scanner.ScannedComponent, ctx components.ComponentContext) (prop.ParsedProp, error) {
	return prop.ParseProp(scannedComp, ctx)
}

func (p Parser) ParseObject(scannedComp scanner.ScannedComponent, ctx components.ComponentContext) (object.ParsedObject, error) {
	return object.ParseObject(scannedComp, ctx)
}

func (p Parser) ParseRoute(scannedComp scanner.ScannedComponent, ctx components.ComponentContext) (route.ParsedRoute, error) {
	return route.ParseRoute(scannedComp, ctx)
}
