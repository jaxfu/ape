package bus

type Event struct {
	EventType EventType
	Component any
}

type EventType = string

const (
	EVENT_CREATE_COMPONENT EventType = "CREATE"
)
