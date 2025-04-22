package linker

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler"
	"github.com/jaxfu/ape/engine/core/linker/internal"
	"github.com/jaxfu/ape/engine/core/store"
)

type Linker interface {
	LinkRef(id string) bool
	LinkComponent(metadata compiler.CompiledComponentMetadata) (components.ComponentMetadata, error)
	LinkAll(compiler.CompiledComponents) (components.AllComponents, error)
}

func NewLinker(store *store.Store) Linker {
	return internal.NewLinker(store)
}
