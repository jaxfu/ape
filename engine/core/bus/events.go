package bus

type Event struct {
	EventType EventType
	Component Component
}

type (
	EventType           = string
	EventTypesInterface struct {
		CREATE EventType
	}
)

var EventTypes = EventTypesInterface{
	CREATE: "CREATE",
}
