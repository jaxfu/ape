package internal

import (
	"fmt"
	"net/http"

	"github.com/jaxfu/ape/engine/core/events"
)

func (a *Api) GetComponents(w http.ResponseWriter, r *http.Request) {
	result := a.ProcessEvent(events.Event{
		EventType: events.EventTypes.GET_COMPONENTS,
		Component: nil,
	})
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error processing event: %+v\n", result.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result.Bytes)
}
