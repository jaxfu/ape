package refs

import "github.com/jaxfu/ape/components/internal/shared"

func NewReference(tag ReferenceTag) Reference {
	return Reference{TargetId: tag}
}

type ReferenceTag = string

type Reference struct {
	LinkedComponentId *shared.ComponentId
	TargetId          ReferenceTag
}

func (r Reference) IsLinked() bool {
	return r.LinkedComponentId != nil
}
