package body

import (
	"github.com/jaxfu/ape/components/internal/props"
	"github.com/jaxfu/ape/components/internal/refs"
	"github.com/jaxfu/ape/components/internal/shared"
)

type MessageBody struct {
	ComponentMetadata shared.ComponentMetadata
	BodyType          MessageBodyType
	Ref               refs.ReferenceTag
	Props             props.PropsMap
}

type MessageBodyType = string

const (
	REF   MessageBodyType = "REF"
	PROPS MessageBodyType = "PROPS"
)
