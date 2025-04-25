package object

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler/internal/props"
	asmshared "github.com/jaxfu/ape/compiler/internal/assembler/internal/shared"
	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/components"
)

func AssembleObject(parsedObj parser.ParsedObject) (components.Object, error) {
	metadata, err := asmshared.AssembleComponentMetadata(
		parsedObj.ComponentMetadata,
		parsedObj.Context,
	)
	if err != nil {
		return components.Object{}, fmt.Errorf("Assembler.AssembleComponentMetadata: %+v", err)
	}

	propsMap, err := props.AssembleProps(parsedObj.Props, &metadata.ComponentId)
	if err != nil {
		return components.Object{}, fmt.Errorf("Assembler.AssembleProps: %+v", err)
	}

	return components.Object{
		ComponentMetadata: metadata,
		Props:             propsMap,
	}, nil
}
