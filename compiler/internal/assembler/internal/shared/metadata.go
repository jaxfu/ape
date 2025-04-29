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
	name := ""
	if ctx.Name != nil {
		name = *ctx.Name
	}

	meta := components.ComponentMetadata{
		ComponentType: ctx.ComponentType,
		Name:          name,
		IsRoot:        ctx.IsRoot,
		ParentId:      ctx.ParentId,
		Category:      metadata.Category,
		Description:   metadata.Description,
	}
	componentId, err := components.GenerateComponentId(meta)
	if err != nil {
		return components.ComponentMetadata{},
			fmt.Errorf("IdHandler.Generate: %+v", err)
	}
	meta.ComponentId = componentId
	if meta.Name == "" {
		meta.Name = componentId
	}

	return meta, nil
}
