package components

type Request struct {
	ComponentMetadata
	Headers HeadersMap
	Body    MessageBody
}

type HeadersMap = map[string]string
