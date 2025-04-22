package shared

import (
	"fmt"

	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
)

func AssembleComponentMetadata(
	metadata parser.ParsedComponentMetadata,
	ctx shared.CompilationContext,
) (shared.CompiledComponentMetadata, error) {
	componentId, err := GenerateComponentId(
		GenerateIdParams{
			ComponentType: ctx.ComponentType,
			Name:          ctx.Name,
			IsRoot:        ctx.IsRoot,
			ParentId:      ctx.ParentId,
			Category:      metadata.Category,
		})
	if err != nil {
		return shared.CompiledComponentMetadata{},
			fmt.Errorf("IdHandler.Generate: %+v", err)
	}

	name := ""
	if ctx.Name == nil {
		name = componentId
	} else {
		name = *ctx.Name
	}

	return shared.CompiledComponentMetadata{
		ComponentType: ctx.ComponentType,
		Name:          name,
		ComponentId:   componentId,
		IsRoot:        ctx.IsRoot,
		Category:      metadata.Category,
		Description:   metadata.Description,
	}, nil
}
