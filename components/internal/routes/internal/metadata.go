package internal

type RouteMetadata struct {
	Url    string     `json:"url" toml:"url"`
	Method HttpMethod `json:"method" toml:"method"`
}

type HttpMethod = string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)
