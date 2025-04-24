package bus

type Status struct {
	Running chan bool
	Error   chan error
}

func NewStatus() Status {
	return Status{
		Running: make(chan bool),
		Error:   make(chan error),
	}
}
