package components

type MessageBody struct {
	ComponentMetadata
	BodyType MessageBodyType
	Ref      ReferenceTag
	Props    PropsMap
}

type MessageBodyType = string

const (
	MESSAGE_BODY_TYPE_REF   MessageBodyType = "REF"
	MESSAGE_BODY_TYPE_PROPS MessageBodyType = "PROPS"
)
