package components

type (
	ResponsesMap = map[string]Response
)

type Response struct {
	ComponentMetadata
	Name       string       `json:"name" toml:"name"`
	StatusCode uint         `json:"status_code" toml:"status_code"`
	Body       *MessageBody `json:"body" toml:"body"`
}
