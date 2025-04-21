package shared

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
	"github.com/jaxfu/ape/engine/pkg/idhandler"
)

func AssembleComponentMetadata(
	parsedMetadata parser.ParsedComponentMetadata,
	ctx components.ComponentContext,
) (shared.CompiledComponentMetadata, error) {
	componentId, err := idhandler.NewIdHandler().Generate(
		idhandler.GenerateIdParams{
			ComponentContext: ctx,
			Category:         parsedMetadata.Category,
		})
	if err != nil {
		return shared.CompiledComponentMetadata{},
			fmt.Errorf("IdHandler.Generate: %+v", err)
	}

	var description *string = nil
	if parsedMetadata.Description != nil {
		if *parsedMetadata.Description != "" {
			description = parsedMetadata.Description
		}
	}

	name := ""
	if ctx.Name == nil {
		name = componentId.Display
	} else {
		name = *ctx.Name
	}
	metadata := shared.CompiledComponentMetadata{
		ComponentType: ctx.ComponentType,
		Name:          name,
		ComponentId:   componentId.Display,
		IsRoot:        ctx.IsRoot,
		Category:      parsedMetadata.Category,
		Description:   description,
	}

	return metadata, nil
}
