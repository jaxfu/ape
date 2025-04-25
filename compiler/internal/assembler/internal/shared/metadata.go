package shared

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/ape/components"
)

func AssembleComponentMetadata(
	metadata parser.ParsedComponentMetadata,
	ctx shared.CompilationContext,
) (components.ComponentMetadata, error) {
	componentId, err := GenerateComponentId(
		GenerateIdParams{
			ComponentType: ctx.ComponentType,
			Name:          ctx.Name,
			IsRoot:        ctx.IsRoot,
			ParentId:      ctx.ParentId,
			Category:      metadata.Category,
		})
	if err != nil {
		return components.ComponentMetadata{},
			fmt.Errorf("IdHandler.Generate: %+v", err)
	}

	name := ""
	if ctx.Name == nil {
		name = componentId
	} else {
		name = *ctx.Name
	}

	return components.ComponentMetadata{
		ComponentType: ctx.ComponentType,
		Name:          name,
		ComponentId:   componentId,
		IsRoot:        ctx.IsRoot,
		ParentId:      ctx.ParentId,
		Category:      metadata.Category,
		Description:   metadata.Description,
	}, nil
}
