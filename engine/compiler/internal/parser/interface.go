package parser

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/body"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/object"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/prop"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/route"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
)

type Parser interface {
	ParseProp(scanner.ScannedComponent, components.ComponentContext) (ParsedProp, error)
	ParseObject(scanner.ScannedComponent, components.ComponentContext) (ParsedObject, error)
	ParseRoute(scanner.ScannedComponent, components.ComponentContext) (ParsedRoute, error)
}

func NewParser() Parser {
	return internal.DefaultParser()
}

type (
	ParsedComponentMetadata = shared.ParsedComponentMetadata
	ParsedMessageBody       = body.ParsedMessageBody
	ParsedProp              = prop.ParsedProp
	ParsedProps             = prop.ParsedProps
	ParsedPropMetadata      = prop.ParsedPropMetadata
	ParsedObject            = object.ParsedObject
	ParsedRoute             = route.ParsedRoute
	ParsedRouteMetadata     = route.ParsedRouteMetadata
	ParsedRequest           = route.ParsedRequest
	ParsedResponse          = route.ParsedResponse
	ParsedResponses         = route.ParsedResponsesMap
)
