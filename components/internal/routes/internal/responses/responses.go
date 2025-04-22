package responses

import (
	"github.com/jaxfu/ape/components/internal/routes/internal/body"
	"github.com/jaxfu/ape/components/internal/shared"
)

type (
	ResponsesMap = map[string]Response
)

type Response struct {
	ComponentMetadata shared.ComponentMetadata
	Name              string            `json:"name" toml:"name"`
	StatusCode        uint              `json:"status_code" toml:"status_code"`
	Body              *body.MessageBody `json:"body" toml:"body"`
}
