package request

import (
	"github.com/jaxfu/ape/components/internal/routes/internal/body"
	"github.com/jaxfu/ape/components/internal/shared"
)

type (
	HeadersMap = map[string]string
)

type Request struct {
	ComponentMetadata shared.ComponentMetadata
	Headers           HeadersMap
	Body              *body.MessageBody
}
