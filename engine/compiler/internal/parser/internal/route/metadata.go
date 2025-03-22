package route

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/compiler/internal/parser/internal/shared"
)

const (
	KEY_URL    string = "url"
	KEY_METHOD string = "method"
)

func parseRouteMetadata(fields map[string]any) (ParsedRouteMetadata, error) {
	url, _, err := shared.GetStringFromMap(fields, KEY_URL)
	if err != nil {
		return ParsedRouteMetadata{}, fmt.Errorf("url not found")
	}

	method, _, _ := shared.GetStringFromMap(fields, KEY_METHOD)

	return ParsedRouteMetadata{
		Url:    url,
		Method: &method,
	}, nil
}

type ParsedRouteMetadata struct {
	Url    string                 `json:"url,omitempty" toml:"url,omitempty"`
	Method *components.HttpMethod `json:"method,omitempty" toml:"method,omitempty"`
}
