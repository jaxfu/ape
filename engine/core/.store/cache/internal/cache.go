package internal

import "github.com/jaxfu/ape/components"

type ComponentCache struct {
	Objects []components.Object
}

func DefaultComponentCache() *ComponentCache {
	return &ComponentCache{}
}

func (c *ComponentCache) AddObject(components.Object) error {
	return nil
}

func (c *ComponentCache) GetObject(id components.ComponentId) components.Object {
	return components.Object{}
}

func (c *ComponentCache) AddProp(components.Prop) error {
	return nil
}

func (c *ComponentCache) GetProp(id components.ComponentId) components.Prop {
	return components.Prop{}
}
