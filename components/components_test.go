package components

import (
	"testing"
)

var (
	parentId = "objects.Todo"
	types    = ComponentTypesImpl
	meta     = ComponentMetadata{
		ComponentType: types.PROP,
		Name:          "Username",
		IsRoot:        false,
		ParentId:      &parentId,
	}
)

func TestGenerateComponentId(t *testing.T) {
	id, err := GenerateComponentId(meta)
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Log(id)
}
