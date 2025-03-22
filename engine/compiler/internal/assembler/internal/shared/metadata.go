package shared

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/pkg/idhandler"
)

func AssembleComponentMetadata(
	parsedMetadata parser.ParsedComponentMetadata,
	ctx components.ComponentContext,
) (components.ComponentMetadata, error) {
	componentId, err := idhandler.NewIdHandler().Generate(
		idhandler.GenerateIdParams{
			ComponentContext: ctx,
			Category:         parsedMetadata.Category,
		})
	if err != nil {
		return components.ComponentMetadata{},
			fmt.Errorf("IdHandler.Generate: %+v", err)
	}

	var category *components.CategoryId = nil
	if parsedMetadata.Category != nil {
		if *parsedMetadata.Category != "" {
			category = &components.CategoryId{
				Display: *parsedMetadata.Category,
			}
		}
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
	metadata := components.ComponentMetadata{
		ComponentType: ctx.ComponentType,
		ComponentId:   componentId,
		Name:          name,
		Category:      category,
		Description:   description,
		Context:       ctx,
	}

	return metadata, nil
}
