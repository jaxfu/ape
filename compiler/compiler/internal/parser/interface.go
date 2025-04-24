package parser

import (
	"github.com/jaxfu/ape/compiler/internal/parser/internal"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/body"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/object"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/prop"
	route2 "github.com/jaxfu/ape/compiler/internal/parser/internal/route"
	"github.com/jaxfu/ape/compiler/internal/parser/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/scanner"
)

type Parser interface {
	ParseProp(scanner.ScannedComponent, bool) (ParsedProp, error)
	ParseObject(scanner.ScannedComponent, bool) (ParsedObject, error)
	ParseRoute(scanner.ScannedComponent, bool) (ParsedRoute, error)
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
	ParsedRoute             = route2.ParsedRoute
	ParsedRouteMetadata     = route2.ParsedRouteMetadata
	ParsedRequest           = route2.ParsedRequest
	ParsedResponse          = route2.ParsedResponse
	ParsedResponses         = route2.ParsedResponsesMap
)
