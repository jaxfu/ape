package cache

import (
	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/cache/internal"
)

type ComponentCacheInterface interface {
	AddObject(components.Object) error
	GetObject(components.ComponentId) components.Object
	AddProp(components.Prop) error
	GetProp(components.ComponentId) components.Prop
}

func NewComponentCache() ComponentCacheInterface {
	return internal.DefaultComponentCache()
}
