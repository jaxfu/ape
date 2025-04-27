package internal

import (
	"fmt"
	"time"

	"github.com/jaxfu/ape/engine/core/events"
)

func (a *Api) ProcessEvent(event events.Event) events.Result {
	resultChan := make(chan events.Result, 1)
	event.ResultChan = resultChan
	a.Bus.Events <- event

	select {
	case res := <-resultChan:
		return res
	case <-time.After(5 * time.Second):
		return events.Result{
			Error: fmt.Errorf("timeout"),
		}
	}
}
