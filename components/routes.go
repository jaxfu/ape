package components

type Route struct {
	ComponentMetadata
	RouteMetadata RouteMetadata
	Request       Request      `json:"request" toml:"request"`
	Responses     ResponsesMap `json:"responses" toml:"responses"`
}

type RouteMetadata struct {
	Url    string     `json:"url" toml:"url"`
	Method HttpMethod `json:"method" toml:"method"`
}

type HttpMethod = string

const (
	HTTP_METHOD_GET  HttpMethod = "GET"
	HTTP_METHOD_POST HttpMethod = "POST"
)
